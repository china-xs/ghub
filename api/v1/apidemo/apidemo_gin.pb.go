// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:v1.0.1
// protoc-gen-go-gin v1.0.1

package apidemo

import (
	gin_tpl "github.com/china-xs/gin-tpl"
	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// gin.gin_tpl.binding.strings.
type ApidemoGinServer interface {
	ApisignCheckDemo(*gin.Context, *SignRequest) (*SignReply, error)
	CreateSignDemo(*gin.Context, *CreateSignRequest) (*CreateSignReply, error)
	CreateTokenDemo(*gin.Context, *CreateTokenRequest) (*CreateTokenReply, error)
	GetTokenInfo(*gin.Context, *GetTokenInfoRequest) (*GetTokenInfoReply, error)
}

func RegisterApidemoGinServer(s *gin_tpl.Server, srv ApidemoGinServer, ms ...gin.HandlerFunc) {
	route := s.Engine.Use(ms...)
	route.GET("/api/v1/get-sign", _Apidemo_ApisignCheckDemo0_Gin_Handler(s, srv))
	route.GET("/api/v1/create-sign", _Apidemo_CreateSignDemo0_Gin_Handler(s, srv))
	route.GET("/api/v1/create-token", _Apidemo_CreateTokenDemo0_Gin_Handler(s, srv))
	route.GET("/api/v1/get-token-info", _Apidemo_GetTokenInfo0_Gin_Handler(s, srv))
}

func _Apidemo_ApisignCheckDemo0_Gin_Handler(s *gin_tpl.Server, srv ApidemoGinServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var in SignRequest
		switch c.Request.Method {
		case "POST", "PUT":
			if err := c.ShouldBindBodyWith(&in, binding.JSON); err != nil {
				s.Enc(c, nil, err)
				return
			}
			if strings.Contains(c.Request.URL.String(), "?") {
				if err := c.ShouldBindQuery(&in); err != nil {
					s.Enc(c, nil, err)
					return
				}
			}

		case "GET", "DELETE":
			if err := c.ShouldBindQuery(&in); err != nil {
				s.Enc(c, nil, err)
				return
			}
		}
		c.Set(gin_tpl.OperationKey, "/api.v1.apidemo.Apidemo/ApisignCheckDemo")
		h := s.Middleware(func(c *gin.Context, req interface{}) (interface{}, error) {
			return srv.ApisignCheckDemo(c, req.(*SignRequest))
		})
		out, err := h(c, &in)
		s.Enc(c, out, err)
		return
	}
}

func _Apidemo_CreateSignDemo0_Gin_Handler(s *gin_tpl.Server, srv ApidemoGinServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var in CreateSignRequest
		switch c.Request.Method {
		case "POST", "PUT":
			if err := c.ShouldBindBodyWith(&in, binding.JSON); err != nil {
				s.Enc(c, nil, err)
				return
			}
			if strings.Contains(c.Request.URL.String(), "?") {
				if err := c.ShouldBindQuery(&in); err != nil {
					s.Enc(c, nil, err)
					return
				}
			}

		case "GET", "DELETE":
			if err := c.ShouldBindQuery(&in); err != nil {
				s.Enc(c, nil, err)
				return
			}
		}
		c.Set(gin_tpl.OperationKey, "/api.v1.apidemo.Apidemo/CreateSignDemo")
		h := s.Middleware(func(c *gin.Context, req interface{}) (interface{}, error) {
			return srv.CreateSignDemo(c, req.(*CreateSignRequest))
		})
		out, err := h(c, &in)
		s.Enc(c, out, err)
		return
	}
}

func _Apidemo_CreateTokenDemo0_Gin_Handler(s *gin_tpl.Server, srv ApidemoGinServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var in CreateTokenRequest
		switch c.Request.Method {
		case "POST", "PUT":
			if err := c.ShouldBindBodyWith(&in, binding.JSON); err != nil {
				s.Enc(c, nil, err)
				return
			}
			if strings.Contains(c.Request.URL.String(), "?") {
				if err := c.ShouldBindQuery(&in); err != nil {
					s.Enc(c, nil, err)
					return
				}
			}

		case "GET", "DELETE":
			if err := c.ShouldBindQuery(&in); err != nil {
				s.Enc(c, nil, err)
				return
			}
		}
		c.Set(gin_tpl.OperationKey, "/api.v1.apidemo.Apidemo/CreateTokenDemo")
		h := s.Middleware(func(c *gin.Context, req interface{}) (interface{}, error) {
			return srv.CreateTokenDemo(c, req.(*CreateTokenRequest))
		})
		out, err := h(c, &in)
		s.Enc(c, out, err)
		return
	}
}

func _Apidemo_GetTokenInfo0_Gin_Handler(s *gin_tpl.Server, srv ApidemoGinServer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var in GetTokenInfoRequest
		switch c.Request.Method {
		case "POST", "PUT":
			if err := c.ShouldBindBodyWith(&in, binding.JSON); err != nil {
				s.Enc(c, nil, err)
				return
			}
			if strings.Contains(c.Request.URL.String(), "?") {
				if err := c.ShouldBindQuery(&in); err != nil {
					s.Enc(c, nil, err)
					return
				}
			}

		case "GET", "DELETE":
			if err := c.ShouldBindQuery(&in); err != nil {
				s.Enc(c, nil, err)
				return
			}
		}
		c.Set(gin_tpl.OperationKey, "/api.v1.apidemo.Apidemo/GetTokenInfo")
		h := s.Middleware(func(c *gin.Context, req interface{}) (interface{}, error) {
			return srv.GetTokenInfo(c, req.(*GetTokenInfoRequest))
		})
		out, err := h(c, &in)
		s.Enc(c, out, err)
		return
	}
}
