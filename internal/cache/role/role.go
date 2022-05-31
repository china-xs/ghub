// Package role
// @author: xs
// @date: 2022/5/19
// @Description: role 角色缓存数据,数据更新
package role

import (
	"context"
	"encoding/json"
	"fmt"
	"ghub/internal/data/dao/model"
	roleRepo "ghub/internal/data/role"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const (
	_logKey = "cache-rule"
	// 默认缓存啊
	_cacheKey = "ghub:rule:%v"
	_expire   = 30
	//_longtime =
)

type Cache struct {
	log  *zap.Logger
	rdb  *redis.Client
	repo *roleRepo.Repo
}

func GetCacheKeyById(id int32) string {
	return fmt.Sprintf(_cacheKey, id)
}

//
// NewCache
// @Description: 初始化缓存层
// @param log
// @param rdb
// @param repo
// @return *Cache
//
func NewCache(log *zap.Logger, rdb *redis.Client, repo *roleRepo.Repo) *Cache {
	return &Cache{
		log:  log,
		rdb:  rdb,
		repo: repo,
	}
}

//
// GetRoleById
// @Description: 单个拉去数据
// @receiver cache
// @param ctx
// @param id
//
func (cache *Cache) GetRoleById(ctx context.Context, id int32) (role *model.Role, err error) {
	key := fmt.Sprintf(_cacheKey, id)
	if rest, err := cache.rdb.Get(ctx, key).Result(); !errors.Is(err, redis.Nil) {
		err = json.Unmarshal([]byte(rest), role)
		return nil, err
	}
	role, err = cache.repo.First(ctx)
	if err != nil {
	}
	buf, err := json.Marshal(role)
	if err != nil {
		return
	}
	// 考虑， 存byte 还是存字符串
	expired := time.Duration(_expire+rand.Intn(100)) * time.Minute
	if err := cache.rdb.Set(ctx, key, string(buf), expired).Err(); err != nil {
		log.WithCtx(ctx, cache.log).Warn(_logKey,
			zap.Int32("id", id),
			zap.Error(err),
		)
		return role, nil
	}
	cache.log.Info(_logKey,
		zap.String("cache-key", key),
		zap.String("cache-val", string(buf)),
		zap.String("expired", expired.String()),
	)
	return
}

//
// GetRoleByIds
// @Description: 批量拉取数据
// @receiver cache
// @param ctx
// @param ids
// @return roles
// @return err
//
func (cache *Cache) GetRoleByIds(ctx context.Context, ids []int32) (roles []*model.Role, err error) {
	l := len(ids)
	var keys = make([]string, l)
	for i := 0; i < l; i++ {
		id := ids[i]
		keys[i] = fmt.Sprintf(_cacheKey, id)
	}
	// 待定处理
	res := cache.rdb.MGet(ctx, keys...).Val()
	var unCacheIds []int32
	for i := 0; i < l; i++ {
		str := res[i]
		if str == nil {
			unCacheIds = append(unCacheIds, ids[i])
		} else {
			var m = new(model.Role)
			json.Unmarshal([]byte(str.(string)), m)
			roles = append(roles, m)
		}
	}
	if len(unCacheIds) == 0 {
		return roles, nil
	}
	roles1, err := cache.repo.Find(ctx, roleRepo.QueryIds(unCacheIds))
	if err == nil {
		l1 := len(roles1)
		var mv []string
		for i := 0; i < l1; i++ {
			m := roles1[i]
			k := fmt.Sprintf(_cacheKey, m.ID)
			if v, err := json.Marshal(m); err == nil {
				mv = append(mv, k, string(v))
			}
			roles = append(roles, m)
		}
		cache.rdb.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
			l3 := len(mv)
			for i := 0; i < l3; i += 2 {
				k := mv[i]
				v := mv[i+1]
				t := time.Duration(_expire+rand.Intn(100)) * time.Minute

				if err := pipeliner.Set(ctx, k, v, t).Err(); err != nil {
					return err
				} else {
					cache.log.Info(_logKey,
						zap.String("cache-key", k),
						zap.String("cache-val", v),
						zap.String("expired", t.String()),
					)
				}
			}
			return nil
		})

	}
	return
}

//
// Delete
// @Description: 清除缓存，不可以在repo层调删除缓存
// @receiver cache
// @param ctx
// @param ids
// @return error
//
func (cache Cache) Delete(ctx context.Context, ids []int32) error {
	l := len(ids)
	var keys = make([]string, l)
	for i := 0; i < l; i++ {
		id := ids[i]
		keys[i] = fmt.Sprintf(_cacheKey, id)
	}
	err := cache.rdb.Del(ctx, keys...).Err()
	if err == nil {
		cache.log.Info(_logKey,
			zap.Strings("cache-del", keys),
		)
	}
	return err
}
