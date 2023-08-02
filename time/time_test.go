package time

import (
	"testing"


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

