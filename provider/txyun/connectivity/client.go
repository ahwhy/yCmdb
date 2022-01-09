package connectivity

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
)

// NewTencentCloudClient client
func NewTencentCloudClient(secretID, secretKey, region string) *TencentCloudClient {
	return &TencentCloudClient{
		Region:    region,
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}

// TencentCloudClient client for all TencentCloud service
type TencentCloudClient struct {
	Region    string
	SecretID  string
	SecretKey string
	cvmConn   *cvm.Client
	cdbConn   *cdb.Client
	billConn  *billing.Client
}

// UseCvmClient cvm客户端
func (me *TencentCloudClient) CvmClient() *cvm.Client {
	if me.cvmConn != nil {
		return me.cvmConn
	}

	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	cvmConn, _ := cvm.NewClient(credential, me.Region, cpf)
	me.cvmConn = cvmConn
	
	return me.cvmConn
}

// CDBClient cdb客户端
func (me *TencentCloudClient) CDBClient() *cdb.Client {
	if me.cdbConn != nil {
		return me.cdbConn
	}

	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	cdbConn, _ := cdb.NewClient(credential, me.Region, cpf)
	me.cdbConn = cdbConn

	return me.cdbConn
}

// UseBillingClient billing客户端
func (me *TencentCloudClient) BillingClient() *billing.Client {
	if me.billConn != nil {
		return me.billConn
	}
	credential := common.NewCredential(
		me.SecretID,
		me.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 300
	cpf.Language = "en-US"

	billConn, _ := billing.NewClient(credential, me.Region, cpf)
	me.billConn = billConn

	return me.billConn
}