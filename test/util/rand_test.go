package util

import (
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"testing"
)

func TestNewLen(t *testing.T) {
	password := util.NewLen(16)

	log.Print(password)

}
