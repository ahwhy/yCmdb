package protocol

import (
	"net"

	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/conf"

	"github.com/infraboard/mcube/grpc/middleware/recovery"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		rc.UnaryServerInterceptor(),
	)))

	return &GRPCService{
		svr: grpcServer,
		l:   log,
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	app.LoadGrpcApp(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPCAddr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}
