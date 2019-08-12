package main

import (
	"ch109_crawler/engine"
	"ch109_crawler/zhenai/parser"
)

func main(){
	engine.Run(
		engine.Request{
			Url:"http://www.zhenai.com/zhenghun",
			ParseFunc:parser.ParseCityList,
		})
}