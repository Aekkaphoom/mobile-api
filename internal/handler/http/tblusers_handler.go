package mobilehttp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblusers"
	"grouplease.co.th/mobile_api/internal/domain/tbluserslogin"
)

type TblUsersHandler struct {
	uc tblusers.TblUsersUseCase
}

func NewTblUsersHandler(uc tblusers.TblUsersUseCase) *TblUsersHandler {
	return &TblUsersHandler{
		uc: uc,
	}
}

func (h *TblUsersHandler) AddUserApi(c *fiber.Ctx) error {
	var user tblusers.TblUsers
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Invalid request user body", err.Error()),
		)
	}

	err := h.uc.Create(c.Context(), &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			domain.ErrorResponse("create user error", err.Error()),
		)
	}

	return c.Status(http.StatusCreated).JSON(
		domain.SuccessResponse("done", ""),
	)
}

func (h *TblUsersHandler) UpdateSidByEidApi(c *fiber.Ctx) error {

	var req tblusers.UpdateSidRequestResponse
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Invalid request body", err.Error()),
		)
	}

	resp, err := h.uc.UpdateSidByEid(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Update sid failed", err.Error()),
		)
	}

	return c.Status(http.StatusOK).JSON(
		domain.SuccessResponse("Update SID done", resp),
	)

}

func (h *TblUsersHandler) GetSidByEidApi(c *fiber.Ctx) error {
	eid := c.Params("eid")
	if eid == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "EID parameter is required",
		})
	}

	sid, err := h.uc.FindSidByEid(c.Context(), eid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve SID",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "SID retrieved successfully",
		"sid":     sid,
	})
}

func (h *TblUsersHandler) LoginUserApi(c *fiber.Ctx) error {
	var loginReq tbluserslogin.LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := h.uc.FindByEIDAndPass(loginReq.EID, loginReq.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find user",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}
