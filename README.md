# go-collection

go-collection provides some types and methods to make easy work with collecctions. 

## Installation

Use the netst go command to download the lib:

```bash
go get github.com/jose78/go-collection
```

## Usage
This is a simple example of how to use the foreach method in lists: 
```go
package main

import (
	"fmt"

	col "github.com/jose78/go-collection/collections"
)

type user struct {
	name string
	age  int
	id   int
}

func main() {
	newList := col.GenerateList(user{"Alvaro", 6, 1}, user{"Sofia", 3, 2})
	newList = append(newList, user{"Mon", 0, 3})

	newList.Foreach(simpleLoop)
}

var simpleLoop col.FnForeachList = func(mapper interface{}, index int) {
	fmt.Printf("%d.- item:%v\n", index, mapper)
}
```

You can find a lot of examples in wiki section.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)