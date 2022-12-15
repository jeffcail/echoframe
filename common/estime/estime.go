package _estime

import "time"

const (
	LAYOUT     = "2006-01-02 15:04:05"
	LayoutDate = "2006-01-02"
)

// FormatTime
func FormatTime(t time.Time) string {
	return t.Format(LAYOUT)
}
