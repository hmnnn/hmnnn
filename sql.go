package main

import (
	"fmt"

	 "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root: SixqgEPqh6.a@tcp(IP:3306)/test01?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {

	//查询
	err := Db.Select( "select * from clothes where clothing size = ? and clothing price<? ","s",100)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	err := Db.Select( "select * from storage where storage volume = (select max(storage volume))")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	err := Db.Select( "select storage volume from storage where storage code="A"")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	err := Db.Select( "select * from clothes where left(clothing code, 1)='A'")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	err := Db.Select( "select supplier infomation  from supply where clothes quality degree="unqualified"")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}



	fmt.Println("select succ:", person)
	//修改
	_, err := Db.Exec("update clothes set clothing price=clothing price*1.1 where clothing size="s"")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	//删除
	_, err := Db.Exec("delete from supply where clothes quality degree=?", "unqualified")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete succ")
		//插入
		r, err := Db.Exec("insert into clothes(clothing code,clothing size,clothing price,clothing type)values(?, ?, ?,?)", "B02", "l", 99,"famale")
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		r, err := Db.Exec("insert into storage(storage code,storage volume)values(?, ?)", "D",20)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		r, err := Db.Exec("insert into supplier(supplier code,supplier name)values(?, ?)", 5,"Luis")
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		r, err := Db.Exec("insert into supply(clothes code,supplier infomation,clothes quality degree)values(?, ?,?)", "B02",5,"qualified")
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}

}
