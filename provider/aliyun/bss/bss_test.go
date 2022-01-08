package bss_test

import (
	"fmt"
	"os"
	"testing"

	op "github.com/ahwhy/yCmdb/provider/aliyun/bss"
	"github.com/ahwhy/yCmdb/provider/aliyun/connectivity"
)

var (
	operater *op.BssOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2021-12"

	pager := operater.PageQuery(req)
	hasNext := true
	for hasNext {
		p := pager.Next()
		if p.Err != nil {
			panic(p.Err)
		}
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

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-zhangjiakou")

	ec, err := client.BssClient()
	if err != nil {
		panic(err)
	}
	operater = op.NewBssOperater(ec)
}
