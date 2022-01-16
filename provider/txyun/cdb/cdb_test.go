package cdb_test

import (
	"fmt"
	"os"
	"testing"

	op "github.com/ahwhy/yCmdb/provider/txyun/cdb"
	"github.com/ahwhy/yCmdb/provider/txyun/connectivity"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

var (
	operater *op.CDBOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

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

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Shanghai)
	operater = op.NewCDBOperater(client.CDBClient())
}
