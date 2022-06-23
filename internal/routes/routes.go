// Package routes
// @author: xs
// @date: 2022/5/17
// @Description: routes
package routes

import (
	pV1apidemo "ghub/api/v1/apidemo"
	pV1signup "ghub/api/v1/auth"
	"ghub/api/v1/helloword"
	v1Apidemo "ghub/internal/service/v1/apidemo"
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
	V1Apidemo *v1Apidemo.ApidemoService
}

func (r *Routes) InitRoutes(app *tpl.Server) {
	helloword.RegisterGreeterGinServer(app, r.HelloSrv)
	//query.SetDefault(r.Db)
	pV1signup.RegisterSignupGinServer(app, r.V1Signup)
	pV1apidemo.RegisterApidemoGinServer(app, r.V1Apidemo)
}
