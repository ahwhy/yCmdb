package http

import (
	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/task"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	task task.ServiceServer
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Task")
	h.task = app.GetGrpcApp(task.AppName).(task.ServiceServer)

	return nil
}

func (h *handler) Name() string {
	return task.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	tr := r.ResourceRouter("task")
	tr.Permission(true)
	tr.Handle("GET", "/tasks", h.QueryTask).AddLabel(label.List)
	tr.Handle("POST", "/tasks", h.CreatTask).AddLabel(label.Create)
}

func init() {
	app.RegistryHttpApp(h)
}
