package impl

import (
	"context"
	"fmt"

	"github.com/ahwhy/yCmdb/api/pkg/secret"
)

const (
	insertSecretSQL = `INSERT INTO secret (
		id,create_at,description,vendor,address,allow_regions,crendential_type,api_key,api_secret,request_rate
	) VALUES (?,?,?,?,?,?,?,?,?,?);`

	querySecretSQL = `SELECT * FROM secret`

	deleteSecret = `DELETE FROM secret WHERE id = ?;`
)

func (s *service) createSecret(ctx context.Context, ins *secret.Secret) error {
	if ins == nil {
		return fmt.Errorf("secret is nil")
	}

	stmt, err := s.db.Prepare(insertSecretSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ins.Id, ins.CreateAt, ins.Description, ins.Vendor, ins.Address,
		ins.AllowRegionString(), ins.CrendentialType, ins.APIKey, ins.APISecret,
		ins.RequestRate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) deleteSecret(ctx context.Context, ins *secret.Secret) error {
	if ins == nil {
		return fmt.Errorf("secret is nil")
	}

	stmt, err := s.db.Prepare(deleteSecret)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ins.Id)
	if err != nil {
		return err
	}

	return nil
}
