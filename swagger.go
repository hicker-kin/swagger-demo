package main

/*
Package router 生成swagger文档
文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format
访问：http://localhost:8080/swagger/index.html

使用方式：
	go get -u github.com/swag-go/swag/cmd/swag
	swag init -g swagger.go -o ./docs/swagger
*/

// @title my-proj
// @version 1.0.0
// @description 云平台业务框架
// @schemes http https
// @basePath /api/v1
// @contact.name qzj
// @contact.email qzj@163.com
// @host localhost:8080

/*basic mean is user and password*/
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Description for what is this security definition being used
