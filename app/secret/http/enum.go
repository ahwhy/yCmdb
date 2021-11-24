package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/app/secret"
	"github.com/ahwhy/yCmdb/utils"

	"github.com/infraboard/mcube/http/response"
)

func (h *handler) ListCrendentialType(w http.ResponseWriter, r *http.Request) {
	resp := []utils.EnumDescribe{
		{Value: secret.CrendentialType_API_KEY.String(), Describe: "API凭证"},
		{Value: secret.CrendentialType_PASSWORD.String(), Describe: "用户名密码"},
	}

	response.Success(w, resp)
}
