package parser

import (
	"ch110_crawler_imooc/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	fmt.Printf("%s\n", contents)

	ParseCityList(contents)
}
