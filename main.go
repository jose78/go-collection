package main

import (
	"fmt"
)


type key interface {}
type value interface {}
type colMap map[key]value
type result interface {} 
type mapper func(key , value) result 

func main() {
	mm := make(colMap)
	mm["2"] = 8	
	fmt.Println(fmap(mm, mapperd))
}

func mapperd(k key, v  value) result{
	return fmt.Sprintf("RESULTADO %s %d" , k , v)
}


func fmap(collection colMap, mapper mapper) []result{
	var res []result
    	for k,v := range collection {
		res = append(res, mapper(k, v))
	}
	return res
}



