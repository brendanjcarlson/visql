package auth_test

import (
	"testing"

	"github.com/brendanjcarlson/visql/server/src/pkg/domains/auth"
)

func TestGenerateHash(t *testing.T) {
	t.Run("should generate hash", func(t *testing.T) {
		hash, err := auth.GenerateHash("password")
		if err != nil {
			t.Errorf("expected nil error\ngot: %v\n", err.Error())
		}
		if hash == "" {
			t.Errorf("expected hash to not be empty\ngot: %v\n", hash)
		}
	})
}

func TestCompareRawWithHash(t *testing.T) {
	t.Run("should return true with nil error", func(t *testing.T) {
		hash, err := auth.GenerateHash("password")
		if err != nil {
			t.Errorf("expected nil error while generating original hash\ngot: %v\n", err.Error())
		}

		ok, err := auth.CompareRawWithHash("password", hash)
		if err != nil {
			t.Errorf("expected nil error\ngot: %v\n", err.Error())
		}
		if !ok {
			t.Errorf("expected ok to be true, got %v", ok)
		}
	})

	t.Run("should return false with nil error", func(t *testing.T) {
		hash, err := auth.GenerateHash("password")
		if err != nil {
			t.Errorf("expected nil error while generating original hash\ngot: %v\n", err.Error())
		}

		ok, err := auth.CompareRawWithHash("DEFINITELYTHEWRONGPASSWORD", hash)
		if err != nil {
			t.Errorf("expected nil error\ngot: %v\n", err.Error())
		}
		if ok {
			t.Errorf("expected ok to be false, got %v", ok)
		}
	})
}
