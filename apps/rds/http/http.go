package http

import (
	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/rds"
	
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service rds.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(rds.AppName)
	h.service = app.GetGrpcApp(rds.AppName).(rds.ServiceServer)

	return nil
}

func (h *handler) Name() string {
	return rds.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	hr := r.ResourceRouter("bill")
	hr.Permission(true)
	// hr.Handle("POST", "/rds", h.SaveRds)
	// hr.Handle("GET", "/rds", h.QueryRds)
}

func init() {
	app.RegistryHttpApp(h)
}
