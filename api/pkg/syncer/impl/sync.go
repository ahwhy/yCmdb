package impl

import (
	"context"

	"github.com/ahwhy/yCmdb/api/pkg/resource"
	"github.com/ahwhy/yCmdb/api/pkg/syncer"
	"github.com/infraboard/mcube/exception"
)

func (s *service) Sync(ctx context.Context, req *syncer.SyncRequest) (*syncer.SyncReponse, error) {
	var (
		resp *syncer.SyncReponse
		err  error
	)

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate sync request error, %s", err)
	}

	secret, err := s.DescribeSecret(ctx, syncer.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}

	// 如果不是vsphere 需要检查region
	if !secret.Vendor.Equal(resource.VendorVsphere) {
		if req.Region == "" {
			return nil, exception.NewBadRequest("region required")
		}
		if !secret.IsAllowRegion(req.Region) {
			return nil, exception.NewBadRequest("this secret not allow sync region %s", req.Region)
		}
	}

	switch req.ResourceType {
	case resource.HostResource:
		resp, err = s.syncHost(ctx, secret, req.Region)
	case resource.RdsResource:
		// resp, err = s.syncRds(ctx, secret, req.Region)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
