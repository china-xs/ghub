// Package gorm
// @author: xs
// @date: 2022/5/17
// @Description: 根目录 cli:make gorm
package main

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gen/field"

	//"github.com/casbin/casbin/v2"
	"github.com/china-xs/gin-tpl/pkg/config"
	db2 "github.com/china-xs/gin-tpl/pkg/db"

	"go.uber.org/zap"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../internal/data/dao/query",
		FieldNullable: true,
		Mode:          gen.WithDefaultQuery, // 初始化调用,减少没次使用表都要 query.user(db).***
	})
	l, _ := zap.NewProduction()
	// config 注意表前缀问题
	v, err := config.New("../../configs/app.yaml")
	opts, err := db2.New(v)
	db, fc, err := db2.NewDb(opts, l)

	//db.AutoMigrate(
	//	model.Account{},
	//	model.Role{},
	//	)
	// 自动生成 casbin_rule 表
	gormadapter.NewAdapterByDBUseTableName(db, "a", "casbin_rule")

	if err != nil {
		panic(fmt.Sprintf("init-db:%v", err.Error()))
	}
	defer fc()
	// 复用工程原本使用的SQL连接配置
	g.UseDB(db)
	// 所有需要实现查询方法的结构体 增加表不能把原来表删除
	// 不配置外建也可以，不过删除管理需要代码自行检查是否把关系删除完整
	user2role := g.GenerateModel("a_user2role",
		gen.FieldGORMTag("account_id", "column:account_id;primaryKey"),
		gen.FieldGORMTag("role_id", "column:role_id;primaryKey"),
	)
	role := g.GenerateModel("a_role")
	g.ApplyBasic(
		role,
		user2role,
		g.GenerateModel("a_casbin_rule"),
		g.GenerateModel("a_rule"),
		g.GenerateModel("a_account",
			gen.FieldRelate(
				field.Many2Many,
				"Roles",
				role,
				&field.RelateConfig{
					GORMTag: "many2many:user2role;",
				},
			),
			gen.FieldType("state", "int32"),
		),
		/*
			g.GenerateModel("a_role",
				gen.FieldRelate(
					field.Many2Many,
					"Accounts",
					user2role,
					&field.RelateConfig{
						GORMTag: "foreignKey:RoleRefer",
					},
				),
			),
		*/
	)

	g.Execute()

}
