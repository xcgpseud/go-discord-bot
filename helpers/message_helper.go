package helpers

import (
	"os"
	"strings"
)

func IsCommand(s string) bool {
	prefix := os.Getenv("PREFIX")
	return len(s) > len(prefix) && s[0:len(prefix)] == prefix
}

func GetCommandAndContent(s string) (command string, content string) {
	parts := strings.Split(s, " ")
	command = parts[0][len(os.Getenv("PREFIX")):]
	if len(parts) < 2 {
		return
	}
	content = strings.Join(parts[1:], " ")
	return
}
