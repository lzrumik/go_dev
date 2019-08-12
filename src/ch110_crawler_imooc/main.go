package main

import (
	"ch110_crawler_imooc/engine"
	"ch110_crawler_imooc/scheduler"
	"ch110_crawler_imooc/zhenai/parser"
)

func main() {
	//simpleEngine :=engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:50,//50个协程
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
