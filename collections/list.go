package collections

import "fmt"

// GenerateList is the default item
func GenerateList(items ...interface{}) ListType {
	return items
}

//Map is the default
func (list ListType) Map(mapper func(interface{}, int) interface{}) ListType {
	result := ListType{}
	for index, item := range list {
		fmt.Println(item)
		result = append(result, mapper(item, index))
	}
	return result
}

// Join is the default method
func (list ListType) Join(separator string) string {
	var result = ""
	var newSeparator = ""
	for _, value := range list {
		result = fmt.Sprintf("%s%s%v", result, newSeparator, value)
		newSeparator = separator
	}
	return result
}

// Reverse is the default method
func (list ListType) Reverse() ListType {
	var res []interface{}
	for index := len(list) - 1; index >= 0; index-- {
		res = append(res, list[index])
	}
	return res
}
