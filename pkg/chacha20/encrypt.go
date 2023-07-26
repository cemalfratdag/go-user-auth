package chacha20

import (
	"cfd/myapp/internal/common"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
)

var secret = "8fd32616a75f1131d5e37f9664b9089d9c329e550c79e9ad192499fefe125963"

func Encrypt(userID int) (string, error) {
	key, err := hex.DecodeString(secret)
	if err != nil {
		return "", err
	}

	c, err := chacha20poly1305.New(key)
	if err != nil {
		return "", err
	}

	userIDBytes := []byte(fmt.Sprintf("%d", userID))

	nonce := make([]byte, c.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := c.Seal(nil, nonce, userIDBytes, nil)
	encodedString := base64.URLEncoding.EncodeToString(append(nonce, ciphertext...))

	return encodedString, nil
}

func Decrypt(encodedUserID string) (int, error) {
	key, err := hex.DecodeString(secret)
	if err != nil {
		return 0, err
	}

	c, err := chacha20poly1305.New(key)
	if err != nil {
		return 0, err
	}

	data, err := base64.URLEncoding.DecodeString(encodedUserID)
	if err != nil {
		return 0, err
	}

	if len(data) < c.NonceSize() {
		return 0, common.ErrInterval
	}

	nonce, ciphertext := data[:c.NonceSize()], data[c.NonceSize():]
	decryptedUserID, err := c.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return 0, err
	}

	var userID int
	fmt.Sscanf(string(decryptedUserID), "%d", &userID)

	return userID, nil
}
