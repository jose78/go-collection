package main

import (
	"fmt"
	"github.com/jose78/go-collenction-utils/collections"
)

type user struct{
	name string
	age int
	id int
}




func main() {

	lst := []user{{"Alvaro" , 1 , 1}, {"Sofi" , 2, 2}}



	for index, item := range lst{
		fmt.Printf("%d - %s \n" , index, item.name)
	}

	var newList collections.ListType = collections.GenerateList(user{"Alvaro",6,1},user{"Sofia",3,2})


	results := newList.Map(mapperLst)


	fmt.Println(results)
}




func mapperLst( mapper interface{} , index int) interface{}{
	user1Item := mapper.(user)
	return fmt.Sprintf ("%d -> Name:%s - Age:%d",index, user1Item.name , user1Item.age) 
}