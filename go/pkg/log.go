package pkg

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ApiLogPath      = "./go/log/api/"
	ErrorLogPath    = "./go/log/error/"
	GourmetFileName = "gourmet-"
	LogExtension    = ".log"
)

// API実行時のログファイルを出力
func ApiLogFile(filePath string, c *gin.Context) {
	// 指定したディレクトリにログファイルを出力(日付でログファイルを分ける)
	now := NowJST()
	logDate := FormatDateTime(now, DateFormat)
	logFile, err := os.OpenFile(ApiLogPath+GourmetFileName+logDate+LogExtension, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("APIログファイルを作成時にエラーが発生しました" + err.Error())
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(c.ClientIP())
}

// エラー発生時のログファイルを生成
func ErrorLogFile(errContent error) {
	// 指定したディレクトリにログファイルを出力(日付でログファイルを分ける)
	now := NowJST()
	logDate := FormatDateTime(now, DateFormat)
	logFile, err := os.OpenFile(ErrorLogPath+logDate+LogExtension, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("エラーログファイルを作成時にエラーが発生しました" + err.Error())
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(errContent)
}
