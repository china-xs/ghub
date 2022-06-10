package service

import (
	"errors"
	"fmt"
	pb "ghub/api/v1/apisign"
	"github.com/china-xs/gin-tpl/pkg/api_sign"
	"github.com/china-xs/gin-tpl/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/parkingwang/go-sign"
	"github.com/spf13/viper"
	"net/url"
	"strconv"
	"time"
)

type ApisignService struct {
	pb.UnimplementedApisignServer
	V *viper.Viper
}

const (
	JWT_USER_ID = "jwt_user_id"
)

func NewApisignService(v *viper.Viper) *ApisignService {
	return &ApisignService{
		V: v,
	}
}

//验证接口签名 demo
///api/v1/get-sign?name=李先后&id=12&time_stamp=1654756018&nonce_str=s1231sdfsdhfwerwe&appid=1231231abc&sign=05d7e4b9cbfe26b037421cd228089c67
func (s *ApisignService) ApisignCheckDemo(ctx *gin.Context, req *pb.SignRequest) (*pb.SignReply, error) {
	return &pb.SignReply{}, nil
}

//生成api 签名demo
///api/v1/get-sign?name=李先后&id=12&time_stamp=1654756018&nonce_str=s1231sdfsdhfwerwe&appid=1231231abc
func (s *ApisignService) CreateSignDemo(ctx *gin.Context, req *pb.CreateSignRequest) (*pb.CreateSignReply, error) {
	options, err := api_sign.NewOps(s.V)
	if err != nil {
		return nil, err
	}

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

	signer.SetAppSecretWrapBody(options.Secret)
	return &pb.CreateSignReply{
		Sign: signer.GetSignedQuery(),
	}, nil
}

//生成jwt-token demo
func (s *ApisignService) CreateTokenDemo(ctx *gin.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	options, err := jwt_auth.NewOps(s.V)
	if err != nil {
		return nil, err
	}
	token, err := createToken(options.Secret, map[string]interface{}{
		JWT_USER_ID: "256",
	}, 3600)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenReply{
		Token: token,
	}, nil

}

//生成jwt-token
func createToken(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds //过期时间
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

//解析jwt-token中的登录信息
func (s *ApisignService) GetTokenInfo(ctx *gin.Context, req *pb.GetTokenInfoRequest) (*pb.GetTokenInfoReply, error) {
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
