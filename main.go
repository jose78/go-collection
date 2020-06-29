package main

import (
	"fmt"
)

type colMap map[interface{}]interface{}
type colList []interface{}
type mapperList func(interface{}, int) interface{}
type mapperMap func(interface{}, interface{}, int) interface{}

func main() {
	mm := make(colMap)
	mm["2"] = 8
	mm[2] = 88888

	convertedMap := converMap(mm, mapperMap_d)
	fmt.Println(convertedMap)
	fmt.Println(converList(convertedMap, mapperList_d))
	fmt.Println(join(converList(convertedMap, mapperList_d) , ", "))

}

func mapperMap_d(k interface{}, v interface{}, index int) interface{}{
	return fmt.Sprintf("RESULTADO %v %d", k, v)
}

func mapperList_d(v interface{}, index int) interface{}{
	return fmt.Sprintf("%d - RESULT upodated => RESULTADO %v", index, v)
}

func converMap(collection colMap, mapper mapperMap) []interface{} {
	var res []interface{}
	index := 0
	for k, v := range collection {
		res = append(res, mapper(k, v, index))
		index++
	}
	return res
}

func converList(collection colList, mapper mapperList) []interface{} {
	var res []interface{}
	index := 0
	for v := range collection {
		res = append(res, mapper(v, index))
		index++
	}
	return res
}

func join(lst []interface{}, separator string) string{
	var res = ""
	var newSeparator = ""
	for _ , value := range lst{
		res = fmt.Sprintf("%s%s%v", res, newSeparator, value)
		newSeparator = separator
	}
	return res
}
