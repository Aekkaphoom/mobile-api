package middlewareAuthtoken

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/authtoken"
)

type middleAuthtokenHandler struct {
	authUc authtoken.AuthtokenUsecase
}

type IMiddleAuthtoken interface {
	TokenLoginCheckExpire() fiber.Handler
}

func MiddleAuthtokenHandler(authUc authtoken.AuthtokenUsecase) IMiddleAuthtoken {
	return &middleAuthtokenHandler{
		authUc: authUc,
	}
}

func (mu *middleAuthtokenHandler) TokenLoginCheckExpire() fiber.Handler {
	return func(c *fiber.Ctx) error {

		eid := c.Params("eid")
		logintoken := c.Params("authtoken")

		var authReq authtoken.AuthToken
		authReq.Eid = eid
		authReq.Token = logintoken

		expire, err := mu.authUc.GetExpireAccess(authReq)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				domain.ErrorResponse("valid auth token error", err.Error()),
			)

		}

		if !expire {
			return c.Status(http.StatusForbidden).JSON(
				domain.ErrorResponse("auth token expired", domain.ErrAuthTokenExpire),
			)

		}

		return c.Next()
	}
}
