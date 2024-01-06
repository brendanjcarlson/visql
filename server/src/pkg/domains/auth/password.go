package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	SALT_LENGTH uint32 = 16
	ITERATIONS  uint32 = 1
	MEMORY      uint32 = 64 * 1024
	THREADS     uint8  = 1
	KEY_LENGTH  uint32 = 32
)

func GenerateHash(raw string) (string, error) {
	salt := make([]byte, SALT_LENGTH)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(raw), salt, ITERATIONS, MEMORY, THREADS, KEY_LENGTH)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	hashString := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, MEMORY, ITERATIONS, THREADS, b64Salt, b64Hash)

	return hashString, nil
}

func CompareRawWithHash(raw string, encodedHash string) (bool, error) {
	salt, hash, memory, iterations, threads, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	comparisonHash := argon2.IDKey([]byte(raw), salt, iterations, memory, threads, uint32(len(hash)))

	if subtle.ConstantTimeCompare(hash, comparisonHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(hash string) (salt []byte, key []byte, memory uint32, iterations uint32, threads uint8, err error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return nil, nil, 0, 0, 0, errors.New("malformed hash")
	}

	var version int
	fmt.Sscanf(parts[2], "v=%d", &version)
	if version != argon2.Version {
		return nil, nil, 0, 0, 0, errors.New("incompatible version of argon2")
	}

	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &threads)

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, 0, 0, 0, err
	}

	key, err = base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, 0, 0, 0, err
	}

	return salt, key, memory, iterations, threads, nil
}
