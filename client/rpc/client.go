package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/ahwhy/yCmdb/apps/bill"
	"github.com/ahwhy/yCmdb/apps/host"
	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/apps/secret"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

// NewClient todo
func NewClientSet(conf *rpc.Config) (*ClientSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "yCmdb"),
		grpc.WithPerRPCCredentials(rpc.NewAuthentication(conf.ClientID, conf.ClientSecret)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(rpc.NewExceptionUnaryClientInterceptor().UnaryClientInterceptor),
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("sdk.yCmdb"),
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *rpc.Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Instance服务的SDK
func (c *ClientSet) Health() healthgrpc.HealthClient {
	return healthgrpc.NewHealthClient(c.conn)
}

// Resource todo
func (c *ClientSet) Resource() resource.RPCClient {
	return resource.NewRPCClient(c.conn)
}

// Host todos
func (c *ClientSet) Secret() secret.RPCClient {
	return secret.NewRPCClient(c.conn)
}

// Host todos
func (c *ClientSet) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Bill service
func (c *ClientSet) Bill() bill.ServiceClient {
	return bill.NewServiceClient(c.conn)
}

// Rds service
// func (c *ClientSet) Rds() rds.ServiceClient {
// 	return rds.NewServiceClient(c.conn)
// }
