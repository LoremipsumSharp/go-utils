package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	s := struct {
		Foo string `json:"foo,omitempty"`
	}{
		Foo: "Bar",
	}
	m,err:=ToMap(s)
	assert.Nil(t,err)
	assert.True(t,len(m)==1)
	assert.True(t,m["foo"] == "Bar")
}
