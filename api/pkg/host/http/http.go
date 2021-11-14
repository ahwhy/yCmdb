package http

import (
	"fmt"

	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/host"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service host.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Host")
	if pkg.Host == nil {
		return fmt.Errorf("dependence service host not ready")
	}

	h.service = pkg.Host

	return nil
}

var (
	api = &handler{}
)

func RegistAPI(r *httprouter.Router) {
	api.Config()

	r.POST("/hosts", api.CreateHost)

	r.GET("/hosts", api.QueryHost)
	r.GET("/hosts/:id", api.DescribeHost)

	r.PATCH("/hosts/:id", api.PatchHost)
	r.PUT("/hosts/:id", api.PutHost)

	r.DELETE("/hosts/:id", api.DeleteHost)
}
