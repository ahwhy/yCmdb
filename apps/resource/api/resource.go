package http

import (
	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) SearchResource(r *restful.Request, w *restful.Response) {
	query, err := resource.NewSearchRequestFromHTTP(r.Request)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("new request error, %s", err))
		return
	}

	set, err := h.service.Search(r.Request.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
