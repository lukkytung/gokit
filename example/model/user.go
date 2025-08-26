package model

import (
	"fmt"
	"time"

	"github.com/lukkytung/gokit/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid      string `gorm:"unique"`
	Email    string `gorm:"uniqueIndex"`
	Avatar   string
	Gender   string
	Birthday time.Time
}

// BeforeCreate 创建用户时生成唯一 ID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	uid, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Failed to generate user ID", err)
		return
	}
	u.Uid = uid
	return
}
