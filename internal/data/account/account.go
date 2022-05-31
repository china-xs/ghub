// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
package account

import (
	"context"
	"ghub/internal/data/dao/model"
	"ghub/internal/data/dao/query"
	"github.com/china-xs/gin-tpl/pkg/db"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Repo struct {
	data *db.Data
	log  *zap.Logger
}

//
// NewRepo
// @Description: 初始化Account Repo
// @param data
// @param log
// @return *Repo
//
func NewRepo(data *db.Data, log *zap.Logger) *Repo {
	return &Repo{
		data: data,
		log:  log,
	}
}

//
// First
// @Description: 单个数据查询
// @receiver repo
// @param ctx
// @param fns
// @return *model.Account
// @return error
//
func (repo *Repo) First(ctx context.Context, fns ...func(dao gen.Dao) gen.Dao) (*model.Account, error) {
	return query.Account.WithContext(ctx).Scopes(fns...).First()
}

//
// Find
// @Description: 多条查询
// @receiver repo
// @param ctx
// @param fns
// @return []*model.Account
// @return error
//
func (repo Repo) Find(ctx context.Context, fns ...func(dao gen.Dao) gen.Dao) ([]*model.Account, error) {
	return query.Account.WithContext(ctx).Scopes(fns...).Find()
}

//
// Create
// @Description: 增删改会涉及事物操作 不可以直接
// @receiver repo
// @param ctx
// @param account
// @return error
//
func (repo Repo) Create(ctx context.Context, account *model.Account) error {
	db := repo.data.DB(ctx)
	r := query.Use(db).Account
	return r.WithContext(ctx).Create(account)
}

//
// Save
// @Description: 更新
// @receiver repo
// @param ctx
// @param account
// @param fns
// @return error
//
func (repo *Repo) Save(ctx context.Context, account *model.Account, fns ...func(gen.Dao) gen.Dao) error {
	db := repo.data.DB(ctx)
	r := query.Use(db).Account
	l := len(fns)
	if l == 0 && account.ID > 0 {
		fns = append(fns, func(tx gen.Dao) gen.Dao {
			return tx.Where(r.ID.Eq(account.ID))
		})
		l = 1
	}
	if l == 0 {
		return gorm.ErrMissingWhereClause
	}
	_, err := r.WithContext(ctx).Scopes(fns...).Omit(r.ID).Updates(account)
	return err
}

//
// Delete
// @Description:
// @receiver repo
// @param ctx
// @param fns
// @return error
//
func (repo Repo) Delete(ctx context.Context, fns ...func(gen.Dao) gen.Dao) error {
	if len(fns) == 0 {
		return gorm.ErrMissingWhereClause
	}
	db := repo.data.DB(ctx)
	r := query.Use(db).Account
	_, err := r.WithContext(ctx).Scopes(fns...).Delete()
	return err
}

//
// Count
// @Description:
// @receiver repo
// @param ctx
// @param fns...func(gen.Dao)gen.Dao
// @return int64
// @return error
//
func (repo Repo) Count(ctx context.Context, fns ...func(gen.Dao) gen.Dao) (int64, error) {
	return query.Account.WithContext(ctx).Scopes(fns...).Count()
}

//
// Value
// @Description: 查询单个字段,注意后续迭代字段类型变更导致获取参数失败问题
// @receiver repo
// @param ctx
// @param fieldName
// @param fns
// @return interface{}
// @return error
//
func (repo Repo) Value(ctx context.Context, fieldName string, fns ...func(gen.Dao) gen.Dao) (interface{}, error) {
	var result map[string]interface{}
	q := query.Account
	column, isOk := q.GetFieldByName(fieldName)
	if !isOk {
		return nil, gorm.ErrInvalidField
	}
	q.WithContext(ctx).Scopes(fns...).Select(column.As("value")).Scan(&result)
	if v, isOk := result["value"]; isOk {
		return v, nil
	}
	return nil, gorm.ErrRecordNotFound
}

//
// Pluck
// @Description: 暂时无法获取到字段类型只能在调用
// @receiver repo
// @param ctx
// @param fieldName
// @param fns...func(gen.Dao)
// @return []interface{}
// @return error
//
func (repo Repo) Pluck(ctx context.Context, fieldName string, fns ...func(gen.Dao) gen.Dao) ([]interface{}, error) {
	//var result map[string]interface{}
	var result []interface{}
	q := query.Account
	column, isOk := q.GetFieldByName(fieldName)
	if !isOk {
		return nil, gorm.ErrInvalidField
	}
	q.WithContext(ctx).Scopes(fns...).Pluck(column, &result)
	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return result, nil
}

//=======================================WHERE==========================================================

//
// QueryId
// @Description: 单个查询条件
// @param id
// @return func(gen.Dao) gen.Dao
//
func QueryId(id int32) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if id <= 0 {
			return tx
		}
		return tx.Where(query.Account.ID.Eq(id))
	}
}

//
// QueryIds
// @Description: 批量查询
// @param ids
// @return func(gen.Dao) gen.Dao
//
func QueryIds(ids []int32) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		if len(ids) == 0 {
			return tx
		}
		return tx.Where(query.Account.ID.In(ids...))
	}
}

//
// Select
// @Description: 选择更新字段
// @param fields
// @return func(gen.Dao) gen.Dao
//
func Select(fields []string) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		q := query.Account
		l := len(fields)
		var columns = make([]field.Expr, l)
		for i := 0; i < l; i++ {
			column, isOk := q.GetFieldByName(fields[i])
			if !isOk {
				continue
			}
			columns[i] = column
		}
		return tx.Select(columns...)
	}
}

//
// Omit
// @Description: 忽略更新字段
// @param fields
// @return func(gen.Dao) gen.Dao
//
func Omit(fields []string) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		q := query.Account
		l := len(fields)
		var columns = make([]field.Expr, l)
		for i := 0; i < l; i++ {
			column, isOk := q.GetFieldByName(fields[i])
			if !isOk {
				continue
			}
			columns[i] = column
		}
		return tx.Omit(columns...)
	}
}
