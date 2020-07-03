package main

import (
	"fmt"
)

type colMap map[interface{}]interface{}
type colList []interface{}
type mapperList func(interface{}, int) interface{}
type mapperMap func(interface{}, interface{}, int) interface{}

// User is the basic structure for demo
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

	convertedMap := convertMap(mm, fnMapperMap)
	fmt.Println(convertedMap)
	fmt.Println(convertList(convertedMap, fnMapperList))
	fmt.Println(reverseList(convertList(convertedMap, fnMapperList)))
	fmt.Println(join(convertList(convertedMap, fnMapperList) , ", "))
	
	
	
	data := func (){
		fmt.Println("Hola")
	}
	
	
	data()

}

func fnMapperMap(k interface{}, v interface{}, index int) interface{}{
	user := v.(User)	
	return fmt.Sprintf("id:[%v], name:[%s]", user.id, user.name)
}

func fnMapperList(v interface{}, index int) interface{}{
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

func reverseList(lst []interface{}) []interface{} {
	var res []interface{}
	for index := len(lst) -1; index >= 0; index --{
		res = append (res , lst[index])
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
