// Package cache
// @author: xs
// @date: 2022/5/27
// @Description: cache
package cache

import (
	"ghub/internal/cache/role"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	role.NewCache,
)
