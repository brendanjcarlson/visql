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

func (r *Repository) Create(e *NewEntity) (*Entity, error) {
	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(
		`
        INSERT INTO accounts (email, password)
        VALUES ($1, $2)
        `,
		e.Email,
		e.Password,
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

	row := r.client.DB().QueryRowx(
		`
        SELECT *
        FROM accounts
        WHERE email = $1
        `,
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

	row := r.client.DB().QueryRowx(
		`
        SELECT *
        FROM accounts
        WHERE id = $1
        `,
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

func (r *Repository) Update(n *Entity) (*Entity, error) {
	o, err := r.FindById(n.Id)
	if err != nil {
		return nil, err
	}

	if n.Email == "" {
		n.Email = o.Email
	}
	if n.Password == "" {
		n.Password = o.Password
	}

	_, cancel := common.NewQueryContext()
	defer cancel()

	row := r.client.DB().QueryRowx(
		`
        UPDATE accounts
        SET email = $1, password = $2
        WHERE id = $3
        RETURNING *
        `,
		n.Email,
		n.Password,
		n.Id.String(),
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

	row := r.client.DB().QueryRowx(
		`
        DELETE FROM accounts
        WHERE id = $1
        RETURNING *
        `,
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
