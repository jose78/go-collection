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
	//	var newList collections.ListType = collections.GenerateList(user{"Alvaro",6,1},user{"Sofia",3,2})
	newList := collections.GenerateList(user{"Alvaro", 6, 1}, user{"Sofia", 3, 2})
	results, err := newList.Map(mapperLst)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Println(results.Reverse().JoinAsString("(♥)"))
	fmt.Println(results)

	resultFiltered, index, _ := newList.FilterLast(filterUserByAge)
	fmt.Printf("result of filter %v with index %d\n", resultFiltered, index)

	listTuples, _ := collections.Zip(results.Reverse(), results)
	fmt.Println(listTuples.JoinAsString("(♥)"))
}



func filterUserByAge(value interface{}) bool {
	user := value.(user)
	return user.age > 3
}

var mapperLst collections.FnMapperList =  func (mapper interface{}, index int) interface{} {
	user1Item := mapper.(user)
	return user1Item.name
}
