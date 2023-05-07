package client

import (
	"github.com/ahwhy/yCmdb/apps/bill"
	"github.com/ahwhy/yCmdb/apps/host"
	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/apps/secret"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

// NewClient todo
func NewClient(conf *Config) (*Client, error) {
	zap.DevelopmentSetup()

	conn, err := grpc.Dial(
		conf.address,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(conf.Authentication),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("CMDB SDK"),
	}, nil
}

// Client 客户端
type Client struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Resource service
func (c *Client) Resource() resource.ServiceClient {
	return resource.NewServiceClient(c.conn)
}

// Host service
func (c *Client) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Secret service
func (c *Client) Secret() secret.ServiceClient {
	return secret.NewServiceClient(c.conn)
}

// Bill service
func (c *Client) Bill() bill.ServiceClient {
	return bill.NewServiceClient(c.conn)
}

// Rds service
// func (c *Client) Rds() rds.ServiceClient {
// 	return rds.NewServiceClient(c.conn)
// }
