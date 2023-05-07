package impl

import (
	"context"
	"fmt"

	"github.com/ahwhy/yCmdb/app/bill"
	"github.com/ahwhy/yCmdb/app/resource"
	"github.com/ahwhy/yCmdb/app/secret"
	"github.com/ahwhy/yCmdb/app/task"

	bssOp "github.com/ahwhy/yCmdb/provider/aliyun/bss"
	aliConn "github.com/ahwhy/yCmdb/provider/aliyun/connectivity"
	hwBssOp "github.com/ahwhy/yCmdb/provider/huawei/bss"
	hwConn "github.com/ahwhy/yCmdb/provider/huawei/connectivity"
	billOp "github.com/ahwhy/yCmdb/provider/txyun/billing"
	txConn "github.com/ahwhy/yCmdb/provider/txyun/connectivity"
)

func (s *service) syncBill(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager bill.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun bill ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}

		operater := bssOp.NewBssOperater(bc)
		req := bssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun bill ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := billOp.NewBillingOperater(client.BillingClient())
		req := billOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun bill ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwBssOp.NewBssOperater(bc)
		req := hwBssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	default:
		t.Failed(fmt.Sprintf("unsuport bill syncing vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		hasNext := true
		for hasNext {
			p := pager.Next()
			hasNext = p.HasNext

			if p.Err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", p.Err))
				return
			}

			// 调用bill服务保持数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				b, err := s.bill.SaveBill(ctx, target)
				if err != nil {
					s.log.Warnf("save bill error, %s", err)
				} else {
					s.log.Debugf("save host %s to db", b.ShortDesc())
				}
			}
		}
	}
}
