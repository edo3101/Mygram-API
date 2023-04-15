package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"tesjwt.go/helpers"
)

type User struct {
	GormModel
	Username string `gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
