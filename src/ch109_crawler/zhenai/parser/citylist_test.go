package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",contents)

	result := ParseCityList(contents)
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng"}
	expectedCitys := []string{"阿坝","阿克苏","阿拉善盟"}

	const resultSize = 470
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d ",resultSize,len(result.Items))
	}

	for i,url := range expectedUrls{
		if result.Request[i].Url != url {
			t.Errorf("expected url #%d:%s;but was %s ",i,url,result.Request[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d ",resultSize,len(result.Items))
	}

	for i,city := range expectedCitys{
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d:%s;but was %s ",i,city,result.Items[i].(string))
		}
	}
}
