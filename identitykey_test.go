package libsignalgo_test

import (
	"testing"

	"github.com/beeper/libsignalgo"
	"github.com/stretchr/testify/assert"
)

// From PublicAPITests.swift:testSignAlternateIdentity
func TestSignAlternateIdentity(t *testing.T) {
	primary, err := libsignalgo.GenerateIdentityKeyPair()
	assert.NoError(t, err)
	secondary, err := libsignalgo.GenerateIdentityKeyPair()
	assert.NoError(t, err)

	signature, err := secondary.SignAlternateIdentity(primary.GetIdentityKey())
	assert.NoError(t, err)

	verified, err :=secondary.GetIdentityKey().VerifyAlternateIdentity(primary.GetIdentityKey(), signature)
	assert.NoError(t, err)
	assert.True(t, verified)
}
