package utils

import (
	uuid "github.com/satori/go.uuid"
)

//GenUUID create uuid based on random numbers
func GenUUID() (uuid.UUID, error) {
	// or error handling
	u, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}
	return u, err
}
