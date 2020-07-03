package collections


import "fmt"

// GenerateList is the default item
func GenerateList(items ...interface{}) ListType{
	return items
} 


//Map is the default
func (list ListType) Map(mapper func (interface{},int) interface{} ) []interface{}{
	result := make([]interface {}, len(list))
	for index, item := range list{
		fmt.Println(item)
		result = append(result, mapper(item, index))
	}
	return result
}