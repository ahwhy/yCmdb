package rds

import (
	"github.com/ahwhy/yCmdb/apps/rds"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"
)

func (o *RdsOperater) PageQuery() rds.Pager {
	return newPager(20, o)
}

func (o *RdsOperater) Query(req *model.ListInstancesRequest) (*rds.RDSSet, error) {
	set := rds.NewRDSSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.Instances).Items

	return set, nil
}
