package main

import (
	"fmt"
	"log/syslog"
	"os"
	"unsafe"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

char *string_from_argv(int, char**);
char *get_user(pam_handle_t *pamh);
int get_uid(char *user);
*/
import "C"

func init() {
	if !disablePtrace() {
	}
}

func log(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-demo")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

func sliceFromArgv(argc C.int, argv **C.char) []string {
	r := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {
		s := C.string_from_argv(C.int(i), argv)
		defer C.free(unsafe.Pointer(s))
		r = append(r, C.GoString(s))
	}
	return r
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	cUsername := C.get_user(pamh)
	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	cPassword := C.get_password(pamh)
	if cPassword != nil {
		defer C.free(unsafe.Pointer(cPassword))
	}

	uid := int(C.get_uid(cUsername))
	if uid < 0 {
		return C.PAM_USER_UNKNOWN
	}
	log("pam_sm_authenticate: user: %v, password: %v flags: %x, slice: %v",
		C.GoString(cUsername), C.GoString(cPassword), flags, sliceFromArgv(argc, argv))
	r := pamAuthenticate(os.Stderr, uid, C.GoString(cUsername), sliceFromArgv(argc, argv))
	if r == AuthError {
		return C.PAM_AUTH_ERR
	}
	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	log("pam_sm_setcred: user: %v, flags: %x, slice: %v", C.GoString(C.get_user(pamh)), flags, sliceFromArgv(argc, argv))
	return C.PAM_IGNORE
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	log("pam_sm_acct_mgmt: user: %v, flags: %x, slice: %v", C.GoString(C.get_user(pamh)), flags, sliceFromArgv(argc, argv))
	// PAM_ACCT_EXPIRED
	// PAM_AUTH_ERR
	// PAM_NEW_AUTHTOK_REQD
	// PAM_PERM_DENIED
	// PAM_USER_UNKNOWN
	return C.PAM_SUCCESS
}

//export pam_sm_open_session
func pam_sm_open_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	log("pam_sm_open_session: user: %v, flags: %x, slice: %v", C.GoString(C.get_user(pamh)), flags, sliceFromArgv(argc, argv))
	// PAM_SESSION_ERR
	return C.PAM_SUCCESS
}

//export pam_sm_close_session
func pam_sm_close_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	log("pam_sm_close_session: user: %v, flags: %x, slice: %v", C.GoString(C.get_user(pamh)), flags, sliceFromArgv(argc, argv))
	// PAM_SESSION_ERR
	return C.PAM_SUCCESS
}

//export pam_sm_chauthtok
func pam_sm_chauthtok(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	log("pam_sm_chauthtok: user: %v, flags: %x, slice: %v", C.GoString(C.get_user(pamh)), flags, sliceFromArgv(argc, argv))
	return C.PAM_SUCCESS
}
