package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myproject/swagger-test/controller"
	"myproject/swagger-test/middleware"

	// _ "myproject/swagger-test/docs/swagger"
	"myproject/swagger-test/httputil"
	"net/http"
)

// https://github.com/swaggo/swag/tree/master/example/celler
func main1() {
	r := gin.New()

	// add skipper and middleware
	skipperPath := []string{"/metrics", "/swagger"}
	skipper := middleware.AllowPathPrefixSkipper(skipperPath...)
	r.Use(middleware.UserAuthMiddleware(skipper))

	c := controller.NewController()
	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
		bottles := v1.Group("/bottles")
		{
			bottles.GET(":id", c.ShowBottle)
			bottles.GET("", c.ListBottles)
		}
		admin := v1.Group("/admin")
		{
			admin.Use(auth())
			admin.POST("/auth", c.Auth)
		}
		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
			examples.GET("calc", c.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c.PathParamsExample)
			examples.GET("header", c.HeaderExample)
			examples.GET("securities", c.SecuritiesExample)
			examples.GET("attribute", c.AttributeExample)
		}
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/swagger", "swagger-test/docs/swagger") // OK
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
