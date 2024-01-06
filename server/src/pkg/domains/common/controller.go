package common

import (
	"github.com/gofiber/fiber/v2"
)

func NOT_YET_IMPLEMENTED(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotImplemented)

	res := c.JSON(&fiber.Map{
		"error": ErrNotYetImplemented.Error(),
	})

	return res
}
