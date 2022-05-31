// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"ghub/internal/data/dao/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newUser2role(db *gorm.DB) user2role {
	_user2role := user2role{}

	_user2role.user2roleDo.UseDB(db)
	_user2role.user2roleDo.UseModel(&model.User2role{})

	tableName := _user2role.user2roleDo.TableName()
	_user2role.ALL = field.NewField(tableName, "*")
	_user2role.ID = field.NewInt32(tableName, "id")
	_user2role.AccountID = field.NewInt32(tableName, "account_id")
	_user2role.RoleID = field.NewInt32(tableName, "role_id")
	_user2role.OperateID = field.NewInt32(tableName, "operate_id")

	_user2role.fillFieldMap()

	return _user2role
}

type user2role struct {
	user2roleDo user2roleDo

	ALL       field.Field
	ID        field.Int32
	AccountID field.Int32
	RoleID    field.Int32
	OperateID field.Int32

	fieldMap map[string]field.Expr
}

func (u user2role) Table(newTableName string) *user2role {
	u.user2roleDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user2role) As(alias string) *user2role {
	u.user2roleDo.DO = *(u.user2roleDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user2role) updateTableName(table string) *user2role {
	u.ALL = field.NewField(table, "*")
	u.ID = field.NewInt32(table, "id")
	u.AccountID = field.NewInt32(table, "account_id")
	u.RoleID = field.NewInt32(table, "role_id")
	u.OperateID = field.NewInt32(table, "operate_id")

	u.fillFieldMap()

	return u
}

func (u *user2role) WithContext(ctx context.Context) *user2roleDo {
	return u.user2roleDo.WithContext(ctx)
}

func (u user2role) TableName() string { return u.user2roleDo.TableName() }

func (u *user2role) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user2role) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["account_id"] = u.AccountID
	u.fieldMap["role_id"] = u.RoleID
	u.fieldMap["operate_id"] = u.OperateID
}

func (u user2role) clone(db *gorm.DB) user2role {
	u.user2roleDo.ReplaceDB(db)
	return u
}

type user2roleDo struct{ gen.DO }

func (u user2roleDo) Debug() *user2roleDo {
	return u.withDO(u.DO.Debug())
}

func (u user2roleDo) WithContext(ctx context.Context) *user2roleDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u user2roleDo) Clauses(conds ...clause.Expression) *user2roleDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u user2roleDo) Returning(value interface{}, columns ...string) *user2roleDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u user2roleDo) Not(conds ...gen.Condition) *user2roleDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u user2roleDo) Or(conds ...gen.Condition) *user2roleDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u user2roleDo) Select(conds ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u user2roleDo) Where(conds ...gen.Condition) *user2roleDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u user2roleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *user2roleDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u user2roleDo) Order(conds ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u user2roleDo) Distinct(cols ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u user2roleDo) Omit(cols ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u user2roleDo) Join(table schema.Tabler, on ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u user2roleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u user2roleDo) RightJoin(table schema.Tabler, on ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u user2roleDo) Group(cols ...field.Expr) *user2roleDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u user2roleDo) Having(conds ...gen.Condition) *user2roleDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u user2roleDo) Limit(limit int) *user2roleDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u user2roleDo) Offset(offset int) *user2roleDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u user2roleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *user2roleDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u user2roleDo) Unscoped() *user2roleDo {
	return u.withDO(u.DO.Unscoped())
}

func (u user2roleDo) Create(values ...*model.User2role) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u user2roleDo) CreateInBatches(values []*model.User2role, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u user2roleDo) Save(values ...*model.User2role) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u user2roleDo) First() (*model.User2role, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2role), nil
	}
}

func (u user2roleDo) Take() (*model.User2role, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2role), nil
	}
}

func (u user2roleDo) Last() (*model.User2role, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2role), nil
	}
}

func (u user2roleDo) Find() ([]*model.User2role, error) {
	result, err := u.DO.Find()
	return result.([]*model.User2role), err
}

func (u user2roleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User2role, err error) {
	buf := make([]*model.User2role, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u user2roleDo) FindInBatches(result *[]*model.User2role, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u user2roleDo) Attrs(attrs ...field.AssignExpr) *user2roleDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u user2roleDo) Assign(attrs ...field.AssignExpr) *user2roleDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u user2roleDo) Joins(field field.RelationField) *user2roleDo {
	return u.withDO(u.DO.Joins(field))
}

func (u user2roleDo) Preload(field field.RelationField) *user2roleDo {
	return u.withDO(u.DO.Preload(field))
}

func (u user2roleDo) FirstOrInit() (*model.User2role, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2role), nil
	}
}

func (u user2roleDo) FirstOrCreate() (*model.User2role, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User2role), nil
	}
}

func (u user2roleDo) FindByPage(offset int, limit int) (result []*model.User2role, count int64, err error) {
	if limit <= 0 {
		count, err = u.Count()
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u user2roleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *user2roleDo) withDO(do gen.Dao) *user2roleDo {
	u.DO = *do.(*gen.DO)
	return u
}
