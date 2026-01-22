package mobilehttp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblcustomersphoto"
)

type tblcustomersphotoHandler struct {
	uc tblcustomersphoto.UsersUsecase
}

func NewTblcustomersphotoHandler(uc tblcustomersphoto.UsersUsecase) *tblcustomersphotoHandler {
	return &tblcustomersphotoHandler{
		uc: uc,
	}
}

// == Function endpoint ==

func (u *tblcustomersphotoHandler) SaveDataPhotoApi(c *fiber.Ctx) error {
	var customerPhotoReq tblcustomersphoto.CustomersPhotoRequest
	if err := c.BodyParser(&customerPhotoReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Invalid request body", err.Error()),
		)
	}

	err := u.uc.AddCustomersPhoto(&customerPhotoReq)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			domain.ErrorResponse("store data failed", err.Error()),
		)
	}

	return c.Status(http.StatusOK).JSON(
		domain.SuccessResponse("Done", ""),
	)
}

func (u *tblcustomersphotoHandler) SavePhotoToLocalApi(c *fiber.Ctx) error {
	var photoReq tblcustomersphoto.PhotoRequest
	if err := c.BodyParser(&photoReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Invalid request body", err.Error()),
		)
	}

	err := u.uc.SavePhoto(&photoReq)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			domain.ErrorResponse("Save photo failed", err.Error()),
		)
	}

	return c.Status(http.StatusOK).JSON(
		domain.SuccessResponse("Done", ""),
	)
}
