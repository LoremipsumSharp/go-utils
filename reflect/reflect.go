package  reflect

import "reflect"


func IsStruct(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() == reflect.Struct
}

func GetTagValue(e interface{}, tagName string) []interface{} {
	rez := make([]interface{}, 0)
	t := reflect.TypeOf(e)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		if len(tag) > 0 {
			rez = append(rez, tag)
		}
	}

	return rez
}

func IsPointerToStruct(p interface{}) bool {

	pv := reflect.ValueOf(p)
	pvk := pv.Kind()

	if pvk != reflect.Ptr {
		return false
	}

	vv := pv.Elem()
	vvk := vv.Kind()

	if vvk != reflect.Struct {
		return false
	}

	return true
}