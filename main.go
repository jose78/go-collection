package main

import (
	"fmt"
)

type key interface{}
type value interface{}
type colMap map[key]value
type result interface{}
type mapper func(key, value, int) result

func main() {
	mm := make(colMap)
	mm["2"] = 8
	mm["3"] = 558
	mm["4"] = 28
	fmt.Println(converMap(mm, mapperd))
}

func mapperd(k key, v value, index int) result {
	return fmt.Sprintf("%d - RESULTADO %s %d", index, k, v)
}

func converMap(collection colMap, mapper mapper) []result {
	var res []result
	index := 0
	for k, v := range collection {
		res = append(res, mapper(k, v, index))
		index++
	}
	return res
}


