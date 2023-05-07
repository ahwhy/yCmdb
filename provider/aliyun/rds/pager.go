package rds

import (
	cmdbRds "github.com/ahwhy/yCmdb/apps/rds"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *RdsOperater, rate int) *pager {
	req := rds.CreateDescribeDBInstancesRequest()
	req.PageSize = requests.NewInteger(pageSize)
	rateFloat := 1 / float64(rate)

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
		tb:       tokenbucket.NewBucketWithRate(rateFloat, 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *RdsOperater
	req      *rds.DescribeDBInstancesRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Next() *cmdbRds.PagerResult {
	result := cmdbRds.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}

	p.total = int64(resp.Total)
	result.Data = resp
	result.HasNext = p.hasNext()
	p.number++

	return result
}

func (p *pager) nextReq() *rds.DescribeDBInstancesRequest {
	// 等待一个可用token
	p.tb.Wait(1)
	p.log.Debug("请求第%d页数据", p.number)
	p.req.PageNumber = requests.NewInteger(p.number)

	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}
