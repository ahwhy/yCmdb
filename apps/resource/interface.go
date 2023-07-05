package resource

import (
	"context"
	"fmt"
	"hash/fnv"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
)

var (
	validate = validator.New()
)

const (
	AppName = "Resource"
)

type Service interface {
	Put(context.Context, *Resource) (*Resource, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	RPCServer
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewSearchRequestFromHTTP(r *http.Request) (*SearchRequest, error) {
	qs := r.URL.Query()
	req := &SearchRequest{
		Page:      request.NewPageRequestFromHTTP(r),
		Keywords:  qs.Get("keywords"),
		Domain:    qs.Get("domain"),
		Namespace: qs.Get("namespace"),
		Env:       qs.Get("env"),
		Status:    qs.Get("status"),
		Owner:     qs.Get("owner"),
		WithTags:  qs.Get("with_tags") == "true",
		Tags:      []*TagSelector{},
	}

	umStr := qs.Get("usage_mode")
	if umStr != "" {
		mode, err := ParseUSAGE_MODEFromString(umStr)
		if err != nil {
			return nil, err
		}
		req.UsageMode = &mode
	}

	rtStr := qs.Get("resource_type")
	if rtStr != "" {
		rt, err := ParseTYPEFromString(rtStr)
		if err != nil {
			return nil, err
		}
		req.Type = &rt
	}

	tgStr := qs.Get("tag")
	if tgStr != "" {
		tg, err := NewTagsFromString(tgStr)
		if err != nil {
			return nil, err
		}
		req.AddTag(tg...)
	}

	return req, nil
}

func (req *SearchRequest) HasTag() bool {
	if req.Tags == nil {
		return false
	}
	return len(req.Tags) > 0
}

func (req *SearchRequest) AddTag(t ...*TagSelector) {
	req.Tags = append(req.Tags, t...)
}

func NewDefaultResource(rt TYPE) *Resource {
	return &Resource{
		Meta: &Meta{},
		Spec: &Spec{
			ResourceType: rt,
			Tags:         []*Tag{},
		},
		Cost:   &Cost{},
		Status: &Status{},
	}
}

func (r *Resource) AddTag(t *Tag) {
	r.Spec.Tags = append(r.Spec.Tags, t)
}

func (i *Status) PrivateIPToString() string {
	return strings.Join(i.PrivateAddress, ",")
}

func (i *Status) PublicIPToString() string {
	return strings.Join(i.PublicAddress, ",")
}

func (i *Status) LoadPrivateIPString(s string) {
	if s != "" {
		i.PrivateAddress = strings.Split(s, ",")
	}
}

func (i *Status) LoadPublicIPString(s string) {
	if s != "" {
		i.PublicAddress = strings.Split(s, ",")
	}
}

func NewDefaultTag() *Tag {
	return &Tag{
		Purpose: TAG_PURPOSE_GROUP,
		Weight:  1,
	}
}

func NewGroupTag(key, value string) *Tag {
	return &Tag{
		Purpose: TAG_PURPOSE_GROUP,
		Key:     key,
		Value:   value,
		Weight:  1,
	}
}

type AccountGetter struct {
	accountId string
}

func (ag *AccountGetter) WithAccountId(id string) {
	ag.accountId = id
}

func (ag *AccountGetter) GetAccountId() string {
	return ag.accountId
}

// key1=v1,v2,v3&key2=~v1,v2,v3
func NewTagsFromString(tagStr string) (tags []*TagSelector, err error)