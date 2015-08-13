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
	"unsafe"

	"github.com/blang/semver"
)

func GetString() (string, error) {
	bufferSize := C.size_t(256)
	str := (*C.char)(C.malloc(bufferSize))
	defer C.free(unsafe.Pointer(str))

	err := C.darwin_get_os(str, bufferSize)
	if err == -1 {
		return "", errors.New("Error running sysctl")
	}
	return C.GoString(str), nil
}

func GetSemanticVersion() (semver.Version, error) {

	// 14.x.x  OS X 10.10.x Yosemite
	// 13.x.x  OS X 10.9.x Mavericks
	// 12.x.x  OS X 10.8.x Mountain Lion
	// 11.x.x  OS X 10.7.x Lion
	// 10.x.x  OS X 10.6.x Snow Leopard
	//  9.x.x  OS X 10.5.x Leopard
	//  8.x.x  OS X 10.4.x Tiger
	//  7.x.x  OS X 10.3.x Panther
	//  6.x.x  OS X 10.2.x Jaguar
	//  5.x    OS X 10.1.x Puma

	str, err := GetString()
	if err != nil {
		return semver.Version{}, err
	}

	return semver.Make(str)
}
