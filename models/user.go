package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/kenny26/maynooth-api/utils"
)

type User struct {
	Model
	Username string `json:"username" gorm:"type:varchar(100);unique_index"`
	Email string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Gender string `json:"gender" gorm:"type:varchar(25)"`
	Dob string `json:"dob" gorm:"type:date"`
}

func (user *User) Validate() (map[string] interface{}, bool) {

	//Username and email must be unique
	temp := &User{}

	//check for errors and duplicate username or email
	err := GetDB().Table("users").
		Or("username = ?", user.Username).
		Where("email = ?", user.Email).
		First(temp).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email == user.Email {
		return utils.Message(false, "Email address already in use by another user."), false
	}
	if temp.Username == user.Username {
		return utils.Message(false, "Username already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

func (user *User) Create() (*User) {

	if _, ok := user.Validate(); !ok {
		return nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	return user
}

func GetUserById(id int) (*User) {

	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func AuthenticateUser(username string, password string) (*User) {

	user := &User{}
	err := GetDB().Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil
	}

	return user
}
