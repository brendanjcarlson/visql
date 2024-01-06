package auth

import (
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/account"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/common"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctrl *Controller) Register(c *fiber.Ctx) error {
	var newAccount account.NewEntity
	err := c.BodyParser(&newAccount)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"message": "invalid request body",
			"errors":  common.StringifyErrs(err),
		})
	}

	errs := newAccount.Validate()
	if errs != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"message": "invalid request body",
			"errors":  common.StringifyErrs(errs...),
		})
	}

	createdAccount, err := ctrl.svc.Register(&newAccount)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&fiber.Map{
			"message": "internal server error",
			"errors":  common.StringifyErrs(err),
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(&fiber.Map{
		"message": "account created",
		"account": &fiber.Map{
			"id": createdAccount.Id,
		},
	})
}

func (ctrl *Controller) Login(c *fiber.Ctx) error {
	return common.ErrNotYetImplemented
}

func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	return common.ErrNotYetImplemented
}
