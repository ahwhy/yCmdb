package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/api/pkg/task"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) CreatTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	req := task.NewCreateTaskRequst()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	// req.SecretId = ps.ByName("id")

	// query := r.URL.Query()
	// // 解析类型
	// t := resource.LoadTypeFromString(query.Get("resource_type"))
	// if t == resource.Unsuport {
	// 	response.Failed(w, exception.NewBadRequest("unsuport resource_type %s", t))
	// 	return
	// }
	// req.ResourceType = t
	// req.Region = query.Get("region")

	set, err := h.task.CreatTask(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := task.NewQueryTaskRequestFromHTTP(r)
	set, err := h.task.QueryTask(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
