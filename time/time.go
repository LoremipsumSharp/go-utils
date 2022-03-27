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

// BeginOfYear begin of year
func BeginOfYear() time.Time {
	now := time.Now()
	return time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// BeginOfMonth begin of month
func BeginOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfWeek begin of week (week begin from Sunday)
func BeginOfWeek() time.Time {
	now := time.Now()
	y, m, d := now.AddDate(0, 0, 0-int(now.Weekday())).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

// EndOfWeek end of week (week begin from Sunday)
func EndOfWeek() time.Time {
	return BeginOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// BeginOfDay begin of day
func BeginOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return BeginOfDay().AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// BeginOfHour begin of hour
func BeginOfHour() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return BeginOfHour().Add(time.Hour).Add(-time.Nanosecond)
}

// BeginOfMinute begin of minute
func BeginOfMinute() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return BeginOfMinute().Add(time.Minute).Add(-time.Nanosecond)
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
