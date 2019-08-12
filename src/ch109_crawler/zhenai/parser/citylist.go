package parser

import (
	"ch109_crawler/engine"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)".+?>(.+?)</a>`
func ParseCityList(content []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListReg)

	match := re.FindAllStringSubmatch(string(content),-1)

	result := engine.ParseResult{}
	for _,m := range match{
		result.Items = append(result.Items,string(m[2]))
		result.Request = append(result.Request,
			engine.Request{Url:string(m[1]),ParseFunc:engine.NilParser})
	}
	return result
}
