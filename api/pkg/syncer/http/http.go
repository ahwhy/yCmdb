package http

import (
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"

	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/syncer"
)

type handler struct {
	service syncer.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Syncer")
	if pkg.Syncer == nil {
		return fmt.Errorf("dependence service syncer not ready")
	}
	h.service = pkg.Syncer

	return nil
}

var (
	api = &handler{}
)

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.POST("/secrets", api.CreateSecret)
	r.GET("/secrets", api.QuerySecret)
	r.GET("/secrets/:id", api.DescribeSecret)
	r.POST("/secrets/:id/sync", api.Sync)
}
