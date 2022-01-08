package rds

import (
	cmdbRds "github.com/ahwhy/yCmdb/app/rds"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

func (o *RdsOperater) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Items.DBInstance)
	set.Total = int64(resp.TotalRecordCount)

	return set, nil
}

func (o *RdsOperater) PageQuery(req *PageQueryRequest) cmdbRds.Pager {
	return newPager(20, o, req.Rate)
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate int
}
