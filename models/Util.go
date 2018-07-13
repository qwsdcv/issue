package models

import (
	"time"
)

// FormatTime formats *time.Time to string with format "2006-01-02 15:04:05"
func FormatTime(t *time.Time) (ret string) {
	return t.Format("2006-01-02 15:04:05")
}
