package middleware

import (
	"net/http"
	"strconv"

	"example.com/web-test/internal/auth"
	"example.com/web-test/util"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before exec
		claims, err := util.JwtParse(c.Request.Header)
		if err != nil {
			util.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		identity, err := strconv.Atoi(claims.Subject)
		if err != nil {
			util.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		user, err := auth.Info(identity)
		if err != nil {
			util.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		// send user to current request
		c.Set("user", user)

		// exec
		c.Next()

		// afrer exec

	}
}
