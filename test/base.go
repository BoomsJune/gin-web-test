package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"example.com/web-test/cmd/router"
	"example.com/web-test/internal/pkg/db/migrate"
	"example.com/web-test/util"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = router.Register()
}

func ResetDB() {
	migrate.DropAll()
	migrate.MigrateAll()
}

func Request(method string, url string, params []byte) (*util.BaseResponse, error) {
	base := &util.BaseResponse{}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(params))
	if method != "GET" {
		request.Header.Set("Content-Type", "application/json")
	}

	r.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		return nil, fmt.Errorf("%d != %d", recorder.Code, http.StatusOK)
	}

	if err := json.Unmarshal(recorder.Body.Bytes(), &base); err != nil {
		return nil, err
	}

	return base, nil
}
