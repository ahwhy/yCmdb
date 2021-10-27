package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/api/pkg/syncer"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) CreateSecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := syncer.NewCreateSecretRequest()

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

	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QuerySecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := syncer.NewQuerySecretRequest()
	set, err := h.service.QuerySecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeSecret(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := syncer.NewDescribeSecretRequest(ps.ByName("id"))
	ins, err := h.service.DescribeSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Desense()
	response.Success(w, ins)
}
