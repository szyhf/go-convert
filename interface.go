package convert // import "go.szyhf.org/di-convert"

import (
	"fmt"
	"reflect"
)

func MustInterfaceArray(array interface{}) (resArray []interface{}) {
	t := reflect.TypeOf(array)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(array)
			resArray = make([]interface{}, v.Len())
			for index, _ := range resArray {
				resArray[index] = v.Index(index).Interface()
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

func ToInterfaceArray(array interface{}) (res []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustInterfaceArray(array)
	return
}
