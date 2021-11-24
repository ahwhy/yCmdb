package ecs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/provider/aliyun/connectivity"
	op "github.com/ahwhy/yCmdb/provider/aliyun/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Printf("%v\n", p.Data)
	}
}

func init() {
	var secretID, secretKey string
	if secretID = os.Getenv("AL_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AL_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-zhangjiakou")

	ec, err := client.EcsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewEcsOperater(ec)
}
