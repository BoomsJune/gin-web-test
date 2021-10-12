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
	var userId uint
	sql := "INSERT INTO auth_user(phone, password, nick_name) VALUES ($1, $2, $3) RETURNING id"
	err := db.DB.QueryRow(sql, u.Phone, u.Password, u.NickName).Scan(&userId)
	if err != nil {
		return "", err
	}
	token, err := util.JwtGet(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *User) Login() (string, error) {
	var userId uint
	var password string

	sql := "SELECT id, password FROM auth_user WHERE phone = $1"
	err := db.DB.QueryRow(sql, u.Phone).Scan(&userId, &password)
	if err != nil {
		return "", err
	}

	if password != string(u.Password) {
		return "", errors.New("password error")
	}

	token, err := util.JwtGet(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Info(userId int) (*User, error) {
	var phone, nick_name string

	sql := "SELECT phone, nick_name FROM auth_user WHERE id = $1"
	err := db.DB.QueryRow(sql, userId).Scan(&phone, &nick_name)
	if err != nil {
		return nil, err
	}
	return &User{ID: uint(userId), Phone: phone, NickName: nick_name}, nil
}
