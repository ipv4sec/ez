package main

import (
	"ez"
	"ez/user"
)

func main(){

	db := ez.GetClient()
	db.Create(&user.User{}) // 要求传入指针，现在程序内部没有判断

	//// Update:
	//db.Update(&User{Id:1,FirstName:"abc"})
	//sqlbuilder.Update(db).Table("#table").Where("id=?",1).Set("FirstName", "abc").Exec()
	//
	//// Delete:
	//e.Delete(&User{Id:1})
	//sqlbuilder.Delete(e).Table("#table").Where("id=?",1).Exec()
	//
	//// Insert:
	//db.Insert(&User{Id:1,FirstName:"abc"})
	//tx.InsertMany(&User{Id:1,FirstName:"abc"},&User{Id:1,FirstName:"abc"})
	//
	//// Select:
	//_,err := sqlbuilder.Select(e, e.Dialect()).Select("*").From("{#table}").Where("id=1").QueryObj(obj)
	//user := &User{Id:1}
	//err := e.Select(u)
	//
	//// Query/Exec:
	//// Query 返回参数与 sql.Query 是相同的
	//sql := "select * from #tbl_name where id=?"
	//rows, err := e.Query(sql, []interface{}{5})
	//sql = "update #tbl_name set name=? where id=?"
	//r, err := e.Exec(sql, []interface{}{"name1", 5})
}
