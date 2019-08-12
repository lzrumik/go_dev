package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string)([]byte,error){
	resp,err := http.Get(url)
	//resp,err := http.Get("http://www.dedecms.com/")

	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("Error : status code = %s",err.Error())
	}

	bodyReader := bufio.NewReader(resp.Body)

	// 其他编码转换utf-8
	e := determineEncodeing(bodyReader)//判断网页编码
	if e == unicode.UTF8{
		return ioutil.ReadAll(bodyReader)
	}
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	return  ioutil.ReadAll(utf8Reader)
}

/**
获取网页的编码
*/
func determineEncodeing(r *bufio.Reader) encoding.Encoding{
	bytes ,err := r.Peek(1024)//获取前1024个字节
	if err != nil {
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e , _ ,_ := charset.DetermineEncoding(bytes,"")
	return e
}
