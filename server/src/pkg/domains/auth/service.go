package auth

import "github.com/brendanjcarlson/visql/server/src/pkg/domains/account"

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

func (svc *Service) Register(newAccount *account.NewEntity) (*account.Entity, error) {
	// todo: hash password
	// todo: send welcome email with email verification link
	return svc.accountRepo.Create(newAccount)
}
