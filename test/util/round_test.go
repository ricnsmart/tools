package util

import (
	"fmt"
	"github.com/ricnsmart/tools/util"
	"testing"
)

func TestBankerRounding(t *testing.T) {
	result := util.BankerRounding(1.3322, 2)
	fmt.Print(result)
}
