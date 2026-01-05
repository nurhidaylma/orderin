package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	domain "github.com/nurhidaylma/orderin/internal/domain/user"
	service "github.com/nurhidaylma/orderin/internal/service/auth"
)

type AuthHandler struct {
	auth service.Service
}

func NewAuthHandler(auth service.Service) *AuthHandler {
	return &AuthHandler{auth: auth}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	var req registerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := h.auth.Register(ctx.Context(),
		req.Name,
		req.Email,
		req.Password,
		domain.Role(req.Role),
	)
	if err != nil {
		return mapAuthError(ctx, err)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
	})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var req loginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.auth.Login(ctx.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		return mapAuthError(ctx, err)
	}

	return ctx.JSON(fiber.Map{
		"access_token": token,
	})
}
