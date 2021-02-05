package secure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		v []byte
		s string
		n string
	}{
		{v: []byte{}, s: "olia", n: "empty"},
		{v: []byte("olia"), s: "olia", n: "simple"},
		{v: []byte("o"), s: "olia", n: "one letter"},
	}

	for i, tc := range tests {
		v, err := Encrypt(tc.v, tc.s)
		assert.Nil(t, err, "fail %d - %s", i, tc.n)
		vd, err := Decrypt(v, tc.s)
		assert.Nil(t, err, "fail %d - %s", i, tc.n)
		assert.Equal(t, tc.v, vd, "fail %d - %s", i, tc.n)
	}
}

func TestEncryptFail(t *testing.T) {
	_, err := Encrypt([]byte("tata"), "")
	assert.NotNil(t, err)
}

func TestDecryptFail(t *testing.T) {
	v, err := Encrypt([]byte("tata"), "aaa")
	_, err = Decrypt(v, "aaa1")
	assert.NotNil(t, err)
}
