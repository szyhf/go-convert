package convert // import "go.szyhf.org/di-convert"

import "reflect"

// 把map转成map[string]interface{}，key的值使用MustString计算。
// 如果子项中也有map，则继续递归执行直到全部转换为map[string]interface{}
// 常用于各种xml\yaml\json转换为map的结果的统一处理。
func MustMapStringInterfaceRecursions(leafMap interface{}) map[string]interface{} {
	leafType := reflect.TypeOf(leafMap)
	if leafType.Kind() != reflect.Map {
		return nil
	}
	leafValue := reflect.ValueOf(leafMap)
	if leafValue.Len() == 0 {
		return nil
	}

	resMap := make(map[string]interface{})
	leafKeyValues := leafValue.MapKeys()
	// key的value
	for _, leafKeyValue := range leafKeyValues {
		// node的value
		nodeValue := leafValue.MapIndex(leafKeyValue)

		// 获得实际的key和node
		k := leafKeyValue.Interface()
		node := nodeValue.Interface()

		strKey := MustString(k)
		nodeType := reflect.TypeOf(node)
		if nodeType.Kind() == reflect.Map {
			temp := MustMapStringInterfaceRecursions(node)
			if temp != nil {
				resMap[strKey] = temp
			}
		} else {
			resMap[strKey] = node
		}
	}
	return resMap
}
