package patient

import (
	"github.com/gofiber/fiber/v3"

	"github.com/HemanthGowda/hospital-management-go/internal/auth"
	"github.com/HemanthGowda/hospital-management-go/internal/domain"
	"github.com/HemanthGowda/hospital-management-go/internal/httpx"
	"github.com/HemanthGowda/hospital-management-go/internal/user"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Update(c fiber.Ctx) error {
	principal, ok := auth.PrincipalFrom(c.Context())
	if !ok {
		return domain.Unauthorized("Authentication is required")
	}
	if !principal.HasAny(user.RoleAdmin, user.RoleDoctor) {
		return domain.Forbidden()
	}

	var request UpdatePatientRequest
	if err := httpx.Decode(c, &request); err != nil {
		return err
	}

	response, err := h.service.Update(c.Context(), c.Params("id"), request)
	if err != nil {
		return err
	}
	return httpx.WriteJSON(c, fiber.StatusOK, response)
}
