package ecs_test

import (
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/api/provider/aliyun/connectivity"
	op "github.com/ahwhy/yCmdb/api/provider/aliyun/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T)

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
