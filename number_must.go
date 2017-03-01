package convert

// 强制转换为int，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt(value interface{}) (res int) {
	if res, err := ToInt(value); err != nil {
		panic(err)
	} else {
		return res
	}
}

// 强制转换为int8，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int8
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt8(value interface{}) (res int8) {
	if res, err := ToInt8(value); err != nil {
		panic(err)
	} else {
		return res
	}
}

// 强制转换为int16，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int16
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt16(value interface{}) (res int16) {
	if res, err := ToInt16(value); err != nil {
		panic(err)
	} else {
		return res
	}
}

// 强制转换为int32，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int32
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt32(value interface{}) (res int32) {
	if res, err := ToInt32(value); err != nil {
		panic(err)
	} else {
		return res
	}
}

// 强制转换为int64，失败则panic
// string会按顺序尝试将数据解析为int64\uint64\float64\bool，然后再转换为int64
// float会抹去小数
// uint8~64、int8~64都会做默认的转换
// bool类型的数据，true-1；false-0
func MustInt64(value interface{}) (res int64) {
	if res, err := ToInt64(value); err != nil {
		panic(err)
	} else {
		return res
	}
}
