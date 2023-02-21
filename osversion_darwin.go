package osversion

/*
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/sysctl.h>

int darwin_get_os(char* str, size_t size) {
    return sysctlbyname("kern.osrelease", str, &size, NULL, 0);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func GetString() (string, error) {
	bufferSize := C.size_t(256)
	str := (*C.char)(C.malloc(bufferSize))
	defer C.free(unsafe.Pointer(str))

	err := C.darwin_get_os(str, bufferSize)
	if err == -1 {
		return "", errors.New(fmt.Sprintf("Error running sysctl: %v", err))
	}
	return C.GoString(str), nil
}

func GetHumanReadable() (string, error) {
	version, err := GetSemanticVersion()
	if err != nil {
		return "", err
	}
	if version.Major < 4 || version.Major > 22 {
		return fmt.Sprintf("Unknown OS X version: %s", version.String()), nil
	}

	decimal := strconv.FormatUint(version.Patch, 10)
	if version.Major >= 18 {
		decimal = strconv.FormatUint(version.Minor-1, 10)
	}
	return strings.Replace(versions[version.Major-4],
		"{decimal}",
		decimal,
		1), nil
}

var versions = []string{
	"OS X 10.0.{decimal} Cheetah",
	"OS X 10.1.{decimal} Puma",
	"OS X 10.2.{decimal} Jaguar",
	"OS X 10.3.{decimal} Panther",
	"OS X 10.4.{decimal} Tiger",
	"OS X 10.5.{decimal} Leopard",
	"OS X 10.6.{decimal} Snow Leopard",
	"OS X 10.7.{decimal} Lion",
	"OS X 10.8.{decimal} Mountain Lion",
	"OS X 10.9.{decimal} Mavericks",
	"OS X 10.10.{decimal} Yosemite",
	"OS X 10.11.{decimal} El Capitan",
	"macOS 10.12.{decimal} Sierra",
	"macOS 10.13.{decimal} High Sierra",
	"macOS 10.14.{decimal} Mojave",
	"macOS 10.15.{decimal} Catalina",
	"macOS 11.{decimal} Big Sur",
	"macOS 12.{decimal} Monterey",
	"macOS 13.{decimal} Ventura",
}
