package auth

import (
	"errors"
	"time"

	"github.com/brendanjcarlson/visql/server/src/pkg/config"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(acc *account.Entity) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:  config.MustGet("JWT_ISSUER"),
		Subject: acc.Id.String(),
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Minute * 30),
		},
		NotBefore: &jwt.NumericDate{
			Time: time.Now(),
		},
		IssuedAt: &jwt.NumericDate{
			Time: time.Now(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.MustGet("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(acc *account.Entity) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:  config.MustGet("JWT_ISSUER"),
		Subject: acc.Id.String(),
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * 24 * 30),
		},
		NotBefore: &jwt.NumericDate{
			Time: time.Now(),
		},
		IssuedAt: &jwt.NumericDate{
			Time: time.Now(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.MustGet("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", err
	}

	hashedTokenString, err := GenerateHash(tokenString)
	if err != nil {
		return "", err
	}

	return hashedTokenString, nil
}

func ExtractAccessClaims(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MustGet("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}
