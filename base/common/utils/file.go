package utils

import (
	"errors"
	"os"
)

func CheckFileExists(p string) bool {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
