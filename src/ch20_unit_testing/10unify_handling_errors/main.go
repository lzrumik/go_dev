package main

import (
	"ch23_error/10unify_handling_errors/filelisting"
	"log"
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
		defer func(){
			if r := recover();r!=nil {
				log.Printf("Panic: %v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
				return
			}
		}()

		err := handler(writer,request)
		if err!= nil {
			log.Printf("Panic: %v",err)
			if usererr,ok := err.(userError);ok{
				http.Error(writer,usererr.Message(),http.StatusBadRequest)
				return
			}

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

/*
给用户看的error
 */
type userError interface {
	error
	Message() string
}

//noinspection ALL
func main (){
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}


}