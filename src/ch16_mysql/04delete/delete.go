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
	database,err := sqlx.Open("mysql","root:123456@tcp(127.0.0.1:3306)/gomysql")
	if err!=nil{
		fmt.Println("exec failed," ,err)
		return
	}
	DB = database
}

func main(){

	_,err := DB.Exec("delete from  person  where user_id = ?",1)
	if err!=nil{
		fmt.Println("delete failed," ,err)
		return
	}

	var person []Person
	err = DB.Select(&person,"select user_id,username,sex,email from person where user_id = ?",1)
	if err!=nil{
		fmt.Println("select failed," ,err)
		return
	}

	fmt.Println(person)
}