package rest

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

type Api interface {
	GetResourceApi(model.ResourceType) (ResourceApi, error)
}

type ResourceApi interface {
	List(mesh string) string
	Item(mesh string, name string) string
}

func NewResourceApi(path string) ResourceApi {
	return &resourceApi{CollectionPath: path}
}

type resourceApi struct {
	CollectionPath string
}

func (r *resourceApi) List(mesh string) string {
	if mesh == "" {
		return fmt.Sprintf("/%s", r.CollectionPath)
	} else {
		return fmt.Sprintf("/meshes/%s/%s", mesh, r.CollectionPath)
	}
}

func (r *resourceApi) Item(mesh string, name string) string {
	if mesh == "" {
		return fmt.Sprintf("/%s/%s", r.CollectionPath, name)
	} else {
		return fmt.Sprintf("/meshes/%s/%s/%s", mesh, r.CollectionPath, name)
	}
}

var _ Api = &ApiDescriptor{}

type ApiDescriptor struct {
	Resources map[model.ResourceType]ResourceApi
}

func (m *ApiDescriptor) GetResourceApi(typ model.ResourceType) (ResourceApi, error) {
	mapping, ok := m.Resources[typ]
	if !ok {
		return nil, errors.Errorf("unknown resource type: %q", typ)
	}
	return mapping, nil
}
