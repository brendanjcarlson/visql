package auth

import (
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/common"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	app.Post("/api/auth/register", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/login", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/logout", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/forgot-password", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/reset-password", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/verify-email", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/send-verification-email", common.NOT_YET_IMPLEMENTED)
	app.Post("/api/auth/send-password-reset-email", common.NOT_YET_IMPLEMENTED)
}
