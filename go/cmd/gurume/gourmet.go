package gurume

import (
	"gurume/go/pkg"

	"github.com/gin-gonic/gin"
)

func FetchGourmet(c *gin.Context, keyword string) pkg.Results {
	// リクルートWEBサービスから情報を取得して返す
	data, err := pkg.Gourmet(c, keyword)
	if err != nil {
		pkg.ErrorLogFile(err)
	}

	// エラー発生時にログは残すが空の構造体を返す
	return data
}
