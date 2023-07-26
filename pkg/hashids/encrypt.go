package hashid

import (
	"cfd/myapp/internal/common"
	"github.com/speps/go-hashids/v2"
)

func loadHD() *hashids.HashIDData {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30

	return hd
}

func Encrypt(userID int) (*string, error) {
	hd := loadHD()

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	e, err := h.Encode([]int{userID})
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func Decrypt(encodedUserID string) (*int, error) {
	hd := loadHD()

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	decodedUserID, err := h.DecodeWithError(encodedUserID)
	if err != nil {
		return nil, err
	}

	if len(decodedUserID) == 0 {
		return nil, common.ErrDecryption
	}

	return &decodedUserID[0], nil
}
