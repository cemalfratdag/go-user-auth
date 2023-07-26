package helper

import "github.com/google/uuid"

func CreateVerificationCode() string {
	return uuid.New().String()[:10]
}
