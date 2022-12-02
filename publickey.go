package libsignalgo

/*
#cgo LDFLAGS: ./libsignal/target/release/libsignal_ffi.a -ldl
#include "./libsignal/libsignal-ffi.h"
*/
import "C"
import "runtime"

type PublicKey struct {
	ptr *C.SignalPublicKey
}

func wrapPublicKey(ptr *C.SignalPublicKey) *PublicKey {
	publicKey := &PublicKey{ptr: ptr}
	runtime.SetFinalizer(publicKey, (*PublicKey).Destroy)
	return publicKey
}

func (pk *PublicKey) Clone() (*PublicKey, error) {
	var cloned *C.SignalPublicKey
	signalFfiError := C.signal_publickey_clone(&cloned, pk.ptr)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapPublicKey(cloned), nil
}

func DeserializePublicKey(keyData []byte) (*PublicKey, error) {
	var pk *C.SignalPublicKey
	signalFfiError := C.signal_publickey_deserialize(&pk, BytesToBuffer(keyData))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapPublicKey(pk), nil
}

func (pk *PublicKey) Serialize() ([]byte, error) {
	var serialized *C.uchar
	var length C.ulong
	signalFfiError := C.signal_publickey_serialize(&serialized, &length, pk.ptr)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return CopyBufferToBytes(serialized, length), nil
}

func (k *PublicKey) Destroy() error {
	runtime.SetFinalizer(k, nil)
	return wrapError(C.signal_publickey_destroy(k.ptr))
}

func (k *PublicKey) Compare(other *PublicKey) (int, error) {
	var comparison C.int
	signalFfiError := C.signal_publickey_compare(&comparison, k.ptr, other.ptr)
	if signalFfiError != nil {
		return 0, wrapError(signalFfiError)
	}
	return int(comparison), nil
}

func (k *PublicKey) Bytes() ([]byte, error) {
	var pub *C.uchar
	var length C.ulong
	signalFfiError := C.signal_publickey_get_public_key_bytes(&pub, &length, k.ptr)
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return CopyBufferToBytes(pub, length), nil
}

func (k *PublicKey) Verify(message, signature []byte) (bool, error) {
	var verify C.bool
	signalFfiError := C.signal_publickey_verify(&verify, k.ptr, BytesToBuffer(message), BytesToBuffer(signature))
	if signalFfiError != nil {
		return false, wrapError(signalFfiError)
	}
	return bool(verify), nil
}
