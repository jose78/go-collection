package collections


// MapperFunc is the default Mapper funtion
type MapperFunc func(interface{}, int) Result
// Action is the default Action funtion
type Action func(interface{}, int)
// Filter is the default Filter funtion
type Filter func(interface{}) bool 


// ListType is the default List
type ListType []interface{}

// Map is the default Map
type Map map[interface{}]interface{}

// Result is the default Result
type Result interface{}
