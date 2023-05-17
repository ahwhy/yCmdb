package http

import (
	"github.com/ahwhy/yCmdb/apps"
	"github.com/ahwhy/yCmdb/apps/bill"

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
	h.service = apps.GetGrpcApp(bill.AppName).(bill.ServiceServer)

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
	apps.RegistryHttpApp(h)
}
