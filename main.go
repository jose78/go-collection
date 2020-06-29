package main

import (
	"fmt"
)

type colMap map[interface{}]interface{}
type colList []interface{}
type mapperList func(interface{}, int) interface{}
type mapperMap func(interface{}, interface{}, int) interface{}


type User struct{
	name string
	age int
	id int
}

func main() {
	mm := make(colMap)
	mm[1] = User{"Mon", 40 ,1}
	mm[3] = User{"Alvaro", 6 ,3}
	mm[4] = User{"Sofi", 3 ,4}

	convertedMap := converMap(mm, mapperMap_d)
	fmt.Println(convertedMap)
	fmt.Println(converList(convertedMap, mapperList_d))
	fmt.Println(join(converList(convertedMap, mapperList_d) , ", "))
	
	
	
	data := func (){
		fmt.Println("Hola")
	}
	
	
	data()

}

func mapperMap_d(k interface{}, v interface{}, index int) interface{}{
	user := v.(User)	
	return fmt.Sprintf("id:[%v], name:[%s]", user.id, user.name)
}

func mapperList_d(v interface{}, index int) interface{}{
	return fmt.Sprintf("%d - RESULT upodated => RESULTADO %v", index, v)
}

func convertMap(collection colMap, mapper mapperMap) []interface{} {
	var res []interface{}
	index := 0
	for k, v := range collection {
		res = append(res, mapper(k, v, index))
		index++
	}
	return res
}

func convertList(collection colList, mapper mapperList) []interface{} {
	var res []interface{}
	for index, v := range collection {
		res = append(res, mapper(v,index))
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
