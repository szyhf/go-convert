package convert // import "go.szyhf.org/di-convert"

import (
	"fmt"
	"reflect"
	"strconv"
)

// 尽最大努力将一个值转为int类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt(value interface{}) (res int, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to int failed", value.(string))
						} else {
							res, err = ToInt(resBool)
						}
					} else {
						res, err = ToInt(resF64)
					}
				} else {
					res, err = ToInt(resU64)
				}

			} else {
				res = int(res64)
			}
		}
	case int:
		{
			res = int(value.(int))
		}
	case int8:
		{
			res = int(value.(int8))
		}
	case int16:
		{
			res = int(value.(int16))
		}
	case int32:
		{
			res = int(value.(int32))
		}
	case int64:
		{
			res = int(value.(int64))
		}
	case uint:
		{
			res = int(value.(uint))
		}
	case uint8:
		{
			res = int(value.(uint8))
		}
	case uint16:
		{
			res = int(value.(uint16))
		}
	case uint32:
		{
			res = int(value.(uint32))
		}
	case uint64:
		{
			res = int(value.(uint64))
		}
	case uintptr:
		{
			res = int(value.(uintptr))
		}

	case float32:
		{
			res = int(value.(float32))
		}
	case float64:
		{
			res = int(value.(float64))
		}

	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int8类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt8(value interface{}) (res int8, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to int8 failed", value.(string))
						} else {
							res, err = ToInt8(resBool)
						}
					} else {
						res, err = ToInt8(resF64)
					}
				} else {
					res, err = ToInt8(resU64)
				}

			} else {
				res = int8(res64)
			}
		}
	case int:
		{
			res = int8(value.(int))
		}
	case int8:
		{
			res = int8(value.(int8))
		}
	case int16:
		{
			res = int8(value.(int16))
		}
	case int32:
		{
			res = int8(value.(int32))
		}
	case int64:
		{
			res = int8(value.(int64))
		}
	case uint:
		{
			res = int8(value.(uint))
		}
	case uint8:
		{
			res = int8(value.(uint8))
		}
	case uint16:
		{
			res = int8(value.(uint16))
		}
	case uint32:
		{
			res = int8(value.(uint32))
		}
	case uint64:
		{
			res = int8(value.(uint64))
		}
	case uintptr:
		{
			res = int8(value.(uintptr))
		}

	case float32:
		{
			res = int8(value.(float32))
		}
	case float64:
		{
			res = int8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{

			valueStr := MustString(value)
			res, err = ToInt8(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int16类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt16(value interface{}) (res int16, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to int16 failed", value.(string))
						} else {
							res, err = ToInt16(resBool)
						}
					} else {
						res, err = ToInt16(resF64)
					}
				} else {
					res, err = ToInt16(resU64)
				}

			} else {
				res = int16(res64)
			}
		}
	case int:
		{
			res = int16(value.(int))
		}
	case int8:
		{
			res = int16(value.(int8))
		}
	case int16:
		{
			res = int16(value.(int16))
		}
	case int32:
		{
			res = int16(value.(int32))
		}
	case int64:
		{
			res = int16(value.(int64))
		}
	case uint:
		{
			res = int16(value.(uint))
		}
	case uint8:
		{
			res = int16(value.(uint8))
		}
	case uint16:
		{
			res = int16(value.(uint16))
		}
	case uint32:
		{
			res = int16(value.(uint32))
		}
	case uint64:
		{
			res = int16(value.(uint64))
		}
	case uintptr:
		{
			res = int16(value.(uintptr))
		}

	case float32:
		{
			res = int16(value.(float32))
		}
	case float64:
		{
			res = int16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt16(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int32类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt32(value interface{}) (res int32, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to int32 failed", value.(string))
						} else {
							res, err = ToInt32(resBool)
						}
					} else {
						res, err = ToInt32(resF64)
					}
				} else {
					res, err = ToInt32(resU64)
				}

			} else {
				res = int32(res64)
			}
		}
	case int:
		{
			res = int32(value.(int))
		}
	case int8:
		{
			res = int32(value.(int8))
		}
	case int16:
		{
			res = int32(value.(int16))
		}
	case int32:
		{
			res = int32(value.(int32))
		}
	case int64:
		{
			res = int32(value.(int64))
		}
	case uint:
		{
			res = int32(value.(uint))
		}
	case uint8:
		{
			res = int32(value.(uint8))
		}
	case uint16:
		{
			res = int32(value.(uint16))
		}
	case uint32:
		{
			res = int32(value.(uint32))
		}
	case uint64:
		{
			res = int32(value.(uint64))
		}
	case uintptr:
		{
			res = int32(value.(uintptr))
		}

	case float32:
		{
			res = int32(value.(float32))
		}
	case float64:
		{
			res = int32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt32(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为int64类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToInt64(value interface{}) (res int64, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to int64 failed", value.(string))
						} else {
							res, err = ToInt64(resBool)
						}
					} else {
						res, err = ToInt64(resF64)
					}
				} else {
					res, err = ToInt64(resU64)
				}

			} else {
				res = int64(res64)
			}
		}
	case int:
		{
			res = int64(value.(int))
		}
	case int8:
		{
			res = int64(value.(int8))
		}
	case int16:
		{
			res = int64(value.(int16))
		}
	case int32:
		{
			res = int64(value.(int64))
		}
	case uint:
		{
			res = int64(value.(uint))
		}
	case uint8:
		{
			res = int64(value.(uint8))
		}
	case uint16:
		{
			res = int64(value.(uint16))
		}
	case uint32:
		{
			res = int64(value.(uint32))
		}
	case uint64:
		{
			res = int64(value.(uint64))
		}
	case uintptr:
		{
			res = int64(value.(uintptr))
		}

	case float32:
		{
			res = int64(value.(float32))
		}
	case float64:
		{
			res = int64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToInt64(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint(value interface{}) (res uint, err error) {
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint failed", value.(string))
						} else {
							res, err = ToUint(resBool)
						}
					} else {
						res, err = ToUint(resF64)
					}
				} else {
					res, err = ToUint(res64)
				}
			} else {
				res = uint(resU64)
			}
		}
	case int:
		{
			res = uint(value.(int))
		}
	case int8:
		{
			res = uint(value.(int8))
		}
	case int16:
		{
			res = uint(value.(int16))
		}
	case int32:
		{
			res = uint(value.(int32))
		}
	case int64:
		{
			res = uint(value.(int64))
		}

	case uint:
		{
			res = uint(value.(uint))
		}
	case uint8:
		{
			res = uint(value.(uint8))
		}
	case uint16:
		{
			res = uint(value.(uint16))
		}
	case uint32:
		{
			res = uint(value.(uint32))
		}
	case uint64:
		{
			res = uint(value.(uint64))
		}
	case uintptr:
		{
			res = uint(value.(uintptr))
		}

	case float32:
		{
			res = uint(value.(float32))
		}
	case float64:
		{
			res = uint(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint(valueStr)
		}
	}
	return
}

// 尽最大努力将一个值转为uint8类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint8(value interface{}) (res uint8, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint8(value)
	return
}

// 尽最大努力将一个值转为uint16类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint16(value interface{}) (res uint16, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint16(value)
	return
}

// 尽最大努力将一个值转为uint32类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint32(value interface{}) (res uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint32(value)
	return
}

// 尽最大努力将一个值转为uint64类型的数据
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func ToUint64(value interface{}) (res uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint64(value)
	return
}

// 尽最大努力将目标转换为[]int
func ToIntArray(value interface{}) (res []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustIntArray(value)
	return
}

// 尽最大努力将目标转换为[]int8
func ToInt8Array(value interface{}) (res []int8, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustInt8Array(value)
	return
}

// 尽最大努力将目标转换为[]int16
func ToInt16Array(value interface{}) (res []int16, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustInt16Array(value)
	return
}

// 尽最大努力将目标转换为[]int32
func ToInt32Array(value interface{}) (res []int32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustInt32Array(value)
	return
}

// 尽最大努力将目标转换为[]int64
func ToInt64Array(value interface{}) (res []int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustInt64Array(value)
	return
}

//尽最大努力将目标转换为[]uint
func ToUintArray(value interface{}) (res []uint, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUintArray(value)
	return
}

//尽最大努力将目标转换为[]uint8
func ToUint8Array(value interface{}) (res []uint8, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint8Array(value)
	return
}

//尽最大努力将目标转换为[]uint16
func ToUint16Array(value interface{}) (res []uint16, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint16Array(value)
	return
}

//尽最大努力将目标转换为[]uint32
func ToUint32Array(value interface{}) (res []uint32, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint32Array(value)
	return
}

//尽最大努力将目标转换为[]uint64
func ToUint64Array(value interface{}) (res []uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustUint64Array(value)
	return
}

// 尽最大努力将目标转换为[]int，失败会panic
func MustIntArray(value interface{}) (resArray []int) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustInt(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

// 尽最大努力将目标转换为[]int8，失败会panic
func MustInt8Array(value interface{}) (resArray []int8) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int8, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustInt8(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

// 尽最大努力将目标转换为[]int16，失败会panic
func MustInt16Array(value interface{}) (resArray []int16) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int16, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustInt16(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

// 尽最大努力将目标转换为[]int32，失败会panic
func MustInt32Array(value interface{}) (resArray []int32) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int32, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustInt32(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

// 尽最大努力将目标转换为[]int64，失败会panic
func MustInt64Array(value interface{}) (resArray []int64) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int64, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustInt64(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

//尽最大努力将目标转换为[]uint,失败会panic
func MustUintArray(value interface{}) (resArray []uint) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustUint(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

//尽最大努力将目标转换为[]uint8,失败会panic
func MustUint8Array(value interface{}) (resArray []uint8) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint8, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustUint8(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

//尽最大努力将目标转换为[]uint16,失败会panic
func MustUint16Array(value interface{}) (resArray []uint16) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint16, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustUint16(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

//尽最大努力将目标转换为[]uint32,失败会panic
func MustUint32Array(value interface{}) (resArray []uint32) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint32, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustUint32(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

//尽最大努力将目标转换为[]uint64,失败会panic
func MustUint64Array(value interface{}) (resArray []uint64) {

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint64, v.Len())
			for index, _ := range resArray {
				resArray[index] = MustUint64(v.Index(index).Interface())
			}
			return
		}
	}
	panic("Not an array or slice")
	return
}

// 强制转换为uint8，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint8(value interface{}) (res uint8) {
	var err error
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint8 failed", value.(string))
						} else {
							res = MustUint8(resBool)
						}
					} else {
						res = MustUint8(resF64)
					}
				} else {
					res = MustUint8(res64)
				}
			} else {
				res = uint8(resU64)
			}
		}
	case int:
		{
			res = uint8(value.(int))
		}
	case int8:
		{
			res = uint8(value.(int8))
		}
	case int16:
		{
			res = uint8(value.(int16))
		}
	case int32:
		{
			res = uint8(value.(int32))
		}
	case int64:
		{
			res = uint8(value.(int64))
		}

	case uint:
		{
			res = uint8(value.(uint))
		}
	case uint8:
		{
			res = uint8(value.(uint8))
		}
	case uint16:
		{
			res = uint8(value.(uint16))
		}
	case uint32:
		{
			res = uint8(value.(uint32))
		}
	case uint64:
		{
			res = uint8(value.(uint64))
		}
	case uintptr:
		{
			res = uint8(value.(uintptr))
		}

	case float32:
		{
			res = uint8(value.(float32))
		}
	case float64:
		{
			res = uint8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint8(valueStr)
		}
	}

	if err != nil {
		panic(err)
	}
	return
}

// 强制转换为uint16，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint16(value interface{}) (res uint16) {
	var err error
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint16 failed", value.(string))
						} else {
							res = MustUint16(resBool)
						}
					} else {
						res = MustUint16(resF64)
					}
				} else {
					res = MustUint16(res64)
				}
			} else {
				res = uint16(resU64)
			}
		}
	case int:
		{
			res = uint16(value.(int))
		}
	case int8:
		{
			res = uint16(value.(int8))
		}
	case int16:
		{
			res = uint16(value.(int16))
		}
	case int32:
		{
			res = uint16(value.(int32))
		}
	case int64:
		{
			res = uint16(value.(int64))
		}

	case uint:
		{
			res = uint16(value.(uint))
		}
	case uint8:
		{
			res = uint16(value.(uint8))
		}
	case uint16:
		{
			res = uint16(value.(uint16))
		}
	case uint32:
		{
			res = uint16(value.(uint32))
		}
	case uint64:
		{
			res = uint16(value.(uint64))
		}
	case uintptr:
		{
			res = uint16(value.(uintptr))
		}

	case float32:
		{
			res = uint16(value.(float32))
		}
	case float64:
		{
			res = uint16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint16(valueStr)
		}
	}

	if err != nil {
		panic(err)
	}
	return
}

// 强制转换为uint32，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint32(value interface{}) (res uint32) {
	var err error
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint32 failed", value.(string))
						} else {
							res = MustUint32(resBool)
						}
					} else {
						res = MustUint32(resF64)
					}
				} else {
					res = MustUint32(res64)
				}
			} else {
				res = uint32(resU64)
			}
		}
	case int:
		{
			res = uint32(value.(int))
		}
	case int8:
		{
			res = uint32(value.(int8))
		}
	case int16:
		{
			res = uint32(value.(int16))
		}
	case int32:
		{
			res = uint32(value.(int32))
		}
	case int64:
		{
			res = uint32(value.(int64))
		}

	case uint:
		{
			res = uint32(value.(uint))
		}
	case uint8:
		{
			res = uint32(value.(uint8))
		}
	case uint16:
		{
			res = uint32(value.(uint16))
		}
	case uint32:
		{
			res = uint32(value.(uint32))
		}
	case uint64:
		{
			res = uint32(value.(uint64))
		}
	case uintptr:
		{
			res = uint32(value.(uintptr))
		}

	case float32:
		{
			res = uint32(value.(float32))
		}
	case float64:
		{
			res = uint32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint32(valueStr)
		}
	}

	if err != nil {
		panic(err)
	}
	return
}

// 强制转换为uint64，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为uint64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustUint64(value interface{}) (res uint64) {
	var err error
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
				if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {
					if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint64 failed", value.(string))
						} else {
							res = MustUint64(resBool)
						}
					} else {
						res = MustUint64(resF64)
					}
				} else {
					res = MustUint64(res64)
				}
			} else {
				res = uint64(resU64)
			}
		}
	case int:
		{
			res = uint64(value.(int))
		}
	case int8:
		{
			res = uint64(value.(int8))
		}
	case int16:
		{
			res = uint64(value.(int16))
		}
	case int32:
		{
			res = uint64(value.(int32))
		}
	case int64:
		{
			res = uint64(value.(int64))
		}

	case uint:
		{
			res = uint64(value.(uint))
		}
	case uint8:
		{
			res = uint64(value.(uint8))
		}
	case uint16:
		{
			res = uint64(value.(uint16))
		}
	case uint32:
		{
			res = uint64(value.(uint32))
		}
	case uint64:
		{
			res = uint64(value.(uint64))
		}
	case uintptr:
		{
			res = uint64(value.(uintptr))
		}

	case float32:
		{
			res = uint64(value.(float32))
		}
	case float64:
		{
			res = uint64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToUint64(valueStr)
		}
	}

	if err != nil {
		panic(err)
	}
	return
}

func MustFloat64(value interface{}) (res float64) {
	var err error
	switch value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0
			}
			//string的情况比较复杂，需要继续区分string是bool、float、int、uint再处理
			if resF64, erro := strconv.ParseFloat(valueString, 0); erro != nil {
				if resU64, erro := strconv.ParseUint(valueString, 0, 0); erro != nil {
					if res64, erro := strconv.ParseInt(valueString, 0, 0); erro != nil {

						if resBool, erro := strconv.ParseBool(valueString); erro != nil {
							err = fmt.Errorf("Convert string \"%s\" to Uint64 failed", value.(string))
						} else {
							res = MustFloat64(resBool)
						}
					} else {
						res = MustFloat64(res64)
					}
				} else {
					res = MustFloat64(resU64)
				}
			} else {
				res = MustFloat64(resF64)
			}
		}
	case int:
		{
			res = float64(value.(int))
		}
	case int8:
		{
			res = float64(value.(int8))
		}
	case int16:
		{
			res = float64(value.(int16))
		}
	case int32:
		{
			res = float64(value.(int32))
		}
	case int64:
		{
			res = float64(value.(int64))
		}

	case uint:
		{
			res = float64(value.(uint))
		}
	case uint8:
		{
			res = float64(value.(uint8))
		}
	case uint16:
		{
			res = float64(value.(uint16))
		}
	case uint32:
		{
			res = float64(value.(uint32))
		}
	case uint64:
		{
			res = float64(value.(uint64))
		}
	case uintptr:
		{
			res = float64(value.(uintptr))
		}

	case float32:
		{
			res = float64(value.(float32))
		}
	case float64:
		{
			res = float64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := MustString(value)
			res, err = ToFloat64(valueStr)
		}
	}

	if err != nil {
		panic(err)
	}
	return
}

func ToFloat64(value interface{}) (res float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	res = MustFloat64(value)
	return
}
