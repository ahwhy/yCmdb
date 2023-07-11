package http

import (
	"github.com/ahwhy/yCmdb/apps"
	"github.com/ahwhy/yCmdb/apps/resource"

	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service resource.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(resource.AppName)
	h.service = apps.GetGrpcApp(resource.AppName).(resource.Service)

	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"资源检索"}

	ws.Route(ws.GET("/search").To(h.SearchResource).
		Doc("检索资源").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Reads(resource.SearchRequest{}).
		Writes(resource.ResourceSet{}).
		Returns(200, "OK", resource.ResourceSet{}))
}

func init() {
	apps.RegistryRESTfulApp(h)
}
