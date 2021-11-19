package http

import (
	"fmt"

	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/task"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

var (
	api = &handler{}
)

type handler struct {
	task task.Service
	log  logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Task")
	if pkg.Task == nil {
		return fmt.Errorf("dependence service task not ready")
	}

	h.task = pkg.Task

	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()

	r.GET("/tasks", api.QueryTask)
	r.POST("/tasks/:id", api.CreatTask)
}
