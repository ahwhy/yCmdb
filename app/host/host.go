package host

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ahwhy/yCmdb/app/resource"

	"github.com/infraboard/mcube/types/ftime"
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "Host"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewDefaultHost() *Host {
	return &Host{
		Base: &resource.Base{
			ResourceType: resource.Type_HOST,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func (h *Host) Put(req *UpdateHostData) {
	h.Information = req.Information
	h.Describe = req.Describe
	h.Information.UpdateAt = ftime.Now().Timestamp() // time 时间戳
	h.GenHash()
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := objectPatch(h.Information, req.Information)
	if err != nil {
		return err
	}

	err = objectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Information.UpdateAt = ftime.Now().Timestamp() // time 时间戳
	h.GenHash()

	return nil
}

// patch JSON {a: 1, b: 2}, {b: 20}  ==> {a: 1, b: 20}
func objectPatch(old, new interface{}) error {
	// {b: 20}
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}

	// {a: 1, b: 20}
	return json.Unmarshal(newByte, old)
}

func (h *Host) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Information)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.Base.ResourceHash = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h.Describe)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.Base.DescribeHash = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

func (d *Describe) KeyPairNameToString() string {
	return strings.Join(d.KeyPairName, ",")
}

func (d *Describe) SecurityGroupsToString() string {
	return strings.Join(d.SecurityGroups, ",")
}

func (d *Describe) LoadKeyPairNameString(s string) {
	if s != "" {
		d.KeyPairName = strings.Split(s, ",")
	}
}

func (d *Describe) LoadSecurityGroupsString(s string) {
	if s != "" {
		d.SecurityGroups = strings.Split(s, ",")
	}
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

func (s *HostSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}

	return &QueryHostRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     UpdateMode_PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

func NewDeleteHostRequestWithID(id string) *DeleteHostRequest {
	return &DeleteHostRequest{Id: id}
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewHostSet(),
	}
}

type PagerResult struct {
	Data    *HostSet
	Err     error
	HasNext bool
}
