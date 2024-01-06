package account

import (
	"database/sql"
	"time"

	"github.com/brendanjcarlson/visql/server/src/pkg/domains/common"
	"github.com/google/uuid"
)

type Entity struct {
	Id          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
	LastLoginAt sql.NullTime `json:"last_login_at" db:"last_login_at"`
	LoginCount  int          `json:"login_count" db:"login_count"`
	Email       string       `json:"email" db:"email"`
	Password    string       `json:"password" db:"password"`
}

type NewEntity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginEntity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *NewEntity) Validate() []error {
	var errs []error

	if e.Email == "" {
		errs = append(errs, common.ErrEmailRequired)
	}
	// todo: validate email format

	if e.Password == "" {
		errs = append(errs, common.ErrPasswordRequired)
	}
	if len(e.Password) < 8 {
		errs = append(errs, common.ErrPasswordTooShort)
	}
	// todo: validate password strength

	if len(errs) > 0 {
		return errs
	}
	return nil
}
