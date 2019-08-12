package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 10 毫秒 等待
var rateLimit = time.Tick(10 *time.Millisecond)
func Fetch(url string) ([]byte, error) {
	// 10 毫秒 等待
	//<-rateLimit
	client := &http.Client{}

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "xxx")
	reqest.Header.Add("X-Requested-With", "xxxx")

	if err != nil {
		return nil, err
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("http fail")
	}
	//utf8Reader :=transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//按照编码解码
	bufioReader := bufio.NewReader(response.Body)
	encode := determineEncoding(bufioReader)
	utf8Reader := transform.NewReader(bufioReader, encode.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)

	//bytes, err :=ioutil.ReadAll(r) //直接读取reader后就不能再读取，所以用bufferReader
	if err != nil {
		log.Printf("Fetcher error :%v", err)
		return unicode.UTF8 //默认utf8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "text/html")
	return e
}
