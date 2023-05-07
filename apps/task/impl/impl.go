package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/app/bill"
	"github.com/ahwhy/yCmdb/app/host"
	"github.com/ahwhy/yCmdb/app/rds"
	"github.com/ahwhy/yCmdb/app/secret"
	"github.com/ahwhy/yCmdb/app/task"
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
	task.UnimplementedServiceServer

	host   host.ServiceServer
	rds    rds.ServiceServer
	secret secret.ServiceServer
	bill   bill.ServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	s.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	s.rds = app.GetGrpcApp(rds.AppName).(rds.ServiceServer)
	s.secret = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	s.bill = app.GetGrpcApp(bill.AppName).(bill.ServiceServer)

	return nil
}

func (s *service) Name() string {
	return task.AppName
}

func (s *service) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
