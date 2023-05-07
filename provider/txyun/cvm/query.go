package cvm

import (
	"github.com/ahwhy/yCmdb/apps/host"
	"github.com/ahwhy/yCmdb/utils"
	
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (o *CVMOperater) PageQuery() host.Pager {
	return newPager(20, o)
}

func (o *CVMOperater) Query(req *cvm.DescribeInstancesRequest) (*host.HostSet, error) {
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	set := o.transferSet(resp.Response.InstanceSet)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)

	return set, nil
}
