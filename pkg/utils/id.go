package utils

import (
	"log"
	"strconv"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func InitIDGenerator() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		log.Fatal("failed to initialize sonyflake")
	}
}

func GenerateID() (uint, error) {
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
func GenerateIDStr() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}
