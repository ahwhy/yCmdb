package billing

import (
	"fmt"
	"strconv"

	"github.com/ahwhy/yCmdb/app/bill"
	"github.com/ahwhy/yCmdb/app/resource"
	"github.com/ahwhy/yCmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
)

func NewBillingOperater(client *billing.Client) *BillingOperater {
	return &BillingOperater{
		client: client,
		log:    zap.L().Named("Tx Billing"),
	}
}

type BillingOperater struct {
	client *billing.Client
	log    logger.Logger
}

func (o *BillingOperater) transferSet(items []*billing.BillResourceSummary, month string) *bill.BillSet {
	set := bill.NewBillSet()
	for i := range items {
		ins := o.transferOne(items[i])
		ins.Vendor = resource.Vendor_TENCENT
		ins.Month = month
		set.Add(ins)
	}

	return set
}

func (o *BillingOperater) transferOne(ins *billing.BillResourceSummary) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = utils.PtrStrV(ins.OwnerUin)
	b.ProductCode = utils.PtrStrV(ins.ProductCode)
	b.ProductType = utils.PtrStrV(ins.ProductCodeName)
	b.PayMode = utils.PtrStrV(ins.PayModeName)
	b.PayModeDetail = utils.PtrStrV(ins.ActionTypeName)
	b.OrderId = utils.PtrStrV(ins.OrderId)
	b.InstanceId = utils.PtrStrV(ins.ResourceId)
	b.InstanceName = utils.PtrStrV(ins.ResourceName)
	b.InstanceConfig = utils.PtrStrV(ins.ConfigDesc)
	b.RegionCode = fmt.Sprintf("%d", utils.PtrInt64(ins.RegionId))
	b.RegionName = utils.PtrStrV(ins.RegionName)

	b.SalePrice, _ = strconv.ParseFloat(utils.PtrStrV(ins.TotalCost), 64)
	b.RealCost, _ = strconv.ParseFloat(utils.PtrStrV(ins.RealTotalCost), 64)
	b.VoucherPay, _ = strconv.ParseFloat(utils.PtrStrV(ins.VoucherPayAmount), 64)
	b.CashPay, _ = strconv.ParseFloat(utils.PtrStrV(ins.CashPayAmount), 64)

	return b
}
