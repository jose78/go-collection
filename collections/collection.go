package collections

import (
	"fmt"
	"reflect"
)

// FnMapperList define how you should implement a correct mapper for Listype
type FnMapperList func(interface{}, int) (interface{}, interface{})

// FnMapperMap define how you should implement a correct mapper for  MapType
type FnMapperMap func(interface{}, interface{}, int) (interface{}, interface{})

// FnFilterList this type define the struture the function to implement if you want to filter the List
type FnFilterList func(interface{}) bool

// FnFilterMap this type define the struture the function to implement if you want to filter the Map
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

// ParseList Convert to ListType another Slice created prevouosly
func ParseList(items ...interface{}) ListType {
	return items
}

// ParseMap Convert to MapType another MAP created prevously
func ParseMap(items map[interface{}]interface{}) MapType {
	return items
}

// ParseItemsToList is the default item
func ParseItemsToList(items ...interface{}) ListType {
	return items
}

// ParseListOfTupleToMap Create a Map from Slice of Tuples
func ParseListOfTupleToMap(tuples []Tuple) (mapped MapType, err error) {
	mapped = MapType{}
	for _, tuple := range tuples {
		mapped[tuple.a] = tuple.b
	}
	return
}

// GenerateMapFromZip is the default item
func GenerateMapFromZip(keys, values []interface{}) MapType {
	tuples, _ := Zip(keys, values)
	if mapped, err := ParseListOfTupleToMap(tuples); err != nil {
		panic(err)
	} else {
		return mapped
	}
}

// Zip merge the elements into Slice of Tuples, each key will be stored inside of the a filed of Tupule Strunct and each value will be stored un b field
func Zip(keys []interface{}, values []interface{}) ([]Tuple, error) {
	sizeA := len(keys)
	if sizeA != len(values) {
		return nil, fmt.Errorf("Zip error, the length of two arrays show be the same len(a)=%d, len(b)=%d", sizeA, len(values))
	}
	list := []Tuple{}
	for index := 0; index < sizeA; index++ {
		list = append(list, Tuple{keys[index], values[index]})
	}
	return list, nil
}

func compareObjects(o1, o2 interface{}) (flagEquals bool) {

	if reflect.TypeOf(o1) == reflect.TypeOf(ListType{}) {
		flagEquals = reflect.DeepEqual(o1.(ListType), o2.(ListType))
	} else {
		flagEquals = reflect.DeepEqual(o1.(MapType), o2.(MapType))
	}
	return
}

func checkIfAIsContainInB(a, b interface{}) bool {

	if reflect.TypeOf(a) == reflect.TypeOf(ListType{}) {
		for mainItem := range a.(ListType) {
			flagContained := false
			for item := range b.(ListType) {
				if !flagContained {
					flagContained = reflect.DeepEqual(item, mainItem)
				}
			}
			if !flagContained {
				return false
			}
		}
		return true
	} else if reflect.TypeOf(a) == reflect.TypeOf(MapType{}) {
		aMap := a.(MapType)
		bMap := b.(MapType)
		if len(aMap) == len(bMap) {
			return false
		}
		var flagContained bool
		for key := range b.(MapType) {
			flagContained = aMap[key] == bMap[key]
			if !flagContained {
				return false
			}
		}
		return true
	}
	return false
}
