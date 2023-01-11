//go:build windows

package windows

import "github.com/yusufpapurcu/wmi"

// cf.) https://learn.microsoft.com/windows/win32/cimwin32prov/win32-operatingsystem
type Win32_OperatingSystem struct {
	Caption    string
	Version    string
	CSDVersion string
}

type WinVer struct {
	OSName  string
	Version string
	Release string
}

func GetWinVer() (*WinVer, error) {
	var dst []Win32_OperatingSystem
	q := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}
	winVer := WinVer{}
	for _, v := range dst {
		winVer.OSName = v.Caption
		winVer.Version = v.Version
		winVer.Release = v.CSDVersion
		break //nolint
	}
	return &winVer, nil
}
