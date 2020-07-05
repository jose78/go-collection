package collections

import "fmt"

// GenerateList is the default item
func GenerateList(items ...interface{}) ListType {
	return items
}

//Foreach is the default
func (list ListType) Foreach(fn func(interface{}, int)) {
	for index, item := range list {
		fn(item, index)
	}
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

//FilterAll is the default
func (list ListType) FilterAll(fn func(interface{}) bool) ListType {
	result := ListType{}
	for _, item := range list {
		if fn(item) {
			result = append(result, fn(item))
		}
	}
	return result
}

//FilterFirst is the default
func (list ListType) FilterFirst(fn func(interface{}) bool) (interface{}, int) {
	for index := 0; index < len(list); index++ {
		if fn(list[index]) {
			return list[index], index
		}
	}
	return nil, -1
}

//FilterLast is the default
func (list ListType) FilterLast(fn func(interface{}) bool) (interface{}, int) {
	for index := len(list) - 1; index >= 0; index-- {
		if fn(list[index]) {
			return list[index], index
		}
	}
	return nil, -1
}

// Append is the default way to insesrt elements
func (list ListType) Append(item interface{}) ListType {
	return append(list, item)
}
