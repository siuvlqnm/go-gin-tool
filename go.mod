module github.com/siuvlqnm/go-gin-example

go 1.14

require (
	github.com/astaxie/beego v1.12.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200805065543-0cf7623e9dbd // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/siuvlqnm/go-gin-example/conf => /app/go/src/go-gin-example/pkg/conf
	github.com/siuvlqnm/go-gin-example/middleware => /app/go/src/go-gin-example/middleware
	github.com/siuvlqnm/go-gin-example/models => /app/go/src/go-gin-example/models
	github.com/siuvlqnm/go-gin-example/pkg/setting => /app/go/src/go-gin-example/pkg/setting
	github.com/siuvlqnm/go-gin-example/routers => /app/go/src/go-gin-example/routers
)
