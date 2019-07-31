package json

import (
	"testing"

	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

const (
	jstr = `{"user_id":"A5138C84-9C18-4653-A657-52BFDC04377C","invest_repay_tm":"2016-04-19 13:29:29.0","invest_repay_amount":52.61,"loan_id":"F2B652F7-006D-41A0-B569-3E09E7323293","invest_repay_status":"REPAYED","invest_repay_due_date":"2016-04-19","plan_repay_amount":52.61,"ca_cd":2,"uid":12175268}`
)

//
//func Benchmark_array_by_json(b *testing.B) {
//	b.ReportAllocs()
//	for n := 0; n < b.N; n++ {
//		sample := make([]int, 0, 10)
//		json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
//	}
//}
//
//func Benchmark_array_by_jsoniter(b *testing.B) {
//	b.ReportAllocs()
//	for n := 0; n < b.N; n++ {
//		sample := make([]int, 0, 10)
//		jsoniter.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
//	}
//}

func Benchmark_decode_by_json(b *testing.B) {
	b.ReportAllocs()
	var jsonstr RepayInfo
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(jstr), &jsonstr)
	}
}

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