package rpc_test

import (
	"context"
	"testing"

	"github.com/ahwhy/yCmdb/apps/resource"
	"github.com/ahwhy/yCmdb/conf"
	"github.com/ahwhy/yCmdb/client/rpc"

	"github.com/infraboard/mcenter/apps/health"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	client *rpc.ClientSet
	ctx    = context.Background()
)

func TestClient(t *testing.T) {
	rs, err := client.Resource().Search(ctx, resource.NewSearchRequest())
	if err != nil {
		if e, ok := err.(exception.APIException); ok {
			t.Fatal(e.ToJson())
		} else {
			t.Fatal(err)
		}
	}

	t.Log(rs)
}

func TestHealth(t *testing.T) {
	req := health.NewHealthCheckRequest()
	resp, err := client.Health().Check(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func init() {
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	c, err := rpc.NewClientSet(conf.C().Mcenter)
	if err != nil {
		panic(err)
	}
	client = c
}
