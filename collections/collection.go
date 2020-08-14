package collections

import ("fmt"
"reflect"
)

// FnMapperList define how you should implement a correct mapper for Listype
type FnMapperList func(interface{}, int) (interface{},interface{})

// FnMapperMap define how you should implement a correct mapper for  MapType
type FnMapperMap func(interface{}, interface{}, int) (interface{},interface{})

// FnFilterList this type define the struture the fucntion to implement if you want to filter the List
type FnFilterList func(interface{}) bool

// FnFilterMap this type define the struture the fucntion to implement if you want to filter the Map
type FnFilterMap func(interface{}, interface{}) bool

// FnForeachList define the function to call the foreach method of the ListType
type FnForeachList func(interface{}, int)

// FnForeachMap define the function to call the foreach method of the MapType
type FnForeachMap func(interface{}, interface{}, int)

// ResultType is the default Result object, used as generic object to encapsulate each returned object
type ResultType []interface{}

// ListType is the default List
type ListType []interface{}

// MapType is the default Map
type MapType map[interface{}]interface{}

// Tuple is the default
type Tuple struct {
	a interface{}
	b interface{}
}

// GenerateList is the default item
func GenerateList(items ...interface{}) ListType {
	return items
}

// GenerateMap is the default item
func GenerateMap(a, b interface{}) MapType {
	result := MapType{}
	result[a] = b
	return result
}

// Zip function returns a ListType object, which is an array of Tuples where the first item in each passed iterator is paired together, and then the second item in each passed iterator are paired together etc.
func Zip(a []interface{}, b []interface{}) (ListType, error) {
	sizeA := len(a)
	if sizeA != len(b) {
		return nil, fmt.Errorf("Zip error, the length of two arrays show be the same len(a)=%d, len(b)=%d", sizeA, len(b))
	}
	list := GenerateList()
	for index := 0; index < sizeA; index++ {
		list = append(list, Tuple{a[index], b[index]})
	}
	return list, nil
}



func compareObjects(o1, o2 interface{}) (flagEquals bool){

	if reflect.TypeOf(o1) ==  reflect.TypeOf(ListType{}) {
		flagEquals = reflect.DeepEqual(o1.(ListType), o2.(ListType))
	}else {
		flagEquals = reflect.DeepEqual(o1.(MapType), o2.(MapType))
	}
	return
}


func checkIfAIsContainInB(a,b interface{})  bool{

	if reflect.TypeOf(a) ==  reflect.TypeOf(ListType{}) {
		for mainItem := range a.(ListType){
			flagContained := false
			for item := range b.(ListType){
				if !flagContained{
					flagContained = reflect.DeepEqual(item , mainItem) 
				} 
			}
			if ! flagContained{
				return false
			}
		}
		return true
	}else if reflect.TypeOf(a) ==  reflect.TypeOf(MapType{}) {
		aMap := a.(MapType)
		bMap := b.(MapType)
		if len(aMap) == len(bMap){
			return false
		}
		flagContained := false
		for key := range b.(MapType){
			flagContained = aMap[key] == bMap[key]
			if ! flagContained{
				return false
			}
		}
		return true
	}
	return false
}