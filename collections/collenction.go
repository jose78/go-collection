package collections

import "fmt"

// MapperFunc is the default Mapper funtion
type MapperFunc func(interface{}, int) Result
// Action is the default Action funtion
type Action func(interface{}, int)
// Filter is the default Filter funtion
type Filter func(interface{}) bool 


// ListType is the default List
type ListType []interface{}

// MapType is the default Map
type MapType map[interface{}]interface{}

// Result is the default Result
type Result interface{}

// Tuple is the default
type Tuple struct {
	a interface{}
	b interface{}
} 


// Zip function returns a ListType object, which is an array of Tuples where the first item in each passed iterator is paired together, and then the second item in each passed iterator are paired together etc.
func Zip(a []interface{} , b []interface{}) (ListType , error){
	sizeA := len(a)
	if sizeA != len(b){
		return  nil , fmt.Errorf("Zip error, the length of two arrays show be the same len(a)=%d, len(b)=%d", sizeA, len(b))
	}
	list:= GenerateList()
	for index:= 0; index < sizeA; index++{
		list= list.Append(Tuple{a[index], b[index]})
	}
	return list, nil
}