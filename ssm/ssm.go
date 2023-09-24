package SSM

/*
#include <stdio.h>
#include <stdlib.h>

#cgo darwin,arm64 LDFLAGS: -L. ./ssm/SSM.dylib
#cgo linux LDFLAGS: -L. ./ssm/SSM.so
#include <SSM.h>

graal_isolate_t *isolate = NULL;
graal_isolatethread_t *thread = NULL;

*/
import "C"
import (
	"unsafe"
)

func GenerateKeyPair(publicKeyStr string, privateKeyStr string, userIdStr string, passwordStr string) {
	userId := C.CString(userIdStr)
	password := C.CString(passwordStr)
	privateKey := C.CString(privateKeyStr)
	publicKey := C.CString(publicKeyStr)

	defer C.free(unsafe.Pointer(userId))
	defer C.free(unsafe.Pointer(password))
	defer C.free(unsafe.Pointer(privateKey))
	defer C.free(unsafe.Pointer(publicKey))

	C.graal_create_isolate(nil, &C.isolate, &C.thread)
	defer C.graal_tear_down_isolate(C.thread)

	C.generateKeyPair(C.thread, publicKey, privateKey, userId, password)
}

func EncryptMessage(publicKeyStr string, messageStr string) string {
	message := C.CString(messageStr)
	publicKey := C.CString(publicKeyStr)

	defer C.free(unsafe.Pointer(message))
	defer C.free(unsafe.Pointer(publicKey))

	C.graal_create_isolate(nil, &C.isolate, &C.thread)
	defer C.graal_tear_down_isolate(C.thread)

	encrypted := C.encryptMessage(C.thread, publicKey, message)

	return C.GoString(encrypted)
}

func DecryptMessage(privateKeyStr string, passwordStr string, encryptedStr string) string {
	password := C.CString(passwordStr)
	privateKey := C.CString(privateKeyStr)
	encrypted := C.CString(encryptedStr)

	defer C.free(unsafe.Pointer(password))
	defer C.free(unsafe.Pointer(privateKey))
	defer C.free(unsafe.Pointer(encrypted))

	C.graal_create_isolate(nil, &C.isolate, &C.thread)
	defer C.graal_tear_down_isolate(C.thread)

	decrypted := C.decryptMessage(C.thread, privateKey, password, encrypted)

	return C.GoString(decrypted)
}
