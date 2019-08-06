package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

/**
echo 111>1.txt
echo 222>2.txt

curl -X POST http://localhost:8080/upload  -F "upload[]=@1.txt"   -F "upload[]=@2.txt" -H "Content-Type: multipart/form-data"
*/
func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		fileName := ""
		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			dst,_ := os.Getwd()
			sysType := runtime.GOOS

			if sysType == "linux" {
				// LINUX系统
				dst = dst + "/src/ch88_gin/03muti_upload/" + file.Filename
			}

			if sysType == "windows" {
				// windows系统
				dst = dst + "\\src\\ch88_gin\\03muti_upload\\" + file.Filename
			}
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				panic(err)
			}
			fileName += dst+","
		}

		fileName = strings.TrimRight(fileName,",")
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded! %s ", len(files),fileName))
	})
	router.Run(":8080")
}
