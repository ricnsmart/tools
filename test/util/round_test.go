package util

import (
	"fmt"
	"github.com/ricnsmart/tools/util"
	"testing"
)

func TestBankerRounding(t *testing.T) {
	result := util.BankerRounding(0.0000001, 2)
	fmt.Print(result)
}
