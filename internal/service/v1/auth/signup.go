package service

import (
	"fmt"
	"ghub/api/common"
	"ghub/internal/biz/auth"
	"ghub/internal/cache/role"
	"ghub/internal/data/account"
	roleRepo "ghub/internal/data/role"
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"

	pb "ghub/api/v1/auth"
)

const _signLogUpKey = "sign_controller"

type SignupService struct {
	pb.UnimplementedSignupServer
	log         *zap.Logger // 1 定义log 日志包
	cache       *role.Cache
	accountRepo *account.Repo
	roleRepo    *roleRepo.Repo
	tx          db.Transaction
	biz         *auth.Biz
}

func NewSignupService(log *zap.Logger,
	tx db.Transaction,
	accountRepo *account.Repo,
	roleRepo *roleRepo.Repo,
	cache *role.Cache,
	biz *auth.Biz,
) *SignupService {
	return &SignupService{
		log:         log,
		cache:       cache,
		accountRepo: accountRepo,
		roleRepo:    roleRepo,
		tx:          tx,
		biz:         biz,
	}
}

func (s *SignupService) UsingEmail(c *gin.Context, req *pb.UsingEmailRequest) (*pb.UsingEmailReply, error) {
	ctx := c.Request.Context()
	log.WithCtx(ctx, s.log).Info(_signLogUpKey,
		zap.String("email", req.Email))

	// 事务
	if _, err := s.biz.UsingEmail(ctx, req); err != nil {
		return nil, err
	}

	//user := new(model.Account)
	//user.Username = req.Username
	//user.Email = req.Email
	//user.Pwd = req.Pwd
	//var fns = make([]func(gen.Dao) gen.Dao, 2)
	//fns[0] = account.QueryId(1)
	//fns[1] = account.PLRoles()
	//u, err := s.accountRepo.First(ctx, fns...)
	//var fns1 = make([]func(gen.Dao) gen.Dao, 2)
	//
	//fns1[0] = account.QueryId(1)
	//fns1[1] = account.Select([]string{"state", "pwd"})
	//
	//u.State = 0
	//u.Phone = "134278765545"
	//u.Pwd=req.Pwd
	//s.accountRepo.Save(ctx, u, fns1...)
	//
	//fmt.Printf("err:%v\n", err)
	//fmt.Printf("user:%v\n", u.Roles)
	return &pb.UsingEmailReply{
		Account: &common.AccountSimple{
			Id:       1,
			Email:    req.Email,
			Username: "",
			Phone:    "",
		},
		Token: &common.AuthToken{
			AccessToken: "",
			ExpireAt:    time.Now().Format(time.RFC3339),
		},
	}, nil
}
func (s *SignupService) UsingPhone(c *gin.Context, req *pb.UsingPhoneRequest) (*pb.UsingPhoneReply, error) {
	ctx := c.Request.Context()
	cache, err := s.cache.GetRoleByIds(ctx, []int32{1, 2})
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cache); i++ {
		m := cache[i]
		fmt.Printf("m:%v\n", m.ID)
	}
	//fmt.Printf("cache:%v\n",cache)
	//s.tx.ExecTx(ctx, func(ctx context.Context) error {
	//	//s.roleRepo.Create()
	//	return nil
	//})

	return &pb.UsingPhoneReply{}, nil
}
