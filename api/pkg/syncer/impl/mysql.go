package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/ahwhy/yCmdb/api/conf"
	"github.com/ahwhy/yCmdb/api/pkg"
	"github.com/ahwhy/yCmdb/api/pkg/host"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db   *sql.DB
	log  logger.Logger
	host host.Service
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	s.host = pkg.Host
	s.log = zap.L().Named("Syncer")

	return nil
}
