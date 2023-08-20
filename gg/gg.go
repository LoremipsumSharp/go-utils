package gg

import "reflect"

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func IsZero[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}