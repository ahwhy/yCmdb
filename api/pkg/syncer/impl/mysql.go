package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/api/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	// syncer syncer.Service
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	// s.syncer = pkg.Syncer
	s.log = zap.L().Named("Syncer")

	return nil
}
