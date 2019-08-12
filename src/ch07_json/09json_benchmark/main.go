package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

const (
	jstr_new = `{"user_id":"A5138C84-9C18-4653-A657-52BFDC04377C","invest_repay_tm":"2016-04-19 13:29:29.0","invest_repay_amount":52.61,"loan_id":"F2B652F7-006D-41A0-B569-3E09E7323293","invest_repay_status":"REPAYED","invest_repay_due_date":"2016-04-19","plan_repay_amount":52.61,"ca_cd":2,"uid":12175268}`
)



/**
go run main.go struct.go struct_easyjson.go
 */
func main(){

	var jsonstr RepayInfo

	json.Unmarshal([]byte(jstr_new), &jsonstr)

	fmt.Println(jsonstr)



	var jsonstr2 RepayInfo

	jsoniter.Unmarshal([]byte(jstr_new), &jsonstr2)

	fmt.Println(jsonstr2)


	/**

	easyjson 用法

	cd /Users/kevint/github/go_dev/src/github.com/mailru/easyjson/easyjson/
	go build .
	cd  ~/github/go_dev/src/ch07_json/09json_benchmark/
	/Users/kevint/github/go_dev/src/github.com/mailru/easyjson/easyjson/easyjson -all struct.go

	 */

	var jsonstr3 RepayInfo
	jsonstr3.UnmarshalJSON([]byte(jstr_new))
	fmt.Println(jsonstr3)
}
