package dao

import (
	"database/sql"
	_ "debug/elf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
//数据库配置
//const (
//	username ="root"
//	password =""
//	ip ="127.0.0.1"
//	port ="8080"
//	Dbname ="loginserver"
//)

var Db *sql.DB
//连接数据库
func Initdb()(err error)  {
	dsn:="root:Dyy120612120@/loginserver"
    Db,err=sql.Open("mysql",dsn)
	if err!=nil{
		log.Fatal(err)
	}
    err=Db.Ping()
	if err!=nil {
		fmt.Println("fail")
		return
	}
	fmt.Println("连接成功")
	return
}
//对数据库进行单行查询操作
func Queryone(sqlstr string,args ...interface{}) *sql.Row {//element 表示要查询的元素
	row:=Db.QueryRow(sqlstr,args...)
	//if err!=nil{
	//	fmt.Printf("err:%v/n",err)
	//	return false
	//}
	return row
}
//对数据库进行查询返回多行结果的操作
func QueryMore(sqlstr string,args ...interface{})(*sql.Rows,error){
	 rows,err:=Db.Query(sqlstr,args...)
	 //defer rows.Close()
	 if err!=nil{
		 fmt.Println(err)
		 return nil,err
	 }else{
		 return rows,nil
	 }

}
//对数据库进行插入数据的操作
func Insertdb(sqlstr string,args ...interface{})int64{
	ret,err:=Db.Exec(sqlstr,args...)
	if err!=nil{
		fmt.Printf("insert faild,err:%v\n",err)
		return 0
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		return 0
	}
	return n
}
//对数据库进行修改数据的操作
func Updatedb(sqlstr string ,args ...interface{})int64{
	ret,err:=Db.Exec(sqlstr,args...)
	if err!=nil{
		fmt.Println(err)
		return 0
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err!=nil{
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return 0
	}
	return n
}
//对数据库进行删除数据的操作
func Deletedb(sqlstr string, args ...interface{})int64{
	ret,err:=Db.Exec(sqlstr,args...)
	if err!=nil{
		fmt.Println(err)
		return 0
	}
	n,err :=ret.RowsAffected()
	if err!=nil{
		fmt.Println(err)
		return 0
	}
	return n
}


