package mobilehttp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/authtoken"
)

type authtokenHandler struct {
	authUc authtoken.AuthtokenUsecase
}

func NewAuthtokenHandler(authUc authtoken.AuthtokenUsecase) *authtokenHandler {
	return &authtokenHandler{
		authUc: authUc,
	}
}

func (h authtokenHandler) GetAuthExpire(c *fiber.Ctx) error {

	eid := c.Params("eid")
	token := c.Params("authtoken")

	req := &authtoken.AuthToken{
		Token: token,
		Eid:   eid,
	}

	res, err := h.authUc.GetDatetimeLoginExpireByAuth(*req)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(
			domain.ErrorResponse("Check authentication token failed ", err.Error()),
		)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"dateLoginExpire": res,
	})
}
