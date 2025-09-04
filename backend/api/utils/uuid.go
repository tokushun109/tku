package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID はランダムなUUIDを生成します
func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}