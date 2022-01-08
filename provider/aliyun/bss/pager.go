package bss

import (
	"github.com/ahwhy/yCmdb/app/bill"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func newPager(pageSize int, operater *BssOperater, rate int, month string) *pager {
	req := bssopenapi.CreateQueryInstanceBillRequest()
	req.IsHideZeroCharge = requests.NewBoolean(true)
	req.PageSize = requests.NewInteger(pageSize)
	req.BillingCycle = month
	retaFloat := 1 / float64(rate)

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
		tb:       tokenbucket.NewBucketWithRate(retaFloat, 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *BssOperater
	req      *bssopenapi.QueryInstanceBillRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Next() *bill.PagerResult {
	result := bill.NewPagerResult()

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

func (p *pager) nextReq() *bssopenapi.QueryInstanceBillRequest {
	// 等待一个可用的token
	p.tb.Wait(1)
	p.log.Debugf("请求第%d页数据", p.number)
	p.req.PageNum = requests.NewInteger(p.number)

	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}
