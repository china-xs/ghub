// Package role
// @author: xs
// @date: 2022/6/6
// @Description: role
package role

import (
	"ghub/internal/data/dao/query"
	"gorm.io/gen"
)

func QueryNames(names []string) func(gen.Dao) gen.Dao {
	return func(tx gen.Dao) gen.Dao {
		return tx.Where(query.Role.Name.In(names...))
	}
}
