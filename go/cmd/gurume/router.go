package gurume

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func SetRouter() *gin.Engine {
	router := gin.Default()
	// 静的ファイルの設定
	router.Static("/assets/css", "go/web/template/gurume/assets/css/")
	router.Static("/assets/image", "go/web/template/gurume/assets/image/")
	// HTMLテンプレート 配置
	router.LoadHTMLGlob("./go/web/template/gurume/html/*.html")

	// BASIC認証を挟みたいルーティングにはauthorizedを使用する
	authorized := router.Group(
		"/", gin.BasicAuth(
			// BASIC認証のユーザーデータ
			gin.Accounts{
				// ユーザー名: パスワード
				os.Getenv("BASIC_USERNAME"): os.Getenv("BASIC_PASSWORD"),
			},
		),
	)

	authorized.GET("/", routeIndex)
	authorized.GET("/gourmetSearch", routeGourmetSearch)

	router.NoRoute(routeNotFound)

	return router
}

// 初期画面
func routeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func routeGourmetSearch(c *gin.Context) {
	// 入力値取得(パラメータ キーワード)
	keyword := template.HTMLEscapeString(c.Query("keyword"))
	data := FetchGourmet(c, keyword)
	// APIから取得したデータを返す
	c.HTML(http.StatusOK, "gourmetSearch.html", gin.H{
		"Data": data,
	})
}

func routeNotFound(c *gin.Context) {
	// 存在しないURLへ遷移した場合は404ページへ
	c.HTML(http.StatusNotFound, "404.html", nil)
}
