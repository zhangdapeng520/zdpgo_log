package zap

import (
	"testing"

	"github.com/zhangdapeng520/zdpgo_log/libs/zap/zapcore"

	"github.com/stretchr/testify/assert"
)

func TestRegisterDefaultEncoders(t *testing.T) {
	testEncodersRegistered(t, "console", "json")
}

func TestRegisterEncoder(t *testing.T) {
	testEncoders(func() {
		assert.NoError(t, RegisterEncoder("foo", newNilEncoder), "expected to be able to register the encoder foo")
		testEncodersRegistered(t, "foo")
	})
}

func TestDuplicateRegisterEncoder(t *testing.T) {
	testEncoders(func() {
		RegisterEncoder("foo", newNilEncoder)
		assert.Error(t, RegisterEncoder("foo", newNilEncoder), "expected an error when registering an encoder with the same name twice")
	})
}

func TestRegisterEncoderNoName(t *testing.T) {
	assert.Equal(t, errNoEncoderNameSpecified, RegisterEncoder("", newNilEncoder), "expected an error when registering an encoder with no name")
}

func TestNewEncoder(t *testing.T) {
	testEncoders(func() {
		RegisterEncoder("foo", newNilEncoder)
		encoder, err := newEncoder("foo", zapcore.EncoderConfig{})
		assert.NoError(t, err, "could not create an encoder for the registered name foo")
		assert.Nil(t, encoder, "the encoder from newNilEncoder is not nil")
	})
}

func TestNewEncoderNotRegistered(t *testing.T) {
	_, err := newEncoder("foo", zapcore.EncoderConfig{})
	assert.Error(t, err, "expected an error when trying to create an encoder of an unregistered name")
}

func TestNewEncoderNoName(t *testing.T) {
	_, err := newEncoder("", zapcore.EncoderConfig{})
	assert.Equal(t, errNoEncoderNameSpecified, err, "expected an error when creating an encoder with no name")
}

func testEncoders(f func()) {
	existing := _encoderNameToConstructor
	_encoderNameToConstructor = make(map[string]func(zapcore.EncoderConfig) (zapcore.Encoder, error))
	defer func() { _encoderNameToConstructor = existing }()
	f()
}

func testEncodersRegistered(t *testing.T, names ...string) {
	assert.Len(t, _encoderNameToConstructor, len(names), "the expected number of registered encoders does not match the actual number")
	for _, name := range names {
		assert.NotNil(t, _encoderNameToConstructor[name], "no encoder is registered for name %s", name)
	}
}

func newNilEncoder(_ zapcore.EncoderConfig) (zapcore.Encoder, error) {
	return nil, nil
}
