package json

import (
	"errors"
	"fmt"
	"reflect"
)

func JSONMarshal(v interface{}) (string, error) {

	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Struct {
		ans := ""
		val := reflect.ValueOf(v)
		for k := 0; k < val.NumField(); k++ {
			now_val := val.Field(k)
			now_key := val.Type().Field(k)

			if !now_val.CanInterface() || !IsVaildType(now_key.Type.Kind()) {
				continue
			}

			now_ans, err := JSONMarshal(now_val.Interface())
			if err != nil {
				fmt.Println(err)
				continue
			}

			name := now_key.Name
			if len(val.Type().Field(k).Tag.Get("json")) > 0 {
				name = val.Type().Field(k).Tag.Get("json")
			}

			ans += name + ":" + now_ans + ","
		}
		return "{" + ans + "}", nil
	}
	//结构体

	if t.Kind() == reflect.Int || t.Kind() == reflect.Int8 || t.Kind() == reflect.Int16 || t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64 || t.Kind() == reflect.Uint || t.Kind() == reflect.Uint8 ||
		t.Kind() == reflect.Uint16 || t.Kind() == reflect.Uint32 || t.Kind() == reflect.Uint64 || t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64 || t.Kind() == reflect.Bool {
		return fmt.Sprintf("%v", v), nil
	}
	//普通类型

	if t.Kind() == reflect.String {
		return fmt.Sprintf("\"%v\"", v), nil
	}
	//字符串

	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		ans := ""
		val := reflect.ValueOf(v)

		for k := 0; k < val.Len(); k++ {
			now_ans, err := JSONMarshal(val.Index(k).Interface())
			if err != nil {
				fmt.Println(err)
				continue
			}
			ans += now_ans + ","
		}
		return "[" + ans + "]", nil
	}
	//数组切片

	if t.Kind() == reflect.Map {
		val := reflect.ValueOf(v)
		keys := val.MapKeys()
		ans := ""
		for _, nowKey := range keys {
			nowVal := val.MapIndex(nowKey)
			if !nowVal.IsValid() || !nowVal.CanInterface() {
				continue
			}

			nowAns, err := JSONMarshal(nowVal.Interface())
			if err != nil {
				fmt.Println(err)
				continue
			}
			ans += nowKey.String() + ":" + nowAns + ","
		}
		return "{" + ans + "}", nil
	}
	//map

	return "", errors.New("No match")
}
