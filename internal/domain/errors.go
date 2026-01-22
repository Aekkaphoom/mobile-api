package domain

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	// == User
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	ErrInvalidInput       = errors.New("invalid input")
	ErrSIDExists          = errors.New("SID already exists")
	// == Customer followup
	ErrInvalidFollowUpRequest = errors.New("invalid followup request")
	// == Customer photo
	ErrPhotoRequest = errors.New("invalid photo request")
	// == Other
	ErrNotFound       = errors.New("not found")
	ErrRequestMissing = errors.New("request data missing")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInternal       = errors.New("internal server error")
	// == Authenticate
	ErrAuthTokenMissing = errors.New("Auth token is missing")
	ErrAuthTokenExpire  = errors.New("Token login timeout is expire")
)

func MapError(err error) error {
	switch err {
	case ErrNotFound:
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	case ErrInvalidInput:
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	case ErrAuthTokenMissing:
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	case ErrAuthTokenExpire:
		return fiber.NewError(440, err.Error())
	case ErrInvalidFollowUpRequest:
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}
}

//=== Manage response result ===

type APIResponse struct {
	Success bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
