package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/conf"
	"github.com/ahwhy/yCmdb/apps"
	"github.com/ahwhy/yCmdb/apps/host"
	"github.com/ahwhy/yCmdb/apps/secret"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db   *sql.DB
	log  logger.Logger
	host host.ServiceServer
	secret.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	s.host = apps.GetGrpcApp(host.AppName).(host.ServiceServer)

	return nil
}

func (s *service) Name() string {
	return secret.AppName
}

func (s *service) Registry(server *grpc.Server) {
	secret.RegisterServiceServer(server, svr)
}

func init() {
	apps.RegistryGrpcApp(svr)
}
