package http

import (
	"fmt"

	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/resource"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	service resource.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Resource")
	if pkg.Resource == nil {
		return fmt.Errorf("dependence service resource not ready")
	}

	h.service = pkg.Resource

	return nil
}

var (
	api = &handler{}
)

func RegistAPI(r *httprouter.Router) {
	api.Config()

	r.GET("/search", api.SearchResource)
	r.GET("/vendors", api.ListVendor)
	r.GET("/regions", api.ListVendorRegion)
}
