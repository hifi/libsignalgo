package libsignalgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/beeper/libsignalgo"
)

// From PublicAPITests.swift:testPkOperations
func TestPrivateKeyOperations(t *testing.T) {
	setupLogging()
	var err error
	var privateKey *libsignalgo.PrivateKey

	t.Run("test generate", func(t *testing.T) {
		privateKey, err = libsignalgo.GeneratePrivateKey()
		assert.NoError(t, err)
		assert.NotNil(t, privateKey)
	})

	var privateKeyBytes []byte
	t.Run("serialize", func(t *testing.T) {
		privateKeyBytes, err = privateKey.Serialize()
		assert.NoError(t, err)
		assert.NotNil(t, privateKeyBytes)
	})

	var publicKey *libsignalgo.PublicKey
	t.Run("get public key", func(t *testing.T) {
		publicKey, err = privateKey.GetPublicKey()
		assert.NoError(t, err)
		assert.NotNil(t, publicKey)
	})

	var publicKeyBytes []byte
	t.Run("serialize public key", func(t *testing.T) {
		publicKeyBytes, err = publicKey.Serialize()
		assert.NoError(t, err)
		assert.NotNil(t, publicKeyBytes)

		assert.EqualValues(t, 5, publicKeyBytes[0])
		assert.Len(t, publicKeyBytes, 33)
	})

	var publicKeyRaw []byte
	t.Run("get public key raw", func(t *testing.T) {
		publicKeyRaw, err = publicKey.Bytes()
		assert.NoError(t, err)
		assert.NotNil(t, publicKeyRaw)

		assert.Len(t, publicKeyRaw, 32)
		assert.Equal(t, publicKeyRaw[0:31], publicKeyBytes[1:32])
	})

	var privateKeyReloaded *libsignalgo.PrivateKey
	var publicKeyReloaded *libsignalgo.PublicKey
	t.Run("deserialize private key", func(t *testing.T) {
		privateKeyReloaded, err = libsignalgo.DeserializePrivateKey(privateKeyBytes)
		assert.NoError(t, err)
		assert.NotNil(t, privateKeyReloaded)

		publicKeyReloaded, err = privateKeyReloaded.GetPublicKey()
		assert.NoError(t, err)
		assert.NotNil(t, publicKeyReloaded)

		assert.Equal(t, publicKey, publicKeyReloaded)

		serializedPublicKey, err := publicKey.Serialize()
		assert.NoError(t, err)
		serializedPublicKeyReloaded, err := publicKeyReloaded.Serialize()
		assert.NoError(t, err)
		assert.Equal(t, serializedPublicKey, serializedPublicKeyReloaded)
	})

	t.Run("sign", func(t *testing.T) {
		message := []byte{0x01, 0x02, 0x03}
		signature, err := privateKey.Sign(message)
		assert.NoError(t, err)

		valid, err := publicKey.Verify(message, signature)
		assert.NoError(t, err)
		assert.True(t, valid)

		signature[5] ^= 1

		valid, err = publicKey.Verify(message, signature)
		assert.NoError(t, err)
		assert.False(t, valid)

		signature[5] ^= 1

		valid, err = publicKey.Verify(message, signature)
		assert.NoError(t, err)
		assert.True(t, valid)

		message[1] ^= 1

		valid, err = publicKey.Verify(message, signature)
		assert.NoError(t, err)
		assert.False(t, valid)

		message[1] ^= 1

		valid, err = publicKey.Verify(message, signature)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("agree", func(t *testing.T) {
		privateKey2, err := libsignalgo.GeneratePrivateKey()
		assert.NoError(t, err)

		publicKey2, err := privateKey2.GetPublicKey()
		assert.NoError(t, err)

		sharedSecret1, err := privateKey.Agree(publicKey2)
		sharedSecret2, err := privateKey2.Agree(publicKey)

		assert.Equal(t, sharedSecret1, sharedSecret2)
	})
}
