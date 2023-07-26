package common

import (
	"errors"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidEmail         = errors.New("email already in use")
	ErrBadRequest           = errors.New("bad Request Payload")
	ErrLogin                = errors.New("wrong email or password")
	ErrTokenMissing         = errors.New("refresh token is missing")
	ErrCreateToken          = errors.New("cannot create token")
	ErrInvalidCode          = errors.New("invalid verification code")
	ErrAlreadyVerified      = errors.New("user already verified")
	ErrInactiveUser         = errors.New("user not verified")
	ErrResetPassword        = errors.New("cannot reset password")
	ErrFailedPasswordReset  = errors.New("wrong email or code")
	ErrInterval             = errors.New("something went wrong, try again")
	ErrPasswordResetExpired = errors.New("password reset code is expired")
	ErrMissingContext       = errors.New("userID not found in request context")
	ErrOldPassword          = errors.New("old password must match")
	ErrUpdatePassword       = errors.New("cannot update password")
	ErrPermission           = errors.New("poor authorization")
	ErrInvalidRoleCode      = errors.New("a role with given code already exists")
	ErrInvalidRoleName      = errors.New("a role with given name already exists")
	ErrRoleNotFound         = errors.New("role not found")
	ErrDecryption           = errors.New("decryption failed: invalid UserID")
)
