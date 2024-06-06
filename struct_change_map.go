package struct_change_map

import "reflect"

func New(item interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	itemValue := reflect.ValueOf(item)

	// 如果傳入的是指標，則獲取其所指向的元素
	if itemValue.Kind() == reflect.Ptr {
		itemValue = itemValue.Elem()
	}
	// 確保傳入的是 struct
	if itemValue.Kind() != reflect.Struct {
		return result
	}
	itemType := itemValue.Type()
	for i := 0; i < itemValue.NumField(); i++ {
		if itemType.Field(i).PkgPath != "" {
			// 跳過未導出的字段
			continue
		}
		result[itemType.Field(i).Name] = itemValue.Field(i).Interface()
	}

	return result
}
