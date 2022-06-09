// Package routes
// @author: xs
// @date: 2022/5/17
// @Description: routes
package routes

import (
	pV1apisign "ghub/api/v1/apisign"
	pV1signup "ghub/api/v1/auth"
	"ghub/api/v1/helloword"
	v1Apisign "ghub/internal/service/v1/apisign"
	v1Signup "ghub/internal/service/v1/auth"
	service "ghub/internal/service/v1/helloword"
	tpl "github.com/china-xs/gin-tpl"
	"github.com/google/wire"
)

// ProviderSet 依赖提供
var ProviderSet = wire.NewSet(wire.Struct(new(Routes), "*"))

//
// Routes
// @Description: 定义Routes 公共路由注册struct 可以不定义，主要是为了方便后续代码多了 初始化调用不堆叠参数
//
type Routes struct {
	HelloSrv *service.GreeterService
	//Helloword helloword.GreeterGinServer
	//GreeterService
	//Db *gorm.DB
	V1Signup  *v1Signup.SignupService
	V1Apisign *v1Apisign.ApisignService
}

func (r *Routes) InitRoutes(app *tpl.Server) {
	helloword.RegisterGreeterGinServer(app, r.HelloSrv)
	//query.SetDefault(r.Db)
	pV1signup.RegisterSignupGinServer(app, r.V1Signup)
	pV1apisign.RegisterApisignGinServer(app, r.V1Apisign)
}
