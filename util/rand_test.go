package util

import (
	"github.com/labstack/gommon/log"
	"testing"
)

func TestNewLen(t *testing.T) {
	password := NewLen(16)

	log.Print(password)

}
