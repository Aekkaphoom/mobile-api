package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateAuthToken() string {
	tokenuuid := uuid.NewString()

	return strings.ReplaceAll(tokenuuid, "-", "")
}
