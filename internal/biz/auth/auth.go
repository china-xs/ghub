// Package auth
// @author: xs
// @date: 2022/6/6
// @Description: auth
package auth

import (
	"context"
	pb "ghub/api/v1/auth"
	"ghub/internal/data/account"
	"ghub/internal/data/dao/model"
	"ghub/internal/data/role"
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Biz struct {
	log      *zap.Logger
	tx       db.Transaction
	repo     *account.Repo
	roleRepo *role.Repo
}

func NewBiz(log *zap.Logger, tx db.Transaction,
	repo *account.Repo,
	roleRepo *role.Repo,
) *Biz {
	return &Biz{
		log:      log,
		tx:       tx,
		repo:     repo,
		roleRepo: roleRepo,
	}
}

//
// UsingEmail
// @Description: 创建账号&角色不存在并创建角色数据，存在保存即可
// @receiver biz
// @param ctx
// @param in
// @return *model.Account
// @return error
//
func (biz *Biz) UsingEmail(ctx context.Context, in *pb.UsingEmailRequest) (*model.Account, error) {
	var fns = make([]func(dao gen.Dao) gen.Dao, 1)
	fns[0] = account.QueryEmail(in.Email)
	_, err := biz.repo.First(ctx)
	if err == nil || (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) {
		return nil, errors.New("email已存在，请登录")
	}
	rfns := make([]func(dao gen.Dao) gen.Dao, 1)
	l := len(in.Roles)
	m := make(map[string]struct{}, l)
	for i := 0; i < l; i++ {
		name := in.Roles[i]
		m[name] = struct{}{}
	}
	rfns[0] = role.QueryNames(in.Roles)
	roles, err := biz.roleRepo.Find(ctx, rfns...)
	var addRoles []model.Role
	for i := 0; i < len(roles); i++ {
		role := roles[i]
		if _, ok := m[role.Name]; ok {
			delete(m, role.Name)
		}
		r := *role
		addRoles = append(addRoles, r)
	}
	var nRoles []*model.Role
	for name, _ := range m {
		nRoles = append(nRoles, &model.Role{
			Name:  name,
			Nodes: "",
			Desc:  name,
		})
	}

	a := model.Account{
		Username: in.Username,
		Email:    in.Email,
	}
	// 开启事务 案例
	biz.tx.ExecTx(ctx, func(ctx context.Context) error {
		if len(nRoles) > 0 {
			if err := biz.roleRepo.Create(ctx, nRoles...); err != nil {
				return err
			}
			for _, role := range nRoles {
				r := *role
				addRoles = append(addRoles, r)
			}
		}
		a.Roles = addRoles
		if err := biz.repo.Create(ctx, &a); err != nil {
			return err
		}
		return nil
	})

	return nil, err
}
