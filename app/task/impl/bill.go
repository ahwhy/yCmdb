package impl

import (
	"context"

	"github.com/ahwhy/yCmdb/app/secret"
	"github.com/ahwhy/yCmdb/app/task"
)

func (s *service) syncBill(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
}
