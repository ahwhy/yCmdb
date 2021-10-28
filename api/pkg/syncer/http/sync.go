package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/api/pkg/resource"
	"github.com/ahwhy/yCmdb/api/pkg/syncer"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) Sync(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query := r.URL.Query()
	req := syncer.NewSyncRequest(ps.ByName("id"))

	// 解析类型
	t := resource.LoadTypeFromString(query.Get("resource_type"))
	if t == resource.Unsuport {
		response.Failed(w, exception.NewBadRequest("unsuport resource_type %s", t))
		return
	}
	req.ResourceType = t
	req.Region = query.Get("region")

	set, err := h.service.Sync(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
