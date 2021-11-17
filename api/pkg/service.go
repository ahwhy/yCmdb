package pkg

import (
	"github.com/ahwhy/yCmdb/api/pkg/host"
	"github.com/ahwhy/yCmdb/api/pkg/resource"
	"github.com/ahwhy/yCmdb/api/pkg/secret"
	"github.com/ahwhy/yCmdb/api/pkg/task"
)

var (
	Resource resource.Service
	Host     host.Service
	Secret   secret.Service
	Task     task.Service
)
