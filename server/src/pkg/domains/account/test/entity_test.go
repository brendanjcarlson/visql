package account_test

import (
	"testing"

	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
)

func TestNewEntityValidate(t *testing.T) {
	t.Run("should return true with nil errors for a ok entity", func(t *testing.T) {
		good := &account.NewEntity{
			FullName: "Foo Bar",
			Email:    "foo@bar.baz",
			Password: "foobarbaz",
		}
		ok, errs := good.Validate()
		if !ok {
			t.Errorf("expected ok to be true, got %v", ok)
		}
		if errs != nil {
			t.Errorf("expected errs to be nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for missing full name", func(t *testing.T) {
		missingFullName := &account.NewEntity{
			Email:    "foo@bar.baz",
			Password: "foobarbaz",
		}
		ok, errs := missingFullName.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for missing email", func(t *testing.T) {
		missingEmail := &account.NewEntity{
			FullName: "Foo Bar",
			Password: "foobarbaz",
		}
		ok, errs := missingEmail.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for missing password", func(t *testing.T) {
		missingPassword := &account.NewEntity{
			FullName: "Foo Bar",
			Email:    "foo@bar.baz",
		}
		ok, errs := missingPassword.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for short password", func(t *testing.T) {
		shortPassword := &account.NewEntity{
			FullName: "Foo Bar",
			Email:    "foo@bar.baz",
			Password: "foo",
		}
		ok, errs := shortPassword.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})
}

func TestLoginEntityValidate(t *testing.T) {
	t.Run("should return true with nil errors for a ok entity", func(t *testing.T) {
		good := &account.LoginEntity{
			Email:    "foo@bar.baz",
			Password: "foobarbaz",
		}
		ok, errs := good.Validate()
		if !ok {
			t.Errorf("expected ok to be true, got %v", ok)
		}
		if errs != nil {
			t.Errorf("expected errs to be nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for missing email", func(t *testing.T) {
		missingEmail := &account.LoginEntity{
			Password: "foobarbaz",
		}
		ok, errs := missingEmail.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})

	t.Run("should return false with errors for missing password", func(t *testing.T) {
		missingPassword := &account.LoginEntity{
			Email: "foo@bar.baz",
		}
		ok, errs := missingPassword.Validate()
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
		if errs == nil {
			t.Errorf("expected errs to be non-nil, got %v", errs)
		}
	})
}
