// Package service
// @author: xs
// @date: 2022/5/17
// @Description: service 相关依赖&提供注册点
package service

import (
	v1Apidemo "ghub/internal/service/v1/apidemo"
	v1Signup "ghub/internal/service/v1/auth"
	service "ghub/internal/service/v1/helloword"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	service.NewGreeterService,
	//helloword.NewGreeterClient,
	v1Signup.NewSignupService,
	v1Apidemo.NewApidemoService,
)
