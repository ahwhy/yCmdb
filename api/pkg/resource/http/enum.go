package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	"github.com/ahwhy/yCmdb/api/pkg/resource"
	"github.com/ahwhy/yCmdb/api/utils"
)

func (h *handler) ListVendor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: resource.VendorAliYun.String(), Describe: "阿里云"},
		{Value: resource.VendorTencent.String(), Describe: "腾讯云"},
		{Value: resource.VendorHuaWei.String(), Describe: "华为云"},
		{Value: resource.VendorVsphere.String(), Describe: "Vsphere"},
	}

	response.Success(w, resp)
}
