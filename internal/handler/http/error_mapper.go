package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	domain "github.com/nurhidaylma/orderin/internal/domain/user"
)

func mapAuthError(ctx *fiber.Ctx, err error) error {
	switch err {
	case domain.ErrEmailExists:
		return ctx.Status(http.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	case domain.ErrWrongPassword:
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	case domain.ErrInvalidRole:
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	default:
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
}
