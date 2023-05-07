package bss

import (
	"github.com/ahwhy/yCmdb/apps/bill"
	"github.com/ahwhy/yCmdb/apps/resource"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func NewBssOperater(client *bssopenapi.Client) *BssOperater {
	return &BssOperater{
		client: client,
		log:    zap.L().Named("ALI BSS"),
	}
}

type BssOperater struct {
	client *bssopenapi.Client
	log    logger.Logger
}

func (o *BssOperater) transferSet(list bssopenapi.DataInQueryInstanceBill) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items.Item
	for i := range items {
		ins := o.transferOne(items[i])
		ins.Vendor = resource.Vendor_ALIYUN
		ins.Month = list.BillingCycle
		set.Add(ins)
	}

	return set
}

func (o *BssOperater) transferOne(ins bssopenapi.Item) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = ins.OwnerID
	b.OwnerName = ins.OwnerName
	b.ProductType = ins.ProductType
	b.ProductCode = ins.ProductCode
	b.ProductDetail = ins.ProductDetail
	b.PayMode = ins.Item
	b.PayModeDetail = ins.BillingType
	b.OrderId = ins.SubOrderId
	b.InstanceId = ins.InstanceID
	b.InstanceName = ins.NickName
	b.PublicIp = ins.InternetIP
	b.PrivateIp = ins.IntranetIP
	b.InstanceConfig = ins.InstanceConfig
	b.RegionCode = ins.RegionNo
	b.RegionName = ins.Region

	b.SalePrice = ins.PretaxGrossAmount
	b.SaveCost = ins.InvoiceDiscount
	b.RealCost = ins.PretaxAmount
	b.StoredcardPay = ins.DeductedByPrepaidCard
	b.VoucherPay = ins.DeductedByCashCoupons
	b.CashPay = ins.PaymentAmount
	b.OutstandingAmount = ins.OutstandingAmount

	return b
}
