package datetime

import (
	"time"
)

// FormatDate は time.Time を "YYYY-MM-DD" 形式の文字列に変換
func FormatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}
