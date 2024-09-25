package main

import (
	"gurume/go/cmd/gurume"
	"gurume/go/pkg"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envLoad()

	// サーバーの起動
	r := gurume.SetRouter()
	err := r.Run()
	if err != nil {
		pkg.ErrorLogFile(err)
		log.Fatal("サーバーの起動に失敗しました", err.Error())
	}
}

// 環境変数の読み込み
func envLoad() {
	err := godotenv.Load()
	if err != nil {
		pkg.ErrorLogFile(err)
		log.Fatal(".envファイルを読み込めませんでした", err)
	}
}
