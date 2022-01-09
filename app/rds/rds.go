package rds

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	resource "github.com/ahwhy/yCmdb/app/resource"
)

const (
	AppName = "rds"
)

func NewDefaultRDS() *RDS {
	return &RDS{
		Base:        &resource.Base{},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func (r *RDS) ShortDesc() string {
	return fmt.Sprintf("%s %s", r.Information.Name, r.Information.PrivateIp)
}

func (d *Describe) Hash() string {
	hash := sha1.New()
	b, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	hash.Write(b)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func NewSet() *RDSSet {
	return &RDSSet{
		Items: []*RDS{},
	}
}

func (s *RDSSet) Add(item *RDS) {
	s.Items = append(s.Items, item)
}

func (s *RDSSet) AddSet(set *RDSSet) {
	s.Items = append(s.Items, set.Items...)
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewSet(),
	}
}

type PagerResult struct {
	Data    *RDSSet
	Err     error
	HasNext bool
}
