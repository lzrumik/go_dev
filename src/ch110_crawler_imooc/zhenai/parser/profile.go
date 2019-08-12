package parser

import (
	"ch110_crawler_imooc/engine"
	"ch110_crawler_imooc/model"
	"encoding/json"
	"fmt"
	"regexp"
)

var compiledProfileReg = regexp.MustCompile(`window.__INITIAL_STATE__=(.*);\(function\(\){var s;`)
var compiledIntRed = regexp.MustCompile(`(\d+).*`)

//用户详情解析器
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	match := compiledProfileReg.FindSubmatch(contents)
	if len(match) <= 1 {
		return engine.ParseResult{
			Items: nil,
		}
	}
	sourceProfile := model.ZhenaiProfile{}
	err := json.Unmarshal(match[1], &sourceProfile)
	if err != nil {
		panic(err)
	}

	profile.Age = sourceProfile.ObjectInfo.Age
	profile.Gender = sourceProfile.ObjectInfo.GenderString
	profile.Education = sourceProfile.ObjectInfo.EducationString
	profile.Marriage = sourceProfile.ObjectInfo.MarriageString
	profile.Height = sourceProfile.ObjectInfo.HeightString
	profile.Income = sourceProfile.ObjectInfo.ObjectSalaryString
	profile.Residence = sourceProfile.ObjectInfo.ObjectWorkCityString

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	fmt.Println(profile)
	return result
}
