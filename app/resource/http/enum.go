package http

import (
	"net/http"

	"github.com/ahwhy/yCmdb/app/resource"
	"github.com/ahwhy/yCmdb/utils"

	ali_region "github.com/ahwhy/yCmdb/provider/aliyun/region"
	hw_region "github.com/ahwhy/yCmdb/provider/huawei/region"
	tx_region "github.com/ahwhy/yCmdb/provider/txyun/region"

	"github.com/infraboard/mcube/http/response"
)

func (h *handler) ListVendor(w http.ResponseWriter, r *http.Request) {
	resp := []utils.EnumDescribe{
		{Value: resource.Vendor_ALIYUN.String(), Describe: "阿里云"},
		{Value: resource.Vendor_TENCENT.String(), Describe: "腾讯云"},
		{Value: resource.Vendor_HUAWEI.String(), Describe: "华为云"},
		{Value: resource.Vendor_VSPHERE.String(), Describe: "Vsphere"},
	}
	
	response.Success(w, resp)
}

func (h *handler) ListResourceType(w http.ResponseWriter, r *http.Request) {
	resp := map[string][]utils.EnumDescribe{
		resource.Vendor_ALIYUN.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "阿里云ECS"},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "阿里云RDS"},
		},
		resource.Vendor_TENCENT.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "腾讯云CVM"},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "腾讯云CDB"},
		},
		resource.Vendor_HUAWEI.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "华为云ECS"},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "华为云RDS"},
		},
		resource.Vendor_VSPHERE.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "VMware vm"},
		},
	}

	response.Success(w, resp)
}

func (h *handler) ListVendorRegion(w http.ResponseWriter, r *http.Request) {
	resp := map[string][]utils.EnumDescribe{
		resource.Vendor_ALIYUN.String():  ali_region.Regions,
		resource.Vendor_TENCENT.String(): tx_region.Regions,
		resource.Vendor_HUAWEI.String():  hw_region.Regions,
		resource.Vendor_VSPHERE.String(): hw_region.Regions,
	}
	
	response.Success(w, resp)
}
