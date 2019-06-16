package crypto

import (
	"golang.org/x/crypto/sha3"
)

func Sha3(data []byte) ([]byte, error) {
	lk := sha3.NewLegacyKeccak256()
	_, err := lk.Write(data)
	if err != nil {
		return nil, err
	}
	return lk.Sum(nil), nil
}
