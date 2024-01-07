package account

import (
	"strings"

	"github.com/brendanjcarlson/visql/server/src/pkg/database"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/common"
	"github.com/google/uuid"
)

type Repository struct {
	client *database.Client
}

func NewRepository(client *database.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Create(n *NewEntity) (*Entity, error) {
	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        INSERT INTO accounts (full_name, email, password)
        VALUES ($1, $2, $3)`,
		n.FullName,
		n.Email,
		n.Password,
	)
	if row.Err() != nil {
		err := row.Err()
		if strings.Contains(err.Error(), "unique") {
			return nil, common.ErrEmailAlreadyInUse
		}
		return nil, err
	}

	var created Entity
	err := row.StructScan(&created)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *Repository) FindByEmail(email string) (*Entity, error) {
	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        SELECT *
        FROM accounts
        WHERE email = $1`,
		email,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var found Entity
	err := row.StructScan(&found)
	if err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *Repository) FindById(id uuid.UUID) (*Entity, error) {
	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        SELECT *
        FROM accounts
        WHERE id = $1`,
		id.String(),
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var found Entity
	err := row.StructScan(&found)
	if err != nil {
		return nil, err
	}

	return &found, nil
}

func (r *Repository) Update(u *Entity) (*Entity, error) {
	o, err := r.FindById(u.Id)
	if err != nil {
		return nil, err
	}

	if u.Email == "" {
		u.Email = o.Email
	}
	if u.Password == "" {
		u.Password = o.Password
	}

	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        UPDATE accounts
        SET email = $1, password = $2
        WHERE id = $3
        RETURNING *`,
		u.Email,
		u.Password,
		u.Id.String(),
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var updated Entity
	err = row.StructScan(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *Repository) UpdateOnLogin(id uuid.UUID) (*Entity, error) {
	o, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        UPDATE accounts
        SET last_login_at = CURRENT_TIMESTAMP, login_count = $1
        WHERE id = $2
        RETURNING *`,
		o.LoginCount+1,
		id.String(),
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var updated Entity
	err = row.StructScan(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *Repository) Delete(id uuid.UUID) (*Entity, error) {
	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(`
        DELETE FROM accounts
        WHERE id = $1
        RETURNING *`,
		id.String(),
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var deleted Entity
	err := row.StructScan(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}
