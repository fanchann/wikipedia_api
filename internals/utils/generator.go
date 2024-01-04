package utils

import (
	"strings"

	"github.com/google/uuid"
)

func IDGenerator() string {
	id := uuid.NewString()
	idStr := strings.Replace(id, "-", "", 16)
	return idStr
}
