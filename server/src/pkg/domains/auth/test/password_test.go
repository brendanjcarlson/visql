package auth_test

import (
	"fmt"
	"testing"

	"github.com/brendanjcarlson/visql/server/src/pkg/domains/auth"
)

func TestGenerateHash(t *testing.T) {
	t.Run("should generate hash", func(t *testing.T) {
		hash, err := auth.GenerateHash("password")
		if err != nil {
			t.Error("should not return error")
		}
		if hash == "" {
			t.Error("should return hash")
		}
	})
}

func TestCompareRawWithHash(t *testing.T) {
	t.Run("should return true with nil error", func(t *testing.T) {
		hash, _ := auth.GenerateHash("password")
		ok, err := auth.CompareRawWithHash("password", hash)
		if err != nil {
			t.Error("should not return error")
		}
		if !ok {
			t.Error("should return true")
		}
	})

	t.Run("should return false with nil error", func(t *testing.T) {
		hash, _ := auth.GenerateHash("password")
		ok, err := auth.CompareRawWithHash("DEFINITELYTHEWRONGPASSWORD", hash)
		if err != nil {
			fmt.Println(err)
			t.Error("should not return error")
		}
		if ok {
			t.Error("should return false")
		}
	})
}
