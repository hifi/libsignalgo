package libsignalgo_test

import (
	"testing"
	"time"

	"github.com/beeper/libsignalgo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSenderCertificate_Operations(t *testing.T) {
	setupLogging()
	senderCertBits := []byte{
		0x0a, 0xcd, 0x01, 0x0a, 0x0c, 0x2b, 0x31, 0x34, 0x31, 0x35, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0x10, 0x2a, 0x19,
		0x2d, 0x63, 0xb5, 0x5f, 0x00, 0x00, 0x00, 0x00, 0x22, 0x21, 0x05, 0xbb, 0x25, 0x64, 0x9c, 0x79, 0x4b, 0xb4, 0x6c, 0x8c,
		0x57, 0x97, 0x69, 0x3c, 0xc8, 0x05, 0xb1, 0xb8, 0x46, 0xda, 0x91, 0x17, 0x6f, 0xec, 0x6a, 0x3e, 0xf2, 0x1f, 0x41, 0x0b,
		0xe9, 0x60, 0x43, 0x2a, 0x69, 0x0a, 0x25, 0x08, 0x01, 0x12, 0x21, 0x05, 0x4f, 0xbf, 0xfa, 0x55, 0xeb, 0xd5, 0x23, 0xd2,
		0x55, 0x16, 0x96, 0x0c, 0xed, 0x28, 0x99, 0xf2, 0x6a, 0x72, 0xfe, 0x26, 0xd0, 0xe0, 0x2a, 0x9d, 0xae, 0x81, 0x67, 0x1f,
		0x46, 0x5b, 0xa1, 0x1d, 0x12, 0x40, 0x7a, 0xbf, 0xdb, 0x83, 0x6c, 0x15, 0xcb, 0x3a, 0x8c, 0x61, 0x76, 0xb3, 0x30, 0x70,
		0xdf, 0xbc, 0x47, 0xea, 0x4a, 0x90, 0x52, 0x35, 0x3a, 0xc4, 0x2f, 0xb8, 0x7e, 0x4e, 0x4d, 0x33, 0x4f, 0x69, 0xa5, 0xe0,
		0xd4, 0xab, 0xd2, 0xdd, 0x81, 0x9f, 0x61, 0xa2, 0xc0, 0x2a, 0x51, 0xc2, 0x74, 0x51, 0xc9, 0x31, 0xaa, 0x85, 0x35, 0xf8,
		0x32, 0x8d, 0x1e, 0xc8, 0xce, 0x7a, 0x2b, 0x9a, 0x9e, 0x01, 0x32, 0x24, 0x39, 0x64, 0x30, 0x36, 0x35, 0x32, 0x61, 0x33,
		0x2d, 0x64, 0x63, 0x63, 0x33, 0x2d, 0x34, 0x64, 0x31, 0x31, 0x2d, 0x39, 0x37, 0x35, 0x66, 0x2d, 0x37, 0x34, 0x64, 0x36,
		0x31, 0x35, 0x39, 0x38, 0x37, 0x33, 0x33, 0x66, 0x12, 0x40, 0x06, 0x8b, 0xf0, 0xc5, 0xe8, 0x99, 0x83, 0x81, 0x28, 0xbd,
		0x36, 0xd9, 0x2b, 0x01, 0xec, 0xa9, 0x95, 0x9d, 0x00, 0xf2, 0xdb, 0x0b, 0xcb, 0xb6, 0x8b, 0x2a, 0x62, 0xd4, 0xdf, 0x46,
		0xdb, 0xb4, 0x50, 0x14, 0x9e, 0x9d, 0xcb, 0xc6, 0xbd, 0xdb, 0x2b, 0x28, 0x98, 0xfc, 0xd5, 0xff, 0x5c, 0xaf, 0x1b, 0x8c,
		0xf7, 0x2b, 0x36, 0xff, 0xfe, 0x2f, 0x55, 0xf3, 0xec, 0xeb, 0xab, 0x25, 0x47, 0x88,
	}

	senderCertificate, err := libsignalgo.DeserializeSenderCertificate(senderCertBits)
	assert.NoError(t, err)
	assert.NotNil(t, senderCertificate)

	t.Run("serialize", func(t *testing.T) {
		serialized, err := senderCertificate.Serialize()
		assert.NoError(t, err)
		assert.Equal(t, senderCertBits, serialized)
	})

	t.Run("expiration", func(t *testing.T) {
		expiration, err := senderCertificate.GetExpiration()
		assert.NoError(t, err)
		assert.True(t, time.Date(2020, time.November, 18, 18, 8, 45, 0, time.UTC).Equal(expiration))
	})

	t.Run("device ID", func(t *testing.T) {
		deviceID, err := senderCertificate.GetDeviceID()
		assert.NoError(t, err)
		assert.Equal(t, uint32(42), deviceID)
	})

	t.Run("public key", func(t *testing.T) {
		publicKey, err := senderCertificate.GetKey()
		assert.NoError(t, err)
		assert.NotNil(t, publicKey)

		serialized, err := publicKey.Serialize()
		assert.NoError(t, err)
		assert.Len(t, serialized, 33)
	})

	t.Run("sender UUID", func(t *testing.T) {
		senderUUID, err := senderCertificate.GetSenderUUID()
		assert.NoError(t, err)
		expectedUUID, err := uuid.Parse("9d0652a3-dcc3-4d11-975f-74d61598733f")
		assert.NoError(t, err)
		assert.Equal(t, expectedUUID, senderUUID)
	})

	t.Run("sender E164", func(t *testing.T) {
		senderE164, err := senderCertificate.GetSenderE164()
		assert.NoError(t, err)
		assert.Equal(t, "+14152222222", senderE164)
	})

	t.Run("server certificate", func(t *testing.T) {
		serverCertificate, err := senderCertificate.GetServerCertificate()
		assert.NoError(t, err)
		assert.NotNil(t, serverCertificate)

		keyID, err := serverCertificate.GetKeyId()
		assert.NoError(t, err)
		assert.Equal(t, uint32(1), keyID)

		serverPublicKey, err := serverCertificate.GetKey()
		assert.NoError(t, err)
		assert.NotNil(t, serverPublicKey)

		serverPublicKeyBytes, err := serverPublicKey.Serialize()
		assert.NoError(t, err)
		assert.Len(t, serverPublicKeyBytes, 33)

		serverSignature, err := serverCertificate.GetSignature()
		assert.NoError(t, err)
		assert.Len(t, serverSignature, 64)
	})
}

