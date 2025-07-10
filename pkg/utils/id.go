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

func GenerateID() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		log.Printf("failed to generate ID: %v", err)
		return "", err
	}
	return strconv.Itoa(int(id)), nil
}
