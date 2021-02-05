package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
)

const (
	salt = "UVegh3shahngad4ahJ9luoX3Iengai"
)

//Encrypt encrypts data using secret phrase
func Encrypt(data []byte, secret string) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}
	bSecret, err := prepareKey([]byte(secret), []byte(salt))
	if err != nil {
		return nil, errors.Wrapf(err, "Can't prepare pass key")
	}
	res, err := encrypt(data, bSecret)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't encrypt")
	}
	return res, nil
}

//Decrypt data using secret
func Decrypt(data []byte, secret string) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}
	bSecret, err := prepareKey([]byte(secret), []byte(salt))
	if err != nil {
		return nil, errors.Wrapf(err, "Can't prepare pass key")
	}
	res, err := decrypt(data, bSecret)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't decrypt")
	}
	return res, nil
}

func encrypt(data, key []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data, key []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func prepareKey(password, salt []byte) ([]byte, error) {
	if len(password) == 0 {
		return nil, errors.New("No secret")
	}
	if len(salt) == 0 {
		return nil, errors.New("No secret")
	}
	key, err := scrypt.Key(password, salt, 32768, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return key, nil
}
