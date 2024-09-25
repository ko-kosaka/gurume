package pkg

import (
	"encoding/xml"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// グルメサーチAPI XMLに合わせた構造体
// ( 参考：https://webservice.recruit.co.jp/doc/hotpepper/reference.html#a1to )
type Results struct {
	XMLName          xml.Name `xml:"results"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	ApiVersion       string   `xml:"api_version"`
	ResultsAvailable string   `xml:"results_available"`
	ResultsReturned  string   `xml:"results_returned"`
	ResultsStart     string   `xml:"results_start"`
	Shop             []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id"`
		Name             string `xml:"name"`
		LogoImage        string `xml:"logo_image"`
		NameKana         string `xml:"name_kana"`
		Address          string `xml:"address"`
		StationName      string `xml:"station_name"`
		KtaiCoupon       string `xml:"ktai_coupon"`
		LargeServiceArea struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"large_service_area"`
		ServiceArea struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"service_area"`
		LargeArea struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"large_area"`
		MiddleArea struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"middle_area"`
		SmallArea struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"small_area"`
		Lat   string `xml:"lat"`
		Lng   string `xml:"lng"`
		Genre struct {
			Text  string `xml:",chardata"`
			Code  string `xml:"code"`
			Name  string `xml:"name"`
			Catch string `xml:"catch"`
		} `xml:"genre"`
		SubGenre struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"sub_genre"`
		Budget struct {
			Text    string `xml:",chardata"`
			Code    string `xml:"code"`
			Name    string `xml:"name"`
			Average string `xml:"average"`
		} `xml:"budget"`
		BudgetMemo   string `xml:"budget_memo"`
		Catch        string `xml:"catch"`
		Capacity     string `xml:"capacity"`
		Access       string `xml:"access"`
		MobileAccess string `xml:"mobile_access"`
		Urls         struct {
			Text string `xml:",chardata"`
			Pc   string `xml:"pc"`
		} `xml:"urls"`
		Open           string `xml:"open"`
		Close          string `xml:"close"`
		PartyCapacity  string `xml:"party_capacity"`
		Wifi           string `xml:"wifi"`
		OtherMemo      string `xml:"other_memo"`
		ShopDetailMemo string `xml:"shop_detail_memo"`
		Wedding        string `xml:"wedding"`
		FreeDrink      string `xml:"free_drink"`
		FreeFood       string `xml:"free_food"`
		PrivateRoom    string `xml:"private_room"`
		Horigotatsu    string `xml:"horigotatsu"`
		Tatami         string `xml:"tatami"`
		Card           string `xml:"card"`
		NonSmoking     string `xml:"non_smoking"`
		Charter        string `xml:"charter"`
		Parking        string `xml:"parking"`
		BarrierFree    string `xml:"barrier_free"`
		Show           string `xml:"show"`
		Karaoke        string `xml:"karaoke"`
		Band           string `xml:"band"`
		Tv             string `xml:"tv"`
		English        string `xml:"english"`
		Pet            string `xml:"pet"`
		Child          string `xml:"child"`
		CouponUrls     struct {
			Text string `xml:",chardata"`
			Pc   string `xml:"pc"`
			Sp   string `xml:"sp"`
		} `xml:"coupon_urls"`
		Course string `xml:"course"`
		Photo  struct {
			Text string `xml:",chardata"`
			Pc   struct {
				Text string `xml:",chardata"`
				L    string `xml:"l"`
				M    string `xml:"m"`
				S    string `xml:"s"`
			} `xml:"pc"`
			Mobile struct {
				Text string `xml:",chardata"`
				L    string `xml:"l"`
				S    string `xml:"s"`
			} `xml:"mobile"`
		} `xml:"photo"`
		Lunch    string `xml:"lunch"`
		Midnight string `xml:"midnight"`
	} `xml:"shop"`
}

// グルメサーチAPIから情報を取得する
func Gourmet(c *gin.Context, keyword string) (Results, error) {

	// ホットペッパー グルメサーチAPIのURL
	apiURL := "http://webservice.recruit.co.jp/hotpepper/gourmet/v1/"
	apiKey := os.Getenv("RECRUIT_WEB_SERVICE_KEY")
	shops := Results{}

	// GETリクエストを送信するためのURLを定義する
	url := apiURL + "?key=" + apiKey + "&keyword=" + keyword
	// GETリクエストを送信するためのHTTPクライアントを生成する
	client := http.Client{}

	// GETリクエストを送信する
	response, err := client.Get(url)
	if err != nil {
		ErrorLogFile(err)
		return Results{}, err
	}
	defer response.Body.Close()

	// XMLレスポンスをパースして構造体に変換する
	err = xml.NewDecoder(response.Body).Decode(&shops)
	if err != nil {
		ErrorLogFile(err)
		return Results{}, err
	}

	// 正常終了時にログを残す(logを吐き出すディレクトリのパスは「main.go」から見たパス)
	ApiLogFile(GourmetFileName, c)

	return shops, nil
}
