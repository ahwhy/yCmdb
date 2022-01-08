package bss

import (
	"github.com/ahwhy/yCmdb/app/bill"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func (o *BssOperater) PageQuery(req *PageQueryRequest) bill.Pager {
	return newPager(20, o, req.Rate, req.Month)
}

func (o *BssOperater) Query(req *bssopenapi.QueryInstanceBillRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.QueryInstanceBill(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(resp.Data.TotalCount)
	set.Items = o.transferSet(resp.Data).Items

	return set, nil
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate  int
	Month string
}
