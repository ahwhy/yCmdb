package connectivity_test

import (
	"os"
	"testing"

	"github.com/ahwhy/yCmdb/api/provider/huawei/connectivity"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("HW_CLOUD_ACCESS_KEY"); secretID == "" {
		t.Fatal("empty HW_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("HW_CLOUD_ACCESS_SECRET"); secretKey == "" {
		t.Fatal("empty HW_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewHuaweiCloudClient(secretID, secretKey, "cn-hangzhou")
	client.EcsClient()
}
