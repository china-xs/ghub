// Package account
// @author: xs
// @date: 2022/6/6
// @Description: account
package account

import (
	"ghub/internal/data/dao/query"
	"gorm.io/gen"
)

func QueryEmail(email string) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		return tx.Where(query.Account.Email.Eq(email))
	}
}

func QueryEmailOrPhone(email, phone string) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		return tx.Where(query.Account.Email.Eq(email)).Or(query.Account.Phone.Eq(phone))
	}
}
