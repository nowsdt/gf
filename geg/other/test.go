package main

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/util/gvalid"
)

type User struct {
	Uid   int    `v:"uid      @integer|min:1"`
	Name  string `v:"name     @required|length:6,30#请输入用户名称|用户名称长度非法"`
	Pass1 string `v:"password1@required|password3"`
	Pass2 string `v:"password2@required|password3|same:password1#||两次密码不一致，请重新输入"`
}

func main() {
	user := &User{
		Name:  "john",
		Pass1: "Abc123!@#",
		Pass2: "123",
	}

	// 使用结构体定义的校验规则和错误提示进行校验
	if e := gvalid.CheckStruct(user, nil); e != nil {
		g.Dump(e.Maps())
	}

	// 自定义校验规则和错误提示，对定义的特定校验规则和错误提示进行覆盖
	rules := map[string]string{
		"uid": "min:6",
	}
	msgs := map[string]interface{}{
		"password2": map[string]string{
			"password3": "名称不能为空",
		},
	}
	if e := gvalid.CheckStruct(user, rules, msgs); e != nil {
		g.Dump(e.Maps())
	}
}
