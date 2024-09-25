package pkg

import "time"

const (
	DateFormat = "2006-01-02" // YYYY-MM-DD
)

// 日本時間(JST)の現在時刻を取得
func NowJST() time.Time {
	now := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	return now.In(jst)
}

// 任意の形にフォーマットされた日付を取得
func FormatDateTime(dateTime time.Time, format string) string {
	return dateTime.Format(format)
}
