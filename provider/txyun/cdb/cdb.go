package cdb

import (
	"time"

	"github.com/ahwhy/yCmdb/apps/rds"
	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
)

func NewCDBOperater(client *cdb.Client) *CDBOperater {
	return &CDBOperater{
		client: client,
		log:    zap.L().Named("Tx CDB"),
	}
}

type CDBOperater struct {
	client *cdb.Client
	log    logger.Logger
}

func (o *CDBOperater) transferSet(items []*cdb.InstanceInfo) *rds.RDSSet {
	set := rds.NewRDSSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	
	return set
}

func (o *CDBOperater) transferOne(ins *cdb.InstanceInfo) *rds.RDS {
	r := rds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.Vendor_TENCENT
	b.Region = utils.PtrStrV(ins.Region)
	b.Zone = utils.PtrStrV(ins.Zone)
	b.CreateAt = o.parseTime(utils.PtrStrV(ins.CreateTime))
	b.InstanceId = utils.PtrStrV(ins.InstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(utils.PtrStrV(ins.DeadlineTime))
	info.Name = utils.PtrStrV(ins.InstanceName)
	info.Category = utils.PtrStrV(ins.DeviceType)
	// info.Type = ins.DBInstanceType
	// info.Description = ins.DBInstanceDescription
	// info.Status = ins.DBInstanceStatus
	// info.PayType = ins.PayType

	desc := r.Describe
	desc.EngineType = "MySQL"
	desc.EngineVersion = utils.PtrStrV(ins.EngineVersion)
	// desc.InstanceClass = ins.DBInstanceClass
	// desc.ClassType = ins.DBInstanceClass
	// desc.ExportType = ins.DBInstanceNetType
	// desc.NetworkType = ins.InstanceNetworkType
	// desc.Type = ins.DBInstanceType

	return r
}

func (o *CDBOperater) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
