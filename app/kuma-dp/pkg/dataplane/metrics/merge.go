package metrics

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

const EnvoyClusterLabelName = "envoy_cluster_name"

func MergeClusters(in io.Reader, out io.Writer) error {
	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(in)
	if err != nil {
		return err
	}

	for _, metricFamily := range metricFamilies {
		if !isClusterMetricFamily(metricFamily) {
			if _, err := expfmt.MetricFamilyToText(out, metricFamily); err != nil {
				return err
			}
			if _, err := out.Write([]byte("\n")); err != nil {
				return err
			}
			continue
		}

		// metricsByClusterNames returns the data in the following format:
		// 'cluster_name' ->
		//   - metric1{envoy_cluster_name="cluster_name-_0_",label1="value1"} 10
		//   - metric1{envoy_cluster_name="cluster_name-_1_",label1="value1"} 20
		//   - metric1{envoy_cluster_name="cluster_name-_2_",label1="value1"} 30
		// 'another_cluster_name' ->
		//   - metric1{envoy_cluster_name="another_cluster_name-_0_",response_code="200"} 10
		//   - metric1{envoy_cluster_name="another_cluster_name-_0_",response_code="401"} 20
		//   - metric1{envoy_cluster_name="another_cluster_name-_1_",response_code="200"} 30
		//   - metric1{envoy_cluster_name="another_cluster_name-_2_",response_code="503"} 40
		metricsByClusterName, err := metricsByClusterNames(metricFamily.Metric)
		if err != nil {
			return err
		}

		// renameCluster changes the value of 'envoy_cluster_name' label for every metric.
		// So the data will look like:
		// 'cluster_name' ->
		//   - metric1{envoy_cluster_name="cluster_name",label1="value1"} 10
		//   - metric1{envoy_cluster_name="cluster_name",label1="value1"} 20
		//   - metric1{envoy_cluster_name="cluster_name",label1="value1"} 30
		// 'another_cluster_name' ->
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="200"} 10
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="401"} 20
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="200"} 30
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="503"} 40
		for clusterName, metrics := range metricsByClusterName {
			renameCluster(clusterName, metrics)
		}

		// after the previous step we've got duplicates in the metrics, merge them during this step:
		// 'cluster_name' ->
		//   - metric1{envoy_cluster_name="cluster_name",label1="value1"} 60
		// 'another_cluster_name' ->
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="200"} 40
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="401"} 20
		//   - metric1{envoy_cluster_name="another_cluster_name",response_code="503"} 40
		for clusterName, metrics := range metricsByClusterName {
			metricsByClusterName[clusterName] = mergeDuplicates(metricFamily.Type, metrics)
		}

		metricFamily.Metric = nil
		for _, metric := range metricsByClusterName {
			metricFamily.Metric = append(metricFamily.Metric, metric...)
		}

		if _, err := expfmt.MetricFamilyToText(out, metricFamily); err != nil {
			return err
		}
		if _, err := out.Write([]byte("\n")); err != nil {
			return err
		}
	}

	return nil
}

func renameCluster(clusterName string, metrics []*io_prometheus_client.Metric) {
	for _, metric := range metrics {
		for _, label := range metric.GetLabel() {
			if label.GetName() == EnvoyClusterLabelName {
				label.Value = &clusterName
			}
		}
	}
}

func mergeDuplicates(typ *io_prometheus_client.MetricType, metrics []*io_prometheus_client.Metric) []*io_prometheus_client.Metric {
	hashes := map[string][]*io_prometheus_client.Metric{}
	for _, metric := range metrics {
		hashes[hash(metric)] = append(hashes[hash(metric)], metric)
	}

	var result []*io_prometheus_client.Metric

	for _, dups := range hashes {
		merged := &io_prometheus_client.Metric{
			Label: dups[0].GetLabel(),
		}
		switch *typ {
		case io_prometheus_client.MetricType_COUNTER:
			merged.Counter = mergeCounter(dups)
		case io_prometheus_client.MetricType_GAUGE:
			merged.Gauge = mergeGauge(dups)
		case io_prometheus_client.MetricType_SUMMARY:
			merged.Summary = mergeSummary(dups)
		case io_prometheus_client.MetricType_UNTYPED:
			merged.Untyped = mergeUntyped(dups)
		case io_prometheus_client.MetricType_HISTOGRAM:
			merged.Histogram = mergeHistogram(dups)
		}
		result = append(result, merged)
	}
	return result
}

func hash(metric *io_prometheus_client.Metric) string {
	pairs := []string{}
	for _, l := range metric.GetLabel() {
		pairs = append(pairs, fmt.Sprintf("%s=%s", l.GetName(), l.GetValue()))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ";")
}

