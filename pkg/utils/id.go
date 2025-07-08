package utils

import (
	"log"

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
