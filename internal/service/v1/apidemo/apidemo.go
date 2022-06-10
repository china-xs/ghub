package service

import (
	"errors"
	"fmt"
	pb "ghub/api/v1/apidemo"
	"github.com/china-xs/gin-tpl/pkg/api_sign"
	"github.com/china-xs/gin-tpl/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
	"github.com/parkingwang/go-sign"
	"net/url"
	"strconv"
)

type ApidemoService struct {
	pb.UnimplementedApidemoServer
	authOptions *jwt_auth.Options
	signOptions *api_sign.Options
	jwtAuth     *jwt_auth.JwtAuth
}

const (
	JWT_USER_ID = "jwt_user_id"
)

func NewApidemoService(jo *jwt_auth.Options, ja *jwt_auth.JwtAuth, ao *api_sign.Options) *ApidemoService {
	return &ApidemoService{
		authOptions: jo,
		signOptions: ao,
		jwtAuth:     ja,
	}
}

//验证接口签名 demo
///api/v1/get-sign?name=李先后&id=12&time_stamp=1654756018&nonce_str=s1231sdfsdhfwerwe&appid=1231231abc&sign=05d7e4b9cbfe26b037421cd228089c67
func (s *ApidemoService) ApisignCheckDemo(ctx *gin.Context, req *pb.SignRequest) (*pb.SignReply, error) {
	return &pb.SignReply{}, nil
}

//生成api 签名demo
///api/v1/get-sign?name=李先后&id=12&time_stamp=1654756018&nonce_str=s1231sdfsdhfwerwe&appid=1231231abc
func (s *ApidemoService) CreateSignDemo(ctx *gin.Context, req *pb.CreateSignRequest) (*pb.CreateSignReply, error) {
	signer := sign.NewGoSignerMd5()
	signer.SetAppId(req.Appid)
	signer.SetTimeStamp(req.TimeStamp)
	signer.SetNonceStr(req.NonceStr)
	values, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		return nil, err
	}
	for k, v := range values {
		switch k {
		case "time_stamp", "nonce_str", "appid", "sign":
		default:
			signer.AddBodies(k, v)
		}
	}

	fmt.Printf("source:%s\n", signer.GetSignBodyString())

	signer.SetAppSecretWrapBody(s.signOptions.Secret)
	return &pb.CreateSignReply{
		Sign: signer.GetSignedQuery(),
	}, nil
}

//生成jwt-token demo
func (s *ApidemoService) CreateTokenDemo(ctx *gin.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	token, err := s.jwtAuth.CreateTokenWithMapPayload(map[string]interface{}{
		JWT_USER_ID: "256",
	}, 3600)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenReply{
		Token: token,
	}, nil

}

//解析jwt-token中的登录信息
func (s *ApidemoService) GetTokenInfo(ctx *gin.Context, req *pb.GetTokenInfoRequest) (*pb.GetTokenInfoReply, error) {
	userId, err := GetJwtUserIdFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetTokenInfoReply{
		UserId: userId,
	}, nil
}

//获取登陆后的user_id demo
func GetJwtUserIdFromCtx(ctx *gin.Context) (int64, error) {
	_userId := ctx.GetString(JWT_USER_ID)
	userId, err := strconv.ParseInt(_userId, 10, 64)
	if err != nil {
		return 0, errors.New("未登录")
	}
	return userId, nil
}
