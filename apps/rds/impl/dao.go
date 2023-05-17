package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ahwhy/yCmdb/apps/rds"
	"github.com/ahwhy/yCmdb/apps/resource/impl"
)

func (s *service) save(ctx context.Context, r *rds.RDS) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事务
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 避免SQL注入, 使用Prepare
	stmt, err = tx.Prepare(impl.SQLInsertResource)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 生成描写信息的Hash
	r.Base.ResourceHash = r.Information.Hash()
	r.Base.DescribeHash = r.Describe.Hash()

	base := r.Base
	info := r.Information
	_, err = stmt.Exec(
		base.Id, base.Vendor, base.Region, base.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type, base.InstanceId,
		info.Name, info.Description, info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount, info.PublicIPToString(),
		info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash, base.SecretId,
	)
	if err != nil {
		return fmt.Errorf("save rds resource info error, %s", err)
	}

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertRdsSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	desc := r.Describe
	_, err = stmt.Exec(
		base.Id, desc.Cpu,
	)
	if err != nil {
		return fmt.Errorf("save rds resource describe error, %s", err)
	}

	return tx.Commit()
}
