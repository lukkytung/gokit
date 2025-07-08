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

func GenerateID() (uint64, error) {
	return sf.NextID()
}
