package ecs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/ahwhy/yCmdb/api/pkg/host"
)

func (o *EcsOperater) Query(req *model.ListServersDetailsRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.ListServersDetails(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.Count)
	set.Items = o.transferSet(resp.Servers).Items

	return set, nil
}

func (o *EcsOperater) PageQuery() host.Pager {
	return newPager(20, o)
}
