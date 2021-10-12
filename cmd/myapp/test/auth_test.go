package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/web-test/cmd/myapp/router"
	"github.com/go-playground/assert/v2"
)

func TestRegister(t *testing.T) {
	r := router.Register()

	recorder := httptest.NewRecorder()

	json := []byte(`{"phone":"12345", "password":"12345", "nick_name":"name"}`)
	request, _ := http.NewRequest("POST", "/regist", bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
}
