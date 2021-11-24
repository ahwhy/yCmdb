package connectivity_test

import (
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/provider/txyun/connectivity"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("TX_CLOUD_SECRET_ID"); secretID == "" {
		t.Fatal("empty TX_CLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TX_CLOUD_SECRET_KEY"); secretKey == "" {
		t.Fatal("empty TX_CLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Shanghai)
	client.CvmClient()
}
