package time

import (
	"fmt"
	"time"
)

const (
	DateFormat       = "2006-01-02"
	DatetimeFormat   = "2006-01-02 15:04:05"
	DatetimeMsFormat = "2006-01-02 15:04:05.999"
	DatetimeNsFormat = "2006-01-02 15:04:05.999999999"
)

var (
	UTCLoc     = time.UTC                        // utc location
	BeijingLoc = time.FixedZone("UTC+8", 8*3600) // beijing location
)

// NowUnix now timestamp (second)
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixMs now timestamp (millisecond)
func NowUnixMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// NowUnixNs now timestamp (nanosecond)
func NowUnixNs() int64 {
	return time.Now().UnixNano()
}

// NowDateStr now date string
func NowDateStr() string {
	return time.Now().Format(DateFormat)
}

// NowDatetimeStr now datetime string
func NowDatetimeStr() string {
	return time.Now().Format(DatetimeFormat)
}

// NowDatetimeMsStr now datetime string (millisecond)
func NowDatetimeMsStr() string {
	return time.Now().Format(DatetimeMsFormat)
}

// NowDatetimeNsStr now datetime string (nanosecond)
func NowDatetimeNsStr() string {
	return time.Now().Format(DatetimeNsFormat)
}

// BeginOfMinute return beginning minute time of day.
// Play: https://go.dev/play/p/ieOLVJ9CiFT
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute return end minute time of day.
// Play: https://go.dev/play/p/yrL5wGzPj4z
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour return beginning hour time of day.
// Play: https://go.dev/play/p/GhdGFnDWpYs
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour return end hour time of day.
// Play: https://go.dev/play/p/6ce3j_6cVqN
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay return beginning hour time of day.
// Play: https://go.dev/play/p/94m_UT6cWs9
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay return end time of day.
// Play: https://go.dev/play/p/eMBOvmq5Ih1
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek return beginning week, default week begin from Sunday.
// Play: https://go.dev/play/p/ynjoJPz7VNV
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time {
	var beginFromWeekday = time.Sunday
	if len(beginFrom) > 0 {
		beginFromWeekday = beginFrom[0]
	}
	y, m, d := t.AddDate(0, 0, int(beginFromWeekday-t.Weekday())).Date()
	beginOfWeek := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	if beginOfWeek.After(t) {
		return beginOfWeek.AddDate(0, 0, -7)
	}
	return beginOfWeek
}

// EndOfWeek return end week time, default week end with Saturday.
// Play: https://go.dev/play/p/i08qKXD9flf
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time {
	var endWithWeekday = time.Saturday
	if len(endWith) > 0 {
		endWithWeekday = endWith[0]
	}
	y, m, d := t.AddDate(0, 0, int(endWithWeekday-t.Weekday())).Date()
	var endWithWeek = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	if endWithWeek.Before(t) {
		endWithWeek = endWithWeek.AddDate(0, 0, 7)
	}
	return endWithWeek
}

// BeginOfMonth return beginning of month.
// Play: https://go.dev/play/p/bWXVFsmmzwL
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth return end of month.
// Play: https://go.dev/play/p/_GWh10B3Nqi
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear return the date time at the begin of year.
// Play: https://go.dev/play/p/i326DSwLnV8
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear return the date time at the end of year.
// Play: https://go.dev/play/p/G01cKlMCvNm
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// UnixToTime convert timestamp (second) to time.Time
func UnixToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// UnixMsToTime convert timestamp (millisecond) to time.Time
func UnixMsToTime(msec int64) time.Time {
	return time.Unix(0, msec*1e6)
}

// UnixNsToTime convert timestamp (nanosecond) to time.Time
func UnixNsToTime(nsec int64) time.Time {
	return time.Unix(0, nsec)
}

// TimeToUnix convert time.Time to timestamp (second)
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

// TimeToUnixMs convert time.Time to timestamp (millisecond)
func TimeToUnixMs(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// TimeToUnixNs convert time.Time to timestamp (nanosecond)
func TimeToUnixNs(t time.Time) int64 {
	return t.UnixNano()
}

// TimeToDateStr convert time.Time to date string
func TimeToDateStr(t time.Time) string {
	return t.Format(DateFormat)
}

// DateStrToTime convert date string to time.Time
func DateStrToTime(val string) (time.Time, error) {
	return time.Parse(DateFormat, val)
}

// DatetimeStrToTime convert datetime string (millisecond) to time.Time
func DatetimeStrToTime(val string) (time.Time, error) {
	return time.Parse(DatetimeFormat, val)
}

// DatetimeMsStrToTime convert datetime string (nanosecond) to time.Time
func DatetimeMsStrToTime(val string) (time.Time, error) {
	return time.Parse(DatetimeMsFormat, val)
}

// DatetimeNsStrToTime convert datetime string to time.Time
func DatetimeNsStrToTime(val string) (time.Time, error) {
	return time.Parse(DatetimeNsFormat, val)
}

// TimeToDatetimeStr convert time.Time to datetime string
func TimeToDatetimeStr(t time.Time) string {
	return t.Format(DatetimeFormat)
}

// TimeToDatetimeMsStr convert time.Time to datetime string (millisecond)
func TimeToDatetimeMsStr(t time.Time) string {
	return t.Format(DatetimeMsFormat)
}

// TimeToDatetimeNsStr convert time.Time to datetime string (nanosecond)
func TimeToDatetimeNsStr(t time.Time) string {
	return t.Format(DatetimeNsFormat)
}

// print func exec time cost
func PrintTimeCost(prefix ...string) func() {
	start := time.Now()
	prefixStr := ""
	if len(prefix) > 0 {
		prefixStr = prefix[0]
	}
	return func() {
		fmt.Printf("%s time cost: %v\n", prefixStr, time.Since(start))
	}
}
