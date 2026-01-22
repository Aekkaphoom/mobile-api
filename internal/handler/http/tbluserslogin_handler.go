package mobilehttp

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"grouplease.co.th/mobile_api/internal/domain"
	"grouplease.co.th/mobile_api/internal/domain/tblusers"
	"grouplease.co.th/mobile_api/internal/domain/tbluserslogin"
	"grouplease.co.th/mobile_api/pkg/utils"
	"grouplease.co.th/mobile_api/pkg/utils/dateutils"
)

type usersLoginHandler struct {
	user_uc  tblusers.TblUsersUseCase
	login_uc tbluserslogin.TblUsersLoginUseCase
}

func NewUsersLoginHandler(user_uc tblusers.TblUsersUseCase, login_uc tbluserslogin.TblUsersLoginUseCase) *usersLoginHandler {
	return &usersLoginHandler{
		user_uc:  user_uc,
		login_uc: login_uc,
	}
}

// ==== Statement Function ====

func (h *usersLoginHandler) LoginApi(c *fiber.Ctx) error {
	// == Parse request body & param ==
	var req tbluserslogin.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			domain.ErrorResponse("Invalid request body", err.Error()),
		)
	}
	role := c.Params("role")

	chkLogin, err := h.login_uc.GetLogin(req.EID, req.Password)
	if err != nil {
		var msgError string
		if err.Error() == "sql: no rows in result set" {
			msgError = "Not found, user/pass wrong !!"
		} else {
			msgError = err.Error()
		}

		return c.Status(http.StatusInternalServerError).JSON(
			domain.ErrorResponse("Failed to check user login already.", msgError),
		)
	}

	if chkLogin != nil {

		// === User login already ===
		// return c.Status(http.StatusUnauthorized).JSON(
		// 	domain.ErrorResponse("User already logged in", "User is already logged in"),
		// )

		// === Return result login ===
		resp, err := h.login_uc.GetUserLogin(req.EID, chkLogin.AuthToken)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				domain.ErrorResponse("Get user login failed", err.Error()),
			)
		}

		return c.Status(http.StatusOK).JSON(
			domain.SuccessResponse("login successfully", resp),
		)

	} else {

		// == Validate User  ==
		_, err := h.user_uc.FindByEIDAndPass(req.EID, req.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				domain.ErrorResponse("Failed to find user", err.Error()),
			)
		}

		// ==== Process Login ====
		now := time.Now()
		location, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			panic(err)
		}
		thaiTime := now.In(location)

		var session_time time.Time
		if role == "mobile" {
			session_time = dateutils.New().ConvertDateStringToSQLDate(thaiTime.AddDate(0, 0, 7))
		} else {
			session_time = thaiTime.Add(time.Hour * 8)
		}

		// == For test add 4 hours ==
		session_time = thaiTime.Add(time.Hour * 4)

		userLogin := &tbluserslogin.TblUserLoginRequest{
			AuthToken:       utils.GenerateAuthToken(),
			UserEid:         req.EID,
			RoleAccess:      role + "_user",
			LoginDatetime:   dateutils.New("2006-01-02 15:04:05 ").ConvertDateStringToSQLDate(time.Now()),
			LoginDateExpire: session_time,
			CreatedBy:       "user-" + role,
			UpdatedBy:       "BE-API",
		}

		loginErr := h.login_uc.AddUserLogin(userLogin)
		if loginErr != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to login user",
				"error":   loginErr.Error(),
			})
		}

		// === Return login result ===
		resp, err := h.login_uc.GetUserLogin(req.EID, userLogin.AuthToken)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				domain.ErrorResponse("Get user login failed", err.Error()),
			)
		}

		return c.Status(http.StatusOK).JSON(
			domain.SuccessResponse("login successfully", resp),
			// domain.SuccessResponse("login successfully", userLogin),
			// domain.SuccessResponse("login successfully", nil),
		)

	}

}

func (h *usersLoginHandler) GetUserLoginOnline(c *fiber.Ctx) error {
	eid := c.Params("eid")
	authToken := c.Params("authtoken")

	userlogin, err := h.login_uc.GetUserLogin(eid, authToken)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			domain.ErrorResponse("Get user login failed", err.Error()),
		)
	}

	return c.Status(http.StatusOK).JSON(
		domain.SuccessResponse("login successfully", userlogin),
	)
}

func (h *usersLoginHandler) LogoutUserApi(c *fiber.Ctx) error {
	// == Parse request body ==
	userEid := c.Params("eid")

	logoutErr := h.login_uc.DeleteUserLoginByEid(userEid)
	if logoutErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to logout user",
			"error":   logoutErr.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON("")
}
