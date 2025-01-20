package syslite

import (
	"runtime"
)

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

func IsARM64() bool {
	return runtime.GOARCH == "arm64"
}

func IsAMD64() bool {
	return runtime.GOARCH == "amd64"
}

func IsX86_64() bool {
	return runtime.GOARCH == "amd64"
}

func Is386() bool {
	return runtime.GOARCH == "386"
}

func IsARM() bool {
	return runtime.GOARCH == "arm"
}

func IsPPC64() bool {
	return runtime.GOARCH == "ppc64"
}

func IsMIPS() bool {
	return runtime.GOARCH == "mips"
}

func IsDebian() bool {
	return IsDistro("DEBIAN")
}

func IsUbuntu() bool {
	return IsDistro("UBUNTU")
}

func IsFedora() bool {
	return IsDistro("FEDORA")
}

func IsCentOS() bool {
	return IsDistro("CENTOS")
}

func IsRHEL() bool {
	return IsDistro("RHEL")
}

func IsArch() bool {
	return IsDistro("ARCH")
}

func IsAlpine() bool {
	return IsDistro("ALPINE")
}

func IsOpenSUSE() bool {
	return IsDistro("OPENSUSE", "SUSE")
}
