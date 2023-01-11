//go:build windows

package windows

import (
	"testing"
)

func TestGetWinVer(t *testing.T) {
	winVer, err := GetWinVer()
	if err != nil {
		t.Errorf("should not return error on Windows: %s", err)
	}

	if len(winVer.OSName) <= 0 {
		t.Error("OSName should not be empty")
	}
}
