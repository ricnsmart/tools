package util

import (
	"fmt"
	"testing"
)

func TestBankerRounding(t *testing.T) {
	result := BankerRounding(0.0000001, 2)
	fmt.Print(result)
}
