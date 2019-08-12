package parser

import (
	"ch110_crawler_imooc/fetcher"
	"fmt"
	"testing"
)

func TestCity(t *testing.T) {
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun/chongqing")

	//fmt.Printf("%s\n",contents)
	fmt.Println(ParseCity(contents))

}
