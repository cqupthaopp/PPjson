package json

import (
	"errors"
	"reflect"
	"strconv"
)

func JSONUnMarshal(v interface{}, data string) error {

	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("It must be a ptr")
	}

	valType := reflect.TypeOf(v).Elem()
	valValu := reflect.ValueOf(v).Elem()

	if valType.Kind() == reflect.Slice || valType.Kind() == reflect.Array {

		var datas []interface{}

		PraiseDataToArray(&datas, data)

		valValu.Set(reflect.MakeSlice(valType, len(datas), len(datas)))

		for i := 0; i < len(datas); i++ {
			nowVal := valValu.Index(i)
			nowTyp := nowVal.Type()
			if nowTyp.Kind() != reflect.Ptr {
				nowVal = nowVal.Addr()
			}
			JSONUnMarshal(nowVal.Interface(), datas[i].(string))

		}

		return nil

	}

	if valType.Kind() == reflect.Map {

		datas := make(map[string]string)
		PraiseData(&datas, data)
		valValu.Set(reflect.MakeMapWithSize(valType, len(datas)))

		for k, v := range datas {

			nowKey := reflect.New(valType.Key())
			nowVal := reflect.New(valType.Key())

			JSONUnMarshal(nowKey.Interface(), k)
			JSONUnMarshal(nowVal.Interface(), v)

			valValu.SetMapIndex(nowKey.Elem(), nowVal.Elem())

		}
		return nil
	} //map

	if valType.Kind() == reflect.Struct {

		datas := make(map[string]string)
		PraiseData(&datas, data)

		for k := 0; k < valValu.NumField(); k++ {

			nowType := valType.Field(k)
			nowVal := valValu.Field(k)

			name := nowType.Name
			if (len(nowType.Tag.Get("json")) > 0) {
				name = nowType.Tag.Get("json")
			}

			if a, b := datas[name]; b {
				if nowVal.Kind() != reflect.Ptr {
					nowVal = nowVal.Addr()
					JSONUnMarshal(nowVal.Interface(), a)
				} //非指针

			}
		}
		return nil
	}
	//结构体

	if valType.Kind() == reflect.String {
		if data[0] == '"' && data[len(data)-1] == '"' {
			valValu.SetString(data[1 : len(data)-1])
			return nil
		}
		valValu.SetString(data)
	} //字符串

	if valType.Kind() == reflect.Int || valType.Kind() == reflect.Int8 || valType.Kind() == reflect.Int16 || valType.Kind() == reflect.Int32 || valType.Kind() == reflect.Int64 || valType.Kind() == reflect.Uint || valType.Kind() == reflect.Uint8 ||
		valType.Kind() == reflect.Uint16 || valType.Kind() == reflect.Uint32 || valType.Kind() == reflect.Uint64 {
		d, e := strconv.Atoi(data)
		if e == nil {
			valValu.SetInt(int64(d))
		}
		return e
	} //int

	if valType.Kind() == reflect.Float32 || valType.Kind() == reflect.Float64 {
		d, e := strconv.ParseFloat(data, 32)
		if e == nil {
			valValu.SetFloat(d)
		}
		return e
	} //Float

	if valType.Kind() == reflect.Bool {
		if data == "true" {
			valValu.SetBool(true)
		} else {
			valValu.SetBool(false)
		}
		return nil
	} //bool

	return nil
}
