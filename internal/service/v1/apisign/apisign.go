package service

import (
	"fmt"
	pb "ghub/api/v1/apisign"
	"github.com/china-xs/gin-tpl/pkg/api_sign"
	"github.com/china-xs/gin-tpl/pkg/jwt_auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/parkingwang/go-sign"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

type ApisignService struct {
	pb.UnimplementedApisignServer
	V *viper.Viper
}

func NewApisignService(v *viper.Viper) *ApisignService {
	return &ApisignService{
		V: v,
	}
}

func (s *ApisignService) ApisignGetDemo(ctx *gin.Context, req *pb.SignRequest) (*pb.SignReply, error) {
	return &pb.SignReply{}, nil
}

//生成api 签名demo
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

//生成token demo
func (s *ApisignService) CreateTokenDemo(ctx *gin.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	options, err := jwt_auth.NewOps(s.V)
	if err != nil {
		return nil, err
	}
	token, err := createToken(options.Secret, map[string]interface{}{
		"user_id": "256",
	}, 3600)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTokenReply{
		Token: token,
	}, nil

}

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
