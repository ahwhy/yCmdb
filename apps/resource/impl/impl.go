package impl

import (
	"github.com/ahwhy/yCmdb/apps"
	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	// db  *sql.DB
	db  *gorm.DB
	log logger.Logger

	resource.UnimplementedServiceServer
}

func (s *service) Config() error {
	// db, err := conf.C().MySQL.GetDB()
	orm, err := conf.C().MySQL.ORM()
	if err != nil {
		return err
	}
	// 在冲突时，更新除主键以外的所有列到新值。
	orm = orm.Clauses(clause.OnConflict{
		UpdateAll: true,
	})
	// 是否开启debug
	if conf.C().Log.Level == "debug" {
		orm.Debug()
	}

	s.log = zap.L().Named(s.Name())
	s.db = db

	return nil
}

func (s *service) Name() string {
	return resource.AppName
}

func (s *service) Registry(server *grpc.Server) {
	resource.RegisterServiceServer(server, svr)
}

func init() {
	apps.RegistryGrpcApp(svr)
	apps.RegistryInternalApp(svr)
}
