package secret

import "context"

type Service interface {
	// 创建Secret
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	// 查询Secret
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)    // 列表查询
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error) // 详情页查询
	// 删除Secret
	DeleteSecret(context.Context, *DeleteSecretRequest) (*Secret, error)
	// 同步资源
	// 	Sync(context.Context, *SyncRequest) (*SyncReponse, error)
}
