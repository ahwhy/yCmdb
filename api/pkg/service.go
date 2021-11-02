package pkg

import (
	"github.com/ahwhy/yCmdb/api/pkg/host"
	"github.com/ahwhy/yCmdb/api/pkg/resource"
	"github.com/ahwhy/yCmdb/api/pkg/syncer"
)

var (
	Host     host.Service
	Syncer   syncer.Service
	Resource resource.Service
)
