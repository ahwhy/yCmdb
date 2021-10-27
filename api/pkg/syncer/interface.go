package syncer

import "context"

type SecretService interface {
	// 创建Secret
	CreateSecret(context.Context, *CreateSecretRequest) (*Secret, error)
	// 查询Secret
	QuerySecret(context.Context, *QuerySecretRequest) (*SecretSet, error)
	DescribeSecret(context.Context, *DescribeSecretRequest) (*Secret, error)
}

type SyncService interface {
	// 	Sync(context.Context, *SyncRequest) (*SyncReponse, error)
}

type Service interface {
	SecretService
	SyncService
}
