package cdb

import (
	"github.com/ahwhy/yCmdb/app/rds"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
)

func (o *CDBOperater) Query(req *cdb.DescribeDBInstancesRequest) (*rds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.Items), nil
}
