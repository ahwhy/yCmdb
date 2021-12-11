package client

import (
	"github.com/ahwhy/yCmdb/app/host"
	"github.com/ahwhy/yCmdb/app/resource"
	"github.com/ahwhy/yCmdb/app/secret"

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

// Resource todo
func (c *Client) Resource() resource.ServiceClient {
	return resource.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Secret() secret.ServiceClient {
	return secret.NewServiceClient(c.conn)
}
