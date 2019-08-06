package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime"
)

/**
echo 111>1.txt

curl -X POST http://localhost:8080/upload  -F "file=@1.txt"  -H "Content-Type: multipart/form-data"
 */
func main() {
	fmt.Println(os.Getwd())

	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		dst,_ := os.Getwd()
		sysType := runtime.GOOS

		if sysType == "linux" {
			// LINUX系统
			dst = dst + "/src/ch88_gin/02upload/" + file.Filename
		}

		if sysType == "windows" {
			// windows系统
			dst = dst + "\\src\\ch88_gin\\02upload\\" + file.Filename
		}
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			panic(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!  %s ", file.Filename ,dst))
	})
	router.Run(":8080")
}
