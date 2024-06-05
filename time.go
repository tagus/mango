package mango

import (
	"fmt"
	"time"
)

const (
	DurationDay  = 24 * time.Hour
	DurationYear = 365 * DurationDay
)

func FormatSimpleDate(ts *time.Time) string {
	if ts == nil {
		return "never"
	}
	return ts.Format("01-02-06")
}

func FormatTimeSince(ts *time.Time) string {
	if ts == nil {
		return "never"
	}
	now := time.Now()
	elapsed := now.Sub(*ts)
	return fmt.Sprintf("%v ago", formatDuration(elapsed))
}

// imprecise formatting since we only care about the highest possible unit
func formatDuration(d time.Duration) string {
	if d >= DurationYear {
		years := d / DurationYear
		return fmt.Sprintf("%dy", years)
	} else if d >= DurationDay {
		days := d / DurationDay
		return fmt.Sprintf("%dd", days)
	} else if d >= time.Hour {
		hours := d / time.Hour
		return fmt.Sprintf("%dh", hours)
	} else if d >= time.Minute {
		minutes := d / time.Minute
		return fmt.Sprintf("%dm", minutes)
	}
	seconds := d / time.Second
	return fmt.Sprintf("%ds", seconds)
}
