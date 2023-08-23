package uuid

import (
	"github.com/google/uuid"
	"strings"
)

func New() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(), "-", "")
}
