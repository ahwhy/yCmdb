package syncer

import (
	"github.com/ahwhy/yCmdb/api/pkg/resource"
)

type SyncRequest struct {
	Region       string
	SecretId     string `validate:"required,lte=100"`
	ResourceType resource.Type
}

func NewSyncRequest(secretId string) *SyncRequest {
	return &SyncRequest{
		SecretId: secretId,
	}
}

func (req *SyncRequest) Validate() error {
	return validate.Struct(req)
}

type SyncDetail struct {
	Name      string `json:"name"`       // 资源名称
	IsSuccess bool   `json:"is_success"` // 是否同步成功
	Message   string `json:"message"`    // 同步失败原因
}

type SyncReponse struct {
	TotolSucceed int64         `json:"total_succeed"`
	TotalFailed  int64         `json:"total_failed"`
	Details      []*SyncDetail `json:"details"`
}

func NewSyncReponse() *SyncReponse {
	return &SyncReponse{
		Details: []*SyncDetail{},
	}
}

func (s *SyncReponse) AddFailed(name, message string) {
	s.Details = append(s.Details, &SyncDetail{
		IsSuccess: false,
		Name:      name,
		Message:   message,
	})
	s.TotalFailed++
}

func (s *SyncReponse) AddSucceed(name, message string) {
	s.Details = append(s.Details, &SyncDetail{
		IsSuccess: true,
		Name:      name,
		Message:   message,
	})
	s.TotolSucceed++
}
