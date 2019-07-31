package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

type Place struct {
	Country int `db:"country"`
	City int `db:"city"`
	TelCode int `db:"telcode"`
}

var DB *sqlx.DB

func init(){
	database,err := sqlx.Open("mysql","root:123456@tcp(127.0.0.1:3307)/gomysql")
	if err!=nil{
		fmt.Println("exec failed," ,err)
		return
	}
	DB = database
}

func main(){

	r,err := DB.Exec("insert into person(username,sex,email) values (?,?,?)","stu01","man","stu01@qq.com")
	if err!=nil{
		fmt.Println("insert failed," ,err)
		return
	}

	id,err := r.LastInsertId()
	if err!=nil{
		fmt.Println("get lastid  failed," ,err)
		return
	}

	fmt.Println(id)
}