package parser

import (
	"ch110_crawler_imooc/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>(\p{Han}+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {

	match := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//limit := 10
	for _, v := range match {
		result.Items = append(result.Items, "City "+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(v[1]),
				ParserFunc: ParseCity,
			},
		)
		//if limit--; limit == 0 {
		//	break
		//}
	}
	return result
}
