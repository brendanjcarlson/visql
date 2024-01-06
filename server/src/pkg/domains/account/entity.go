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
	FullName    string       `json:"full_name" db:"full_name"`
	Email       string       `json:"email" db:"email"`
	Password    string       `json:"password" db:"password"`
}

type NewEntity struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *NewEntity) Validate() (bool, []error) {
	var errs []error

	if e.FullName == "" {
		errs = append(errs, common.ErrFullNameRequired)
	}

	if e.Email == "" {
		errs = append(errs, common.ErrEmailRequired)
	}
	// todo: validate email format

	if e.Password == "" {
		errs = append(errs, common.ErrPasswordRequired)
	}
	if len(e.Password) < common.MIN_PASSWORD_LENGTH {
		errs = append(errs, common.ErrPasswordTooShort)
	}
	// todo: validate password strength

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

type LoginEntity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *LoginEntity) Validate() (bool, []error) {
	var errs []error

	if e.Email == "" {
		errs = append(errs, common.ErrEmailRequired)
	}

	if e.Password == "" {
		errs = append(errs, common.ErrPasswordRequired)
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
