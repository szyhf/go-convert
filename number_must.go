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
