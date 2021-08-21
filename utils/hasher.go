package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"github.com/sirupsen/logrus"
)

const DefaultSaltSize = 512

type hasher struct {
	Secret string `json:"secret" bson:"secret"`
}

// type IHasher interface {
// 	generateSalt() (salt []byte, err error)
// 	Hash(password string, salt []byte) (hashedPassword string, err error)
// 	IsMatch(hashedPassword, password string, salt []byte) bool
// }

func NewHasher() *hasher {
	return new(hasher)
}

func (h *hasher) LoadSecret(secret string) *hasher {
	h.Secret = secret
	return h
}

func (h *hasher) Hash(value string) (hashedPassword string, err error) {
	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Password + Secret
	passwordWithSalt := append([]byte(value), []byte(h.Secret)...)

	// Hashing
	if _, err = sha512Hasher.Write(passwordWithSalt); err != nil {
		return "", fmt.Errorf("error hashing value: %v", err)
	}

	// Get hashed SHA-512 value
	var byteHashedPassword = sha512Hasher.Sum(nil)

	return base64.URLEncoding.EncodeToString(byteHashedPassword), nil
}

func (h *hasher) IsMatch(currentValue string, hashedValue string) bool {
	var currentHash, err = h.Hash(currentValue)
	if err != nil {
		logrus.Errorf("can't hash propperly: %v", err)
		return false
	}

	return hashedValue == currentHash
}

func (h *hasher) GenerateSecret(length ...int) (secret string, err error) {
	if length == nil {
		length[0] = DefaultSaltSize
	}
	raw := make([]byte, length[0])

	// create salt
	if _, err = rand.Read(raw[:]); err != nil {
		return "", fmt.Errorf("can't create salt: %v", err)
	}

	// encode b64
	secret = base64.URLEncoding.EncodeToString(raw)

	return
}

// without salt

// func (h *hasher) Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
//
// 	// Since the key is in string, we need to convert decode it to bytes
// 	key, _ := hex.DecodeString(keyString)
// 	plaintext := []byte(stringToEncrypt)
//
// 	// Create a new Cipher Block from the key
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
// 	// https://golang.org/pkg/crypto/cipher/#NewGCM
// 	aesGCM, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	// Create a nonce. Nonce should be from GCM
// 	nonce := make([]byte, aesGCM.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		panic(err.Error())
// 	}
//
// 	// Encrypt the data using aesGCM.Seal
// 	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
// 	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
// 	return fmt.Sprintf("%x", ciphertext)
//
// }

// func (h *hasher) Decrypt(encryptedString string, keyString string) (decryptedString string) {
//
// 	key, _ := hex.DecodeString(keyString)
// 	enc, _ := hex.DecodeString(encryptedString)
//
// 	// Create a new Cipher Block from the key
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	// Create a new GCM
// 	aesGCM, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	// Get the nonce size
// 	nonceSize := aesGCM.NonceSize()
//
// 	// Extract the nonce from the encrypted data
// 	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
//
// 	// Decrypt the data
// 	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	return fmt.Sprintf("%s", plaintext)
// }
