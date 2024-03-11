package expr

import (
	"context"
	"fmt"

	"github.com/expr-lang/expr"
)


func Eval[T any](ctx context.Context, exp string, vars map[string]interface{}) (retval T, reterr error) { //nolint:ireturn
	if len(exp) == 0 {
		return *new(T), nil
	}
	ex, err := expr.Compile(exp)
	if err != nil {
		return *new(T), fmt.Errorf("failed to compiled the expression:%v", err)
	}

	res, err := expr.Run(ex, vars)
	if err != nil {
		return *new(T), fmt.Errorf("evaluate expression: %w", err)
	}

	defer func() {
		if err := recover(); err != nil {
			retval = *new(T)
			reterr = fmt.Errorf("failed to evaluate the expression:%v", err)
		}
	}()

	return res.(T), nil
}