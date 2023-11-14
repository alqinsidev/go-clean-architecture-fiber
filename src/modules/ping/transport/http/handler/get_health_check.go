package handler

import (
	"alqinsidev.io/go-clean-architecture/src/constant"
	httpHelpers "alqinsidev.io/go-clean-architecture/src/helpers/http"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetHealthCheck(c *fiber.Ctx) error {
	result, err := h.usecase.GetHealthCheck()
	if err != nil {
		return httpHelpers.RenderError(c, fiber.StatusInternalServerError, err)
	}

	h.app.GetLogger().LogSuccess(constant.LogCategoryUsecase, MethodGetHealthCheck, "success")
	return httpHelpers.Render(c, fiber.StatusOK, result, nil)
}
