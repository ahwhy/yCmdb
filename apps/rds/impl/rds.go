package impl

import (
	"context"

	"github.com/ahwhy/yCmdb/apps/rds"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

func (s *service) SaveRds(ctx context.Context, r *rds.RDS) (*rds.RDS, error) {
	r.Base.Id = xid.New().String()
	r.Base.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, r); err != nil {
		return nil, err
	}

	return r, nil
}
