// This file implements GlowGetProcAddress for every supported platform. The
// correct version is chosen automatically based on build tags:
// windows: WGL
// darwin: CGL
// linux: GLX
// Use of EGL instead of the platform's default (listed above) is made possible
// via the "egl" build tag.
// It is also possible to install your own function outside this package for
// retrieving OpenGL function pointers, to do this see InitWithProcAddrFunc.
package gl

/*
#cgo windows CFLAGS: -DTAG_WINDOWS
#cgo darwin CFLAGS: -DTAG_DARWIN
#cgo linux CFLAGS: -DTAG_LINUX

#cgo linux LDFLAGS: -lOSMesa -ldl

#if defined(TAG_WINDOWS)
	#define WIN32_LEAN_AND_MEAN 1
	#include <windows.h>
	#include <stdlib.h>
	static HMODULE osmesadll = NULL;
	void* GlowGetProcAddress(const char* name) {
		if (osmesadll == NULL) {
			osmesadll = LoadLibraryA("osmesa.dll");
		}
		return GetProcAddress(osmesadll, (LPCSTR) name);
	}
#else
	#define _GNU_SOURCE
	#include <stdlib.h>
	#include <dlfcn.h>
	void* GlowGetProcAddress(const char* name) {
		return dlsym(RTLD_DEFAULT, name);
	}
#endif
*/
import "C"
import "unsafe"

func getProcAddress(namea string) unsafe.Pointer {
	cname := C.CString(namea)
	defer C.free(unsafe.Pointer(cname))
	return C.GlowGetProcAddress(cname)
}
