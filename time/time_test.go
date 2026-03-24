package time

import (
	"testing"
	stdtime "time"


	"github.com/stretchr/testify/assert"
)

func TestUnixMsToTime(t *testing.T) {
	ts:=1684810374149
	d:=UnixMsToTime(int64(ts))
	assert.Equal(t,23,d.Day())
	assert.Equal(t,10,d.Hour())
	assert.Equal(t,52,d.Minute())
	assert.Equal(t,54,d.Second())
}

func TestNowDatetimeISO8601Str(t *testing.T) {
	value := NowDatetimeISO8601Str()

	_, err := stdtime.Parse(DatetimeISO8601Format, value)
	assert.NoError(t, err)
	assert.Contains(t, value, "T")
	assert.Contains(t, value, "+08:00")
}

