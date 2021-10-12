package auth

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"

	"example.com/web-test/internal/pkg/db"
	"example.com/web-test/internal/pkg/util"
)

type User struct {
	ID       uint    `json:"id,omitempty"`
	Phone    string  `json:"phone" binding:"required"`
	Password Encrypt `json:"password" binding:"required"`
	NickName string  `json:"nick_name,omitempty"`
}

func (User) TableName() string {
	return "auth_user"
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID       uint   `json:"id"`
		Phone    string `json:"phone"`
		NickName string `json:"nick_name,omitempty"`
	}{
		ID:       u.ID,
		Phone:    u.Phone,
		NickName: u.NickName,
	})
}

type Encrypt string

func (e *Encrypt) UnmarshalJSON(b []byte) (err error) {
	*e = Encrypt(fmt.Sprintf("%x", md5.Sum(b)))

	return err
}

func (u *User) Register() (string, error) {
	if err := db.DB.Create(&u).Error; err != nil {
		return "", err
	}

	token, err := util.JwtGet(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) Login() (string, error) {
	var returnUser User

	if err := db.DB.Where("phone = ?", u.Phone).First(&returnUser).Error; err != nil {
		return "", err
	}

	if returnUser.Password != u.Password {
		return "", errors.New("password error")
	}

	token, err := util.JwtGet(returnUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Info(userId int) (*User, error) {
	var user *User
	// db.DB.Debug().Where("id = ?", userId).First(&user)
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
