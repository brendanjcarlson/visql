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
			"message": common.MsgInvalidRequestBody,
			"errors":  common.StringifyErrs(err),
		})
	}

	valid, errs := newAccount.Validate()
	if !valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"message": common.MsgInvalidRequestBody,
			"errors":  common.StringifyErrs(errs...),
		})
	}

	created, err := ctrl.svc.Register(&newAccount)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&fiber.Map{
			"message": common.MsgInternalServerErr,
			"errors":  common.StringifyErrs(err),
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(&fiber.Map{
		"message": "account created",
		"account": &fiber.Map{
			"id": created.Id,
		},
	})
}

func (ctrl *Controller) Login(c *fiber.Ctx) error {
	var loginEntity account.LoginEntity
	err := c.BodyParser(&loginEntity)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"message": common.MsgInvalidRequestBody,
			"errors":  common.StringifyErrs(err),
		})
	}

	valid, errs := loginEntity.Validate()
	if !valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(&fiber.Map{
			"message": common.MsgInvalidRequestBody,
			"errors":  common.StringifyErrs(errs...),
		})
	}

	accessToken, refreshToken, err := ctrl.svc.Login(&loginEntity)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&fiber.Map{
			"message": common.MsgInternalServerErr,
			"errors":  common.StringifyErrs(err),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(&fiber.Map{
		"message":       "login successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	return common.ErrNotYetImplemented
}
