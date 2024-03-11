package gg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	var foo Foo
	assert.True(t, IsZero(foo))

	foo.Bar = "bar"
	assert.False(t, IsZero(foo))
}

type Foo struct {
	Bar string
}