type Serializable interface {
	Serialize() ([]byte, error)
}

func testRoundTrip[T Serializable](t *testing.T, name string, obj T, deserializer func([]byte) (T, error)) {
	t.Run(name, func(t *testing.T) {
		serialized, err := obj.Serialize()
		assert.NoError(t, err)

		deserialized, err := deserializer(serialized)
		assert.NoError(t, err)

		deserializedSerialized, err := deserialized.Serialize()
		assert.NoError(t, err)

		assert.Equal(t, serialized, deserializedSerialized)
	})
}

func TestSenderCertificateSerializationRoundTrip(t *testing.T) {
	keyPair, err := libsignalgo.GenerateIdentityKeyPair()
	assert.NoError(t, err)

	testRoundTrip(t, "key pair", keyPair, libsignalgo.DeserializeIdentityKeyPair)
	testRoundTrip(t, "public key", keyPair.GetPublicKey(), libsignalgo.DeserializePublicKey)
	testRoundTrip(t, "private key", keyPair.GetPrivateKey(), libsignalgo.DeserializePrivateKey)
	testRoundTrip(t, "identity key", keyPair.GetIdentityKey(), libsignalgo.NewIdentityKeyFromBytes)

	preKeyRecord, err := libsignalgo.NewPreKeyRecord(7, keyPair.GetPublicKey(), keyPair.GetPrivateKey())
	assert.NoError(t, err)
	testRoundTrip(t, "pre key record", preKeyRecord, libsignalgo.DeserializePreKeyRecord)

	publicKeySerialized, err := keyPair.GetPublicKey().Serialize()
	assert.NoError(t, err)
	signature, err := keyPair.GetPrivateKey().Sign(publicKeySerialized)
	assert.NoError(t, err)

	signedPreKeyRecord, err := libsignalgo.NewSignedPreKeyRecordFromPrivateKey(
		77,
		time.UnixMilli(42000),
		keyPair.GetPrivateKey(),
		signature,
	)
	assert.NoError(t, err)
	testRoundTrip(t, "signed pre key record", signedPreKeyRecord, libsignalgo.DeserializeSignedPreKeyRecord)
}
