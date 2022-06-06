// Package tmp
// @author: xs
// @date: 2022/5/18
// @Description: tpl repo|cache 根据标名称生成 repo cache
package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//go:embed stubs
var stubsFS embed.FS

func main() {
	//var filePath string
	var models []string
	// 待完成
	models = []string{"account", "role"}
	l := len(models)
	modelData, err := stubsFS.ReadFile("stubs/repo.stub")
	if err != nil {
		panic("stub 文件不存在")
	}
	for i := 0; i < l; i++ {
		var filePath, dir string
		t := models[i]
		dir = fmt.Sprintf("./internal/data/%s", t)
		os.MkdirAll(dir, os.ModePerm)
		filePath = dir + "/" + t + ".gen.go"
		modelStub := string(modelData)
		replaces := make(map[string]string)
		// 添加默认的替换变量
		replaces["{{PackageName}}"] = t
		replaces["{{ModelName}}"] = strings.ToUpper(string(t[0])) + t[1:]

		// 对模板内容做变量替换
		for search, replace := range replaces {
			modelStub = strings.ReplaceAll(modelStub, search, replace)
		}
		//fmt.Printf("%v\n", modelStub)
		if err := ioutil.WriteFile(filePath, []byte(modelStub), 0644); err != nil {
			panic(err)
		}

	}
	/*
		filePath = "./internal/data/role/role.go"
		//filePath ="role.go"
		// 实现最后一个参数可选

		replaces := make(map[string]string)
		// 读取 stub 模板文件
		modelData, err := stubsFS.ReadFile("stubs/repo.stub")
		if err != nil {
			panic("stub 文件不存在")
		}
		modelStub := string(modelData)

		// 添加默认的替换变量
		replaces["{{PackageName}}"] = "role"
		replaces["{{ModelName}}"] = "Role"

		// 对模板内容做变量替换
		for search, replace := range replaces {
			modelStub = strings.ReplaceAll(modelStub, search, replace)
		}
		fmt.Printf("%v\n", modelStub)
		if err := ioutil.WriteFile(filePath, []byte(modelStub), 0644); err != nil {
			panic(err)
		}

	*/
}
