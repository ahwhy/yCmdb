package rds

import (
	cmdbRds "github.com/ahwhy/yCmdb/app/rds"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

func (o *RdsOperater) PageQuery(req *PageQueryRequest) cmdbRds.Pager {
	return newPager(20, o, req.Rate)
}

func (o *RdsOperater) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.RDSSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := cmdbRds.NewRDSSet()

	// 更详细的数据 需要通过DescribeDBInstanceAttribute获取，比如cpu和内存信息
	for _, ins := range resp.Items.DBInstance {
		descReq := rds.CreateDescribeDBInstanceAttributeRequest()
		descReq.DBInstanceId = ins.DBInstanceId
		detail, err := o.client.DescribeDBInstanceAttribute(descReq)
		if err != nil {
			return nil, err
		}
		set.AddSet(o.transferSet(detail.Items.DBInstanceAttribute))
	}
	set.Total = int64(resp.TotalRecordCount)

	return set, nil
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate int
}
