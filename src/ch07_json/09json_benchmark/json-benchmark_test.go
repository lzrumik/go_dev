package main

import (
	"testing"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

const (
	jstr = `{"user_id":"A5138C84-9C18-4653-A657-52BFDC04377C","invest_repay_tm":"2016-04-19 13:29:29.0","invest_repay_amount":52.61,"loan_id":"F2B652F7-006D-41A0-B569-3E09E7323293","invest_repay_status":"REPAYED","invest_repay_due_date":"2016-04-19","plan_repay_amount":52.61,"ca_cd":2,"uid":12175268}`
)

/**
go test -bench=. json-benchmark_test.go struct.go struct_easyjson.go
 */
//json go原生解析
func Benchmark_decode_by_json(b *testing.B) {
	b.ReportAllocs()
	var jsonstr RepayInfo
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(jstr), &jsonstr)
	}
}

//jsoniter  直接可以解析
func Benchmark_decode_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	var jsonstr RepayInfo
	for n := 0; n < b.N; n++ {
		jsoniter.Unmarshal([]byte(jstr), &jsonstr)
	}
}


//func Benchmark_decode_by_msgpack(b *testing.B) {
//	b.ReportAllocs()
//	//var jsonstr RepayInfo
//	for n := 0; n < b.N; n++ {
//		v, _, e := msgpack.Unpack([]byte(jstr))
//	}
//}

func Benchmark_decode_by_easyjson(b *testing.B) {
	b.ReportAllocs()
	var jsonstr RepayInfo
	for n := 0; n < b.N; n++ {
		jsonstr.UnmarshalJSON([]byte(jstr))
	}
}