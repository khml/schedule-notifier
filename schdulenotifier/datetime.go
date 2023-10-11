package schdulenotifier

import (
	"time"
)

func ParseDate(dateString string, location *time.Location) (time.Time, error) {
	// dateString := "202310110930" // "YYYYMMDDhhmm"形式の日付文字列
	layout := "200601021504" // フォーマットに対応したレイアウト文字列

	// 文字列をtime.Time型にパース
	return time.ParseInLocation(layout, dateString, location)
}
