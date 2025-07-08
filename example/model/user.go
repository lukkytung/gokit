package model

import (
	"fmt"

	"github.com/lukkytung/gokit/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid   uint   `gorm:"unique"`
	Email string `gorm:"unique"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Failed to generate user ID", err)
		return
	}
	u.Uid = uid
	return
}
