package chacha20

import (
	"math/rand"
	"testing"
	"time"
)

func generateRandomUserID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}

func TestEncryptDecrypt(t *testing.T) {

	for i := 0; i < 100000; i++ {
		userID := generateRandomUserID()

		encryptedUserID, err := Encrypt(userID)
		if err != nil {
			t.Errorf("Error encrypting userID %d: %v", userID, err)
			continue
		}

		decryptedUserID, err := Decrypt(encryptedUserID)
		if err != nil {
			t.Errorf("Error decrypting userID %s: %v", encryptedUserID, err)
			continue
		}

		if decryptedUserID != userID {
			t.Errorf("Decrypted userID (%d) does not match the original userID (%d)", decryptedUserID, userID)
		}
	}
}
