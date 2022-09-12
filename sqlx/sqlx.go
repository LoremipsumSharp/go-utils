package sqlx

import (
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

var ErrPtr = fmt.Errorf("must pass by reference")

func Select[T any](db *sqlx.DB, query string, params ...any) ([]T, error) {
	var dest []T
	return dest, db.Select(&dest, query, params...)
}

func Get[T any](db *sqlx.DB, query string, params ...any) (*T, error) {
	var v T
	return &v, db.Get(&v, query, params...)
}

func Query[T any](db *sqlx.DB, dest T, query string, params ...any) error {
	if rv := reflect.ValueOf(dest); rv.Kind() == reflect.Ptr {
		switch reflect.Indirect(rv).Kind() {
		case reflect.Slice:
			return db.Select(dest, query, params...)
		default:
			return db.Get(dest, query, params...)
		}
	}
	return ErrPtr
}
