package libsignalgo

/*
#cgo LDFLAGS: ./libsignal/target/release/libsignal_ffi.a -ldl
#include "./libsignal/libsignal-ffi.h"
*/
import "C"
import "unsafe"

type DeviceTransferKey struct {
	privateKey []byte
}

func GenerateDeviceTransferKey() (*DeviceTransferKey, error) {
	var resp *C.uchar
	var length C.ulong
	signalFfiError := C.signal_device_transfer_generate_private_key(&resp, &length)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return &DeviceTransferKey{
		privateKey: C.GoBytes(unsafe.Pointer(resp), C.int(length)),
	}, nil
}

func (dtk *DeviceTransferKey) Bytes() []byte {
	return dtk.privateKey
}

func (dtk *DeviceTransferKey) GenerateCertificate(name string, days int) ([]byte, error) {
	var resp *C.uchar
	var length C.ulong
	signalFfiError := C.signal_device_transfer_generate_certificate(&resp, &length, BytesToBuffer(dtk.privateKey), C.CString(name), C.uint32_t(days))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return C.GoBytes(unsafe.Pointer(resp), C.int(length)), nil
}
