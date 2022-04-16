package string

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder:=Builder{}
	builder.AppendString("Hello ")
	builder.AppendString("World ")
	builder.PrependString("Xu ")
	builder.PrependString("Loremipsum ")
	builder.AppendStringIf(false,"are ")
	builder.AppendStringIf(false,"Q ")
	fmt.Println(builder.String())
}
