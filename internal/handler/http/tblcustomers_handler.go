package mobilehttp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblcustomers"
)

type tblcustomersHandler struct {
	uc tblcustomers.TblCustomersUsecase
}

func NewCustomersHandler(uc tblcustomers.TblCustomersUsecase) *tblcustomersHandler {
	return &tblcustomersHandler{
		uc: uc,
	}
}

func (h *tblcustomersHandler) FollowUpCustomer(c *fiber.Ctx) error {
	var followCustomer tblcustomers.CustomerRequest
	if err := c.BodyParser(&followCustomer); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed parser request follow customer.",
			"error":   err.Error(),
		})
	}

	err := h.uc.AddFollowupCustomer(followCustomer)
	if err != nil {
		// return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		// 	"message": "Add follow customer failed",
		// 	"error":   err.Error(),
		// })
		return domain.MapError(err)
	}

	return c.Status(http.StatusOK).JSON("")
}
