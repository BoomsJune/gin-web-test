package auth

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"example.com/web-test/internal/pkg/config"
	"example.com/web-test/internal/pkg/db"
	"example.com/web-test/internal/pkg/util"
)

type User struct {
	util.Model
	Phone    string `json:"phone" binding:"required" gorm:"unique;not null;size:20"`
	Password string `json:"password" binding:"required" gorm:"not null;size:32"`
	NickName string `json:"nick_name,omitempty" gorm:"size:20"`
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

type encryptParam struct {
	password  string
	createdTs int64
}

// 给密码加密，密码+注册时间+签名密钥
func encrypt(p *encryptParam) string {
	if p.createdTs == 0 {
		p.createdTs = time.Now().Unix()
	}

	content := fmt.Sprintf("%s.%d.%s", p.password, p.createdTs, config.Cfg.App.SignKey)
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}

func (u *User) Register() (string, error) {
	encrypt(&encryptParam{password: u.Password})

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

	encrypt(&encryptParam{password: u.Password, createdTs: returnUser.CreatedAt.Unix()})

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
