package libsignalgo

/*
#cgo LDFLAGS: ./libsignal/target/release/libsignal_ffi.a -ldl
#include "./libsignal/libsignal-ffi.h"
*/
import "C"
import "unsafe"

func HKDFDerive(outputLength int, inputKeyMaterial, salt, info []byte) ([]byte, error) {
	output := BorrowedMutableBuffer(outputLength)
	signalFfiError := C.signal_hkdf_derive(output, BytesToBuffer(inputKeyMaterial), BytesToBuffer(salt), BytesToBuffer(info))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return C.GoBytes(unsafe.Pointer(output.base), C.int(output.length)), nil
}