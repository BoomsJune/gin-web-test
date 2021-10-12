package handler

import (
	"net/http"
	"strconv"

	"example.com/web-test/internal/myapp/auth"
	"example.com/web-test/internal/pkg/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user auth.User
	c.BindJSON(&user)

	token, err := user.Register()
	if err != nil {
		util.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	util.ResponseOk(c, gin.H{"access_token": token})
}

func Login(c *gin.Context) {
	var user auth.User
	c.BindJSON(&user)

	token, err := user.Login()
	if err != nil {
		util.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	util.ResponseOk(c, gin.H{"access_token": token})
}

func Info(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := auth.Info(userId)
	if err != nil {
		util.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	util.ResponseOk(c, user)
}
