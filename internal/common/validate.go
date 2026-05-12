package common

import "strings"

// Quick function for checking is a string has URL format
func IsURL(s string) bool {
	parts := strings.Split(s, "://")

	return len(parts) >= 2
}
