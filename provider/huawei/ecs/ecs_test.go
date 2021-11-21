package ecs_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/api/provider/huawei/connectivity"
	op "github.com/ahwhy/yCmdb/api/provider/huawei/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}
}

func init() {
	var secretID, secretKey string
	if secretID = os.Getenv("HW_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty HW_CLOUD_ACCESS_KEY")
	}
	
	if secretKey = os.Getenv("HW_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty HW_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewHuaweiCloudClient(secretID, secretKey, "cn-north-4")

	ec, err := client.EcsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewEcsOperater(ec)
}
