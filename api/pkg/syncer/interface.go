package syncer

import "context"

type SecretService interface {
	// 创建Secret
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	// 查询Secret
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)    // 列表查询
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error) // 详情页查询
}

type SyncService interface {
	Sync(context.Context, *SyncRequest) (*SyncReponse, error)
}

type Service interface {
	SecretService
	SyncService
}
