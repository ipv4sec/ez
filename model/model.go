package model

import (
	"reflect"
	"errors"
	"fmt"
)

// ---- ---- ---- ----

type TableNamer interface {
	TableName() string
}

// ---- ---- ---- ----

var model *Model
var GetInstance = func() (*Model) {
	if model == nil {
		model = &Model{}
	}
	return model
}

// ---- ---- ---- ---

var SetupEzORMContext = func(v interface{}) {
	t := reflect.TypeOf(v).Elem()

	cols := map[string]*Column{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		cols[f.Name] = &Column{
			f.Name,f.Type,0,f.Name,
		}
		// fmt.Println(f.Name, f.Type, f.Tag.Get("ez"))
	}
	GetInstance().Cols = cols
	GetInstance().parseTableName(v)

	fmt.Println(GetInstance())
}


// ----- ---- ---- ----
type Model struct {
	Name          string                 // 表的名称
	Cols          map[string]*Column     // 所有的列
	//UniqueIndexes map[string][]*Column   // 唯一索引列
	//PK            []*Column              // 主键
}

func (m *Model) parseTableName(v interface{}) error {
	tableName, ok := v.(TableNamer)
	if !ok {
		return errors.New(`TableName Not Set`)
	}
	GetInstance().Name = tableName.TableName()
	return nil
}


