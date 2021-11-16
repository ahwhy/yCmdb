package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/api/conf"
	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/host"
	"github.com/ahwhy/yCmdb/api/pkg/secret"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db     *sql.DB
	log    logger.Logger
	host   host.Service
	secret secret.Service
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named("Task")
	s.db = db
	s.host = pkg.Host
	s.secret = pkg.Secret

	return nil
}
