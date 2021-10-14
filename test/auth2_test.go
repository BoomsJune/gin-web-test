package test

import (
	"crypto/md5"
	"fmt"
	"testing"

	"example.com/web-test/config"
	"example.com/web-test/internal/auth"
	"example.com/web-test/internal/pkg/db"
	"github.com/go-playground/assert/v2"
)

func TestRegister2(t *testing.T) {
	defer ResetDB()

	params := []byte(`{"phone":"12345", "password":"12345", "nick_name":"name"}`)
	res, err := Request("POST", "/regist", params)

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Code)

	data := res.Data.(map[string]interface{})

	assert.NotEqual(t, "", data["access_token"])

	// check database
	var user *auth.User
	err = db.DB.Where("phone = ?", "12345").First(&user).Error

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, user.ID)
	pwdEyt := fmt.Sprintf("%s.%d.%s", "12345", user.CreatedAt.Unix(), config.Cfg.App.SignKey)
	assert.Equal(t, user.Password, fmt.Sprintf("%x", md5.Sum([]byte(pwdEyt))))

}

func TestRegisterWithUniquePhone2(t *testing.T) {
	defer ResetDB()

	// create A user
	params := []byte(`{"phone":"12345", "password":"12345", "nick_name":"A"}`)
	res, err := Request("POST", "/regist", params)

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Code)

	// create B user with same phone number
	params = []byte(`{"phone":"12345", "password":"67890", "nick_name":"B"}`)
	res, err = Request("POST", "/regist", params)

	assert.Equal(t, nil, err)
	assert.Equal(t, 400, res.Code)

}

func TestRegisterWithNilField2(t *testing.T) {
	defer ResetDB()

	params := [][]byte{
		[]byte(`{"phone":"", "password":"12345", "nick_name":"A"}`), // nil phone
		[]byte(`{"phone":"12345", "password":"", "nick_name":"A"}`), // nil password
	}

	for _, p := range params {
		res, err := Request("POST", "/regist", p)

		assert.Equal(t, nil, err)
		assert.Equal(t, 400, res.Code)
	}
}
