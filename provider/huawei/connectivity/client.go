package connectivity

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	rds_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/region"
	dcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"
	bss "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
)

// NewHuaweiCloudClient client
func NewHuaweiCloudClient(ak, sk, region string) *HuaweiCloudClient {
	return &HuaweiCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type HuaweiCloudClient struct {
	Region       string
	AccessKey    string
	AccessSecret string
	accountId    string
	ecsConn      *ecs.EcsClient
	rdsConn      *rds.RdsClient
	dcsConn      *dcs.DcsClient
	bssConn      *bss.BssClient
}

func (c *HuaweiCloudClient) Credentials() basic.Credentials {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()

	return auth
}

func (c *HuaweiCloudClient) GlobalCredentials() global.Credentials {
	auth := global.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()

	return auth
}

// EcsClient 客户端
func (c *HuaweiCloudClient) EcsClient() (*ecs.EcsClient, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client := ecs.EcsClientBuilder().
		WithRegion(ecs_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.ecsConn = ecs.NewEcsClient(client)

	return c.ecsConn, nil
}

// RdsClient 客户端
func (c *HuaweiCloudClient) RdsClient() (*rds.RdsClient, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client := rds.RdsClientBuilder().
		WithRegion(rds_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.rdsConn = rds.NewRdsClient(client)

	return c.rdsConn, nil
}

// DcsClient 客户端
func (c *HuaweiCloudClient) DcsClient() (*dcs.DcsClient, error) {
	if c.dcsConn != nil {
		return c.dcsConn, nil
	}

	client := dcs.DcsClientBuilder().
		WithRegion(ecs_region.ValueOf(c.Region)).
		WithCredential(c.Credentials()).
		Build()

	c.dcsConn = dcs.NewDcsClient(client)

	return c.dcsConn, nil
}

// BssClient 客户端
func (c *HuaweiCloudClient) BssClient() (*bss.BssClient, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client := bss.BssClientBuilder().
		WithEndpoint("bss.myhuaweicloud.com").
		WithCredential(c.GlobalCredentials()).
		Build()

	c.bssConn = bss.NewBssClient(client)
	
	return c.bssConn, nil
}