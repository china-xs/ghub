// Package biz
// @author: xs
// @date: 2022/6/6
// @Description: biz 文件夹区分业务方便，方法名简单 不用考虑函数名重复问题
// 控制层的主要作用就是协调model层和view层直接的调用和转换。
// 能够有效的避免请求直接进行数据库内容调用，而忽略了逻辑处理的部分。
// 实际上biz就起到了一个server服务的角色，很好的沟通了上层和下层直接的转换，
// 避免在model层进行业务处理（代码太混乱，不利于维护）
package biz

import (
	"ghub/internal/biz/auth"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	auth.NewBiz,
)
