package utils

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zheng-ji/goSnowFlake"
)

//GenUUID create uuid based on random numbers
func GenUUID() (string, error) {
	// or error handling
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), err
}

//GenUID uidGenerater
func GenUID(workID int64) (int64, error) {
	iw, err := goSnowFlake.NewIdWorker(workID)
	if err != nil {
		return 0, err
	}
	return iw.NextId()
}
