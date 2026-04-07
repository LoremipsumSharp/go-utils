package gg

import (
	"fmt"
	"reflect"
)

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func IsZero[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}

func ToType[T any](v any) (T, bool) {
	var zero T
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return zero, false
	}
	targetType := reflect.TypeOf(zero)
	if val.Type() == targetType {
		return v.(T), true
	}
	if val.Type().ConvertibleTo(targetType) {
		return val.Convert(targetType).Interface().(T), true
	}
	return zero, false
}

func MustToType[T any](v any) T {
	if res, ok := ToType[T](v); ok {
		return res
	}
	var zero T
	targetType := reflect.TypeOf(zero)
	val := reflect.ValueOf(v)
	if val.IsValid() {
		panic(fmt.Sprintf("cannot convert value of type %s to %s", val.Type(), targetType))
	}
	panic(fmt.Sprintf("cannot convert invalid value to %s", targetType))
}
