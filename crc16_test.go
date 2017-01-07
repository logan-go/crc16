package crc16

import (
	"fmt"
	"testing"
)

func TestCrc16(t *testing.T) {
	source := "123456789"
	if fmt.Sprintf("%X", Crc16(source)) != "31C3" {
		t.FailNow()
	}
}
