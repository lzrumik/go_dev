package main

import (
	"ch23_error/10unify_handling_errors/filelisting"
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,request *http.Request)error

/**
输入一个函数  输出一个函数
输入的函数统一返回err 输出的函数返回
 */
func errWrapper( handler appHandler)/*前面参数 后面返回值 */ func(w http.ResponseWriter,r *http.Request){

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer,request)
		if err!= nil {
			logs.Warn("Error handling request : %s",err.Error())
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}


//noinspection ALL
func main (){
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}


}