package json

import (
	"reflect"
)

func IsVaildType(kind reflect.Kind) bool {

	if kind == reflect.Chan || kind == reflect.Complex128 || kind == reflect.Complex64 || kind == reflect.Func || kind == reflect.Invalid {
		return false
	}
	return true
}

func PraiseData(datas *map[string]string, data string) {

	l := 1

	for k := 0; k < len(data); {

		if data[k] != ':' {
			k++
			continue
		}

		rr := k + 1

		if data[k+1] == '[' {
			for ; rr < len(data); rr++ {
				if data[rr] == ']' {
					break
				}
			}
			(*datas)[data[l:k]] = data[k+1 : rr+1]
			k = rr + 1
			l = rr + 1
			continue
		}

		if data[k+1] != '{' {
			for ; rr < len(data); rr++ {
				if data[rr] == ',' {
					break
				}
			}
			(*datas)[data[l:k]] = data[k+1 : rr]
			l = rr + 1
			k = l
			continue
		}

		cnt := 1

		for rr = k + 2; cnt > 0; rr++ {
			if data[rr] == '}' {
				cnt--
			} else if data[rr] == '{' {
				cnt++
			}
		}

		(*datas)[data[l:k]] = data[k+1 : rr+1]
		k = rr + 1
		l = k
	}

}

func PraiseDataToArray(datas *[]interface{}, data string) {

	*datas = make([]interface{}, 0)

	l := 1
	r := 1

	for l < len(data) && r < len(data) {

		if data[r] != ',' {
			r++
			continue
		}

		(*datas) = append((*datas), data[l:r])

		l = r + 1
		r = l

	}

}
