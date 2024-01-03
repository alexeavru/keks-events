//go:build tools
// +build tools

package tools

import (
	"time"

	_ "github.com/99designs/gqlgen"
)

func ParseTime(value string) time.Time {
	t, _ := time.Parse(time.RFC3339, value)
	return t
}
