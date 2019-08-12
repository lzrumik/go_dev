package engine

import (
	"ch110_crawler_imooc/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)//将请求url 加入队列等待执行
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]  //循环处理请求
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, v := range parseResult.Items {
			log.Printf("Get Item %v", v)
		}
	}
}

//请求单个url 返回处理后的结果
/**
地址列表的结果 返回各个城市的地址
 */
func worker(r Request) (ParseResult, error) {

	log.Printf("Fecthing %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error at %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
