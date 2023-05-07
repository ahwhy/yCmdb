package rds

import (
	"github.com/ahwhy/yCmdb/apps/rds"
	"github.com/ahwhy/yCmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"
)

func newPager(pageSize int, operater *RdsOperater) *pager {
	req := &model.ListInstancesRequest{}
	req.Limit = utils.Int32Ptr(int32(pageSize))

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *RdsOperater
	req      *model.ListInstancesRequest
	log      logger.Logger
}

func (p *pager) Next() *rds.PagerResult {
	result := rds.NewPagerResult()

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

func (p *pager) nextReq() *model.ListInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.number)
	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.number))

	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
