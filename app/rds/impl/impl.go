package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/rds"
	"github.com/ahwhy/yCmdb/conf"
	
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger

	rds.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db

	return nil
}

func (s *service) Name() string {
	return rds.AppName
}

func (s *service) Registry(server *grpc.Server) {
	rds.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
