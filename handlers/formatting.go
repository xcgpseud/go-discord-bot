package handlers

import (
	"os"
	"strings"
)

func GetParts(s string) (maybePrefix string, command string) {
	prefix := os.Getenv("PREFIX")
	prefixLen := len(prefix)
	maybePrefix = s[0:prefixLen]
	content := s[prefixLen:]

	parts := strings.Split(content, " ")
	command = parts[0]

	return
}

func GetFollowing(s string) (following string) {
	parts := strings.Split(s, " ")
	following = strings.Join(parts[1:], " ")

	return
}
