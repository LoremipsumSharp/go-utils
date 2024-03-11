package expr

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	type foo struct {
		Val  int
		Code string
	}
	var env map[string]interface{} = map[string]interface{}{
		"code1": foo{
			Val:  1,
			Code: "code1",
		},
		"code2": foo{
			Val:  2,
			Code: "code2",
		},
	}
	expr := `code1.Val+code2.Val`
	sum, err := Eval[int](context.Background(), expr, env)
	assert.True(t, err == nil)
	assert.True(t, sum == 3)
}
