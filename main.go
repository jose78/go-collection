package main

import (
	"fmt"

	"github.com/jose78/go-collection/collections"
)

type user struct {
	name string
	age  int
	id   int
}

func main() {
	examplesWithList()
}

func examplesWithList() {
	//	var newList collections.ListType = collections.ParseItemsToList(user{"Alvaro",6,1},user{"Sofia",3,2})
	newList := collections.ParseItemsToList(user{"Alvaro", 6, 1}, user{"Sofia", 3, 2})
	resultsInter, err := newList.Map(mapperLst)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	results := resultsInter.(collections.ListType)
	fmt.Println(results.Reverse().JoinAsString("(♥)"))
	fmt.Println(results)

	resultInterMap, err := results.Map(mapperLstToMap)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Println(resultInterMap)
	resultMap := resultInterMap.(collections.MapType)
	resultMap.Foreach(printLoopMap)

	resultFiltered, index, _ := newList.FilterLast(filterUserByAge)
	fmt.Printf("result of filter %v with index %d\n", resultFiltered, index)


}

func print(item interface{}, index int){
	fmt.Println("item ->" , item)
} 



func filterUserByAge(value interface{}) bool {
	user := value.(user)
	return user.age > 3
}

var mapperLst collections.FnMapperList =  func (mapper interface{}, index int) (key, value interface{}) {
	user1Item := mapper.(user)
	value = user1Item.name
	return 
}


var mapperLstToMap collections.FnMapperList =  func (mapper interface{}, index int) (key, value interface{}) {
	value  = mapper
	key = index
	return 
}


var printLoopMap collections.FnForeachMap = func (key interface{}, value interface{}, index int){
	fmt.Printf("%d - %s \n" , key, value)
}