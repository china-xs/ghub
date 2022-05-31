// Package account
// @author: xs
// @date: 2022/5/20
// @Description: account
package account

import (
	"ghub/internal/data/dao/query"
	"gorm.io/gen"
)

//
// PLRoles
// @Description: 预加载
// @return func(gen.Dao) gen.Dao
//
func PLRoles() func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		a := query.Account
		return tx.Preload(a.Roles).Select()
	}
}
