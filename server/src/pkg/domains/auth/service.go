package auth

import (
	"errors"

	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
)

type Service struct {
	accountRepo *account.Repository
	// sessionRepo *session.Repository
}

func NewService(accountRepo *account.Repository) *Service {
	return &Service{
		accountRepo: accountRepo,
		// sessionRepo: sessionRepo,
	}
}

func (svc *Service) Register(n *account.NewEntity) (*account.Entity, error) {
	hash, err := GenerateHash(n.Password)
	if err != nil {
		return nil, err
	}

	n.Password = hash

	created, err := svc.accountRepo.Create(n)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (svc *Service) Login(loginAccount *account.LoginEntity) (string, string, error) {
	acc, err := svc.accountRepo.FindByEmail(loginAccount.Email)
	if err != nil {
		return "", "", err
	}

	if acc == nil {
		return "", "", errors.New("invalid credentials")
	}

	ok, err := CompareRawWithHash(loginAccount.Password, acc.Password)
	if err != nil {
		return "", "", err
	}
	if !ok {
		return "", "", errors.New("invalid credentials")
	}

	// generate access and refresh tokens

	return "access_token", "refresh_token", nil
}
