package http

import (
	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/bill"

	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service bill.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(bill.AppName)
	// 使用GetGrpcApp方法，http接口的暴露，需要依赖后端的grpc接口
	// 将GRPCApp这个接口对象，断言成host.ServiceServer接口，以使用对应的方法
	h.service = app.GetGrpcApp(bill.AppName).(bill.ServiceServer)

	return nil
}

func (h *handler) Name() string {
	return bill.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	hr := r.ResourceRouter("bill")
	hr.Permission(true)
	hr.Handle("POST", "/bills", h.SaveBill)
	hr.Handle("GET", "/bills", h.QueryBill)
}

func init() {
	app.RegistryHttpApp(h)
}
