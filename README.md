# swagger-demo
go swagger demo

Package router 生成swagger文档    
文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format  
访问：http://localhost:8080/swagger/index.html  

使用方式：
	go get -u github.com/swag-go/swag/cmd/swag  
	swag init -g swagger.go -o ./docs/swagger
