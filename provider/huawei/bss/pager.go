package bss

import (
	"github.com/ahwhy/yCmdb/app/bill"
	"github.com/ahwhy/yCmdb/utils"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
)

func newPager(pageSize int, operater *BssOperater, rate int, month string) *pager {
	req := &model.ListCustomerselfResourceRecordsRequest{}
	req.Cycle = month
	req.Limit = utils.Int32Ptr(int32(pageSize))
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
	operater *BssOperater
	req      *model.ListCustomerselfResourceRecordsRequest
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

	p.total = resp.Total
	result.Data = resp
	result.HasNext = p.hasNext()
	p.number++

	return result
}

func (p *pager) nextReq() *model.ListCustomerselfResourceRecordsRequest {
	p.log.Debugf("请求第%d页数据", p.number)
	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.number))

	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
