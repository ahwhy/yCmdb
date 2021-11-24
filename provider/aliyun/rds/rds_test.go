package rds_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/provider/aliyun/connectivity"
	op "github.com/ahwhy/yCmdb/provider/aliyun/rds"
)

var (
	operater *op.RdsOperater
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
	if secretID = os.Getenv("AL_CLOUD_ACCESS_KEY"); secretID == "" {
		panic("empty AL_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AL_CLOUD_ACCESS_SECRET"); secretKey == "" {
		panic("empty AL_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-hangzhou")

	ec, err := client.RdsClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewRdsOperater(ec)
}
