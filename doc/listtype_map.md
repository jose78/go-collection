### ListType - Map
Map function iterates through a ListType, converting each element into a new value using the function as the transformer.


```go
package mai

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
	//	var newList collections.ListType = collections.ParseItemsToList(user{"Alvaro",6,1},user{"Sofia",3,2})
	newList := collections.ParseItemsToList(user{"Alvaro", 6, 1}, user{"Sofia", 3, 2})
	results, err := newList.Map(mapperLst)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Println(results)
}

func mapperLst(mapper interface{}, index int) interface{} {
	user1Item := mapper.(user)
	return user1Item.name
}
```
Result of execution

```bash
$ go run main.go 
[Alvaro Sofia]
```