package billing_test

import (
	"fmt"
	"os"
	"testing"

	op "github.com/ahwhy/yCmdb/provider/txyun/billing"
	"github.com/ahwhy/yCmdb/provider/txyun/connectivity"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

var (
	operater *op.BillingOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	req.Month = "2021-10"

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
	zap.DevelopmentSetup()

	var secretID, secretKey string
	if secretID = os.Getenv("TX_CLOUD_SECRET_ID"); secretID == "" {
		panic("empty TX_CLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TX_CLOUD_SECRET_KEY"); secretKey == "" {
		panic("empty TX_CLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Shanghai)
	operater = op.NewBillingOperater(client.BillingClient())
}
