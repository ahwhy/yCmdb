package http

import (
	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/host"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service host.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(host.AppName)
	// 使用GetGrpcApp方法，http接口的暴露，需要依赖后端的grpc接口
	// 将GRPCApp这个接口对象，断言成host.ServiceServer接口，以使用对应的方法
	h.service = app.GetGrpcApp(host.AppName).(host.ServiceServer)

	return nil
}

func (h *handler) Name() string {
	return host.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	hr := r.ResourceRouter("host")
	hr.Permission(true)
	hr.Handle("POST", "/hosts", h.CreateHost).AddLabel(label.Create)
	hr.Handle("GET", "/hosts", h.QueryHost).AddLabel(label.List)
	hr.Handle("GET", "/hosts/:id", h.DescribeHost).AddLabel(label.Get)
	hr.Handle("PATCH", "/hosts/:id", h.PatchHost).AddLabel(label.Update)
	hr.Handle("PUT", "/hosts/:id", h.PutHost).AddLabel(label.Update)
	hr.Handle("DELETE", "/hosts/:id", h.DeleteHost).AddLabel(label.Delete)
}

func init() {
	app.RegistryHttpApp(h)
}
