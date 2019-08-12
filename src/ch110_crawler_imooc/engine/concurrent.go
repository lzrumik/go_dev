package engine

import (
	"log"
)

//并行引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler  //调度器
	WorkerCount int  //work数量
}

//调度器接口
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChannel(chan Request)
	WorkerReady(chan Request)
	Run()
}

//并行爬虫 执行
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//创建结果管道
	out := make(chan ParseResult)

	//从main中配置而来
	e.Scheduler.Run()

	//创建50个爬虫协程
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	//将连接  加入调度器
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount :=1
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Item %d :%v",itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

//创建抓取 和 处理结果 的协程
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			//地址列表的结果 返回各个城市的地址
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
