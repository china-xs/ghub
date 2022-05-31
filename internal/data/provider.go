// Package data
// @author: xs
// @date: 2022/5/18
// @Description: data
package data

import (
	"ghub/internal/data/account"
	"ghub/internal/data/dao/query"
	"ghub/internal/data/role"
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	db.New,
	db.NewDb,
	StepUp,
	NewTx,
	account.NewRepo,
	role.NewRepo,
)

func StepUp(gdb *gorm.DB, l *zap.Logger) (*db.Data, error) {
	query.SetDefault(gdb)
	d, _, err := db.NewData(gdb, l)
	return d, err
}

func NewTx(d *db.Data) db.Transaction {
	return db.Transaction(d)
}
