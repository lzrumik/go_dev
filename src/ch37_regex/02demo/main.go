package main

import (
	"log"
	"regexp"
	"sync"
	"time"
)

const text = `fdsa@2345.com`


/**
单机跑不过php7  开协程就秒杀了
 */
func main(){

	var wg sync.WaitGroup
	executeTime := time.Now().UnixNano()

	chanNum := 4
	total := 1000000

	for i := 0; i < chanNum; i++ {
		wg.Add(1)
		go func() {
			r, _ := regexp.Compile(`([_a-z0-9]+?)([_a-z0-9-]*)@([a-z0-9-]*)(\.[a-z]{2,})`)
			str := []byte("fdsa@2345.com")
			for j := 0; j < total/chanNum; j++ {
				r.Match(str)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("execute time : ", (time.Now().UnixNano()-executeTime)/1e6, " MS")


	//re := regexp.MustCompile(`([_a-z0-9]+?)([_a-z0-9-]*)@([a-z0-9-]*)(\.[a-z]{2,})`)
	//
	//var testS []string
	//for i:=0;i<1000000;i++{
	//	testS = re.FindStringSubmatch(text)
	//}
	//
	//fmt.Println(testS)

}