package http

import (
	"fmt"

	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/secret"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service secret.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Secret")
	if pkg.Secret == nil {
		return fmt.Errorf("dependence service secret not ready")
	}

	h.service = pkg.Secret

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

	r.DELETE("/secrets/:id", api.DeleteSecret)

	r.GET("/crendential_types", api.ListCrendentialType)

	// r.POST("/secrets/:id/sync", api.Sync)
}
