package impl

import (
	"database/sql"

	"github.com/ahwhy/yCmdb/api/conf"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type service struct {
	db  *sql.DB
	log logger.Logger
}

var (
	// Service 服务实例
	Service = &service{}
)

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	s.log = zap.L().Named("Host")

	return nil
}
