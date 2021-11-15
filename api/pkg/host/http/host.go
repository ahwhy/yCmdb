package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/api/pkg/host"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ins := host.NewDefaultHost()

	// // GetDataFromRequest 检测请求大小
	// if r.ContentLength == 0 {
	// 	return nil, exception.NewBadRequest("request body is empty")
	// }
	// if r.ContentLength > BodyMaxContenxLength {
	// 	return nil, exception.NewBadRequest(
	// 		"the body exceeding the maximum limit, max size %dM",
	// 		BodyMaxContenxLength/1024/1024)
	// }

	// // 读取body数据
	// body, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()

	// if err != nil {
	// 	return nil, exception.NewBadRequest(
	// 		fmt.Sprintf("read request body error, %s", err))
	// }

	// json.Unmarshal(body, v)

	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveHost(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := host.NewQueryHostRequestFromHTTP(r)
	set, err := h.service.QueryHost(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDescribeHostRequestWithID(ps.ByName("id"))
	set, err := h.service.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) PatchHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewUpdateHostRequest(ps.ByName("id"))
	req.UpdateMode = host.PATCH

	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PutHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewUpdateHostRequest(ps.ByName("id"))

	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDeleteHostRequestWithID(ps.ByName("id"))
	set, err := h.service.DeleteHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
