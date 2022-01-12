package impl

import (
	"context"
	"fmt"

	"github.com/ahwhy/yCmdb/app/host"
	"github.com/ahwhy/yCmdb/app/resource"
	"github.com/ahwhy/yCmdb/app/secret"
	"github.com/ahwhy/yCmdb/app/task"
	aliConn "github.com/ahwhy/yCmdb/provider/aliyun/connectivity"
	ecsOp "github.com/ahwhy/yCmdb/provider/aliyun/ecs"
	hwConn "github.com/ahwhy/yCmdb/provider/huawei/connectivity"
	hwEcsOp "github.com/ahwhy/yCmdb/provider/huawei/ecs"
	txConn "github.com/ahwhy/yCmdb/provider/txyun/connectivity"
	cvmOp "github.com/ahwhy/yCmdb/provider/txyun/cvm"
	vsConn "github.com/ahwhy/yCmdb/provider/vsphere/connectivity"
	vmOp "github.com/ahwhy/yCmdb/provider/vsphere/vm"
)

func (s *service) syncHost(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager host.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun ecs ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := ecsOp.NewEcsOperater(ec)
		req := ecsOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun cvm ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun ecs ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwEcsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	case resource.Vendor_VSPHERE:
		s.log.Debugf("sync vshpere vm ...")
		client := vsConn.NewVsphereClient(secret.Address, secret.ApiKey, secret.ApiSecret)
		ec, err := client.VimClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := vmOp.NewVmOperater(ec)
		// 通过回调直接保存
		err = operater.Query(func(h *host.Host) {
			// 补充管理信息
			h.Base.SecretId = secret.Id
			s.SaveOrUpdateHost(ctx, h, t)
		})
		if err != nil {
			t.Failed(err.Error())
			return
		}
	default:
		t.Failed(fmt.Sprintf("unsuport vendor %s", secret.Vendor))
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

			// 调用host服务保持数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				// 补充管理信息
				target.Base.SecretId = secret.Id
				s.SaveOrUpdateHost(ctx, target, t)
			}
		}
	}
}

// Host主机数据入库
func (s *service) SaveOrUpdateHost(ctx context.Context, ins *host.Host, t *task.Task) {
	h, err := s.host.SaveOrUpdateHost(ctx, ins)
	var detail *task.Record
	if err != nil {
		s.log.Warnf("save host error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.Base.InstanceId, ins.Information.Name, err.Error())
	} else {
		s.log.Debugf("save host %s to db", h.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.Base.InstanceId, ins.Information.Name)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}
