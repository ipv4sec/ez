package model

import "reflect"

type Column struct {
	//model *Model

	Name     string       // 数据库的字段名
	//Nullable bool         // 是否可以为 NULL

	GoType   reflect.Type // Go 语言中的数据类型
	Zero     interface{}  // GoType 的零值
	GoName   string       // 在Go程序中的字段名

	//HasDefault bool
	//Default    string // 默认值
}
