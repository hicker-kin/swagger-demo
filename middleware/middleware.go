package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"myproject/swagger-test/httputil"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		/*sah := &sso.SessionHandle{}
		status, err := sah.SessionAuth(c, true)
		if err != nil {
			ginplus.ResErrorWithStatus(c, err, status)
			return
		}*/
		if c.GetHeader("Authorization") != "123" {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("invalid Authorization"))
			return
		}
		c.Next()
	}
}