func mergeCounter(metrics []*io_prometheus_client.Metric) *io_prometheus_client.Counter {
	var acc float64
	for _, m := range metrics {
		acc += *m.Counter.Value
	}
	return &io_prometheus_client.Counter{Value: &acc}
}

func mergeGauge(metrics []*io_prometheus_client.Metric) *io_prometheus_client.Gauge {
	var acc float64
	for _, m := range metrics {
		acc += *m.Gauge.Value
	}
	return &io_prometheus_client.Gauge{Value: &acc}
}

func mergeUntyped(metrics []*io_prometheus_client.Metric) *io_prometheus_client.Untyped {
	var acc float64
	for _, m := range metrics {
		acc += *m.Untyped.Value
	}
	return &io_prometheus_client.Untyped{Value: &acc}
}

func mergeHistogram(metrics []*io_prometheus_client.Metric) *io_prometheus_client.Histogram {
	bucketMap := map[float64]uint64{}
	var sum float64
	var count uint64
	for _, m := range metrics {
		for _, bucket := range m.Histogram.GetBucket() {
			bucketMap[bucket.GetUpperBound()] += *bucket.CumulativeCount
		}
		sum += m.Histogram.GetSampleSum()
		count += m.Histogram.GetSampleCount()
	}

	buckets := []*io_prometheus_client.Bucket{}
	for upperBound, cumulativeCount := range bucketMap {
		u := upperBound
		c := cumulativeCount
		buckets = append(buckets, &io_prometheus_client.Bucket{
			UpperBound:      &u,
			CumulativeCount: &c,
		})
	}

	return &io_prometheus_client.Histogram{
		SampleCount: &count,
		SampleSum:   &sum,
		Bucket:      buckets,
	}
}

func mergeSummary(metrics []*io_prometheus_client.Metric) *io_prometheus_client.Summary {
	quantileMap := map[float64]float64{}
	var sum float64
	var count uint64
	for _, m := range metrics {
		for _, quantile := range m.Summary.GetQuantile() {
			quantileMap[quantile.GetQuantile()] += quantile.GetValue()
		}
		sum += m.Summary.GetSampleSum()
		count += m.Summary.GetSampleCount()
	}

	quantiles := []*io_prometheus_client.Quantile{}
	for quantile, value := range quantileMap {
		q := quantile
		v := value
		quantiles = append(quantiles, &io_prometheus_client.Quantile{
			Quantile: &q,
			Value:    &v,
		})
	}

	return &io_prometheus_client.Summary{
		SampleCount: &count,
		SampleSum:   &sum,
		Quantile:    quantiles,
	}
}

func isClusterMetricFamily(family *io_prometheus_client.MetricFamily) bool {
	if len(family.Metric) == 0 {
		return false
	}
	_, hasClusterName := getClusterName(family.Metric[0])
	return hasClusterName
}

func metricsByClusterNames(metricsFamily []*io_prometheus_client.Metric) (map[string][]*io_prometheus_client.Metric, error) {
	indexedMetrics := map[string][]*io_prometheus_client.Metric{}
	for _, m := range metricsFamily {
		clusterName, ok := getClusterName(m)
		if !ok {
			return nil, errors.New("failed to get clusterName for non-cluster metric")
		}
		prefix, _, ok := isMergeableClusterName(clusterName)
		if !ok {
			indexedMetrics[clusterName] = append(indexedMetrics[clusterName], m)
			continue
		}
		indexedMetrics[prefix] = append(indexedMetrics[prefix], m)
	}
	return indexedMetrics, nil
}

func getClusterName(metric *io_prometheus_client.Metric) (clusterName string, ok bool) {
	for _, label := range metric.Label {
		if *label.Name == EnvoyClusterLabelName {
			return *label.Value, true
		}
	}
	return "", false
}

func isMergeableClusterName(clusterName string) (prefix string, n int, ok bool) {
	var re = regexp.MustCompile(`(?P<prefix>.*)-_(?P<num>[0-9]+)_`)
	matches := re.FindStringSubmatch(clusterName)
	if matches == nil {
		return "", 0, false
	}

	prefixIndex := re.SubexpIndex("prefix")
	if prefixIndex == -1 {
		return "", 0, false
	}
	numIndex := re.SubexpIndex("num")
	if numIndex == -1 {
		return "", 0, false
	}
	num, err := strconv.Atoi(matches[numIndex])
	if err != nil {
		return "", 0, false
	}
	return matches[prefixIndex], num, true
}
