package mango

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
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

/******************************************************************************/

var validTimestampFormats = []string{
	time.RFC3339,  // "2023-10-01T12:00:00Z"
	time.DateOnly, // "2023-10-01"
	time.DateTime, // "2023-10-01 12:00:00"
}

// ParseTimestamp attempts to parse a string into a Timestamp
func ParseTimestamp(s string) (Timestamp, error) {
	var ts Timestamp
	for _, format := range validTimestampFormats {
		parsedTime, err := time.ParseInLocation(format, s, time.Local)
		if err == nil {
			ts.Time = parsedTime
			return ts, nil
		}
	}
	return ts, fmt.Errorf("invalid timestamp given: %s", s)
}

// Timestamp is a simple wrapper around time.Time that allows for more permissive
// JSON unmarshalling of strings. This can be used as a first class field in structs
// and is fully compatible with sql drivers.
//
// The supported timestamp formats are:
// - RFC3339 (e.g. "2023-10-01T12:00:00Z")
// - date only (e.g. "2023-10-01")
// - date time (e.g. "2023-10-01 12:00:00")
type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "" || s == "null" {
		t.Time = time.Time{}
		return nil
	}
	ts, err := ParseTimestamp(s)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}
	*t = ts
	return nil
}

func (t *Timestamp) Scan(value any) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case string:
		return t.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("unsupported type for Timestamp: %T", value)
	}
}

func (t Timestamp) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

/******************************************************************************/

// Duration is a simple wrapper around time.Duration that supports json unmarshalling
// from a duration string (e.g. "1h30m", "45s") or a raw numeric nanoseconds value.
// This can be used as a first class field in structs.
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		v, err := time.ParseDuration(s)
		if err != nil {
			return fmt.Errorf("failed to parse duration: %w", err)
		}
		*d = Duration{Duration: v}
		return nil
	}

	// fallback: numeric (nanoseconds)
	var n int64
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	*d = Duration{Duration: time.Duration(n)}
	return nil
}
