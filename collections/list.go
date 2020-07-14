package collections

import (
	"errors"
	"fmt"
)



//Foreach method performs the given action for each element of the array/slice until all elements have been processed or the action generates an exception.
func (list ListType) Foreach(action FnForeachList) error {
	for index, item := range list {
		fmt.Printf("%+v\n", index)
		if err := callbackForeach(index, item, action); err != nil {
			return err
		}
	}
	return nil
}

func callbackForeach(index int, item interface{}, fnInternal FnForeachList) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			fmt.Printf("ERROR: %v", err)
		}
	}()
	fnInternal(item, index)
	return err
}

func callbackMap(index int, value interface{}, fnInternal FnMapperList) (item interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			fmt.Printf("ERROR: %v", err)
		}
	}()
	item = fnInternal(value, index)
	return item, err
}

//Map function iterates through a ListType, converting each element into a new value using the function as the transformer.
func (list ListType) Map(mapper FnMapperList) (ListType, error) {
	result := GenerateList()
	for index, item := range list {
		itemMapped, err := callbackMap(index, item, mapper)
		if err != nil {
			return nil, err
		}

		result = append(result, itemMapped)

	}
	return result, nil
}

// JoinAsString concatenates all elements as string .
func (list ListType) JoinAsString(separator string) string {
	var result = ""
	var newSeparator = ""
	for _, value := range list {
		result = fmt.Sprintf("%s%s%v", result, newSeparator, value)
		newSeparator = separator
	}
	return result
}

// Reverse Create a new ListType that is the reverse the elements of the original ListType.
func (list ListType) Reverse() ListType {
	var res []interface{}
	for index := len(list) - 1; index >= 0; index-- {
		res = append(res, list[index])
	}
	return res
}

//FilterAll method finds all ocurrences in a collection that matches with the function criteria.
func (list ListType) FilterAll(fn FnFilterListType) ListType {
	result := ListType{}
	for _, item := range list {
		if fn(item) {
			result = append(result, fn(item))
		}
	}
	return result
}

//FilterFirst method finds the first ocurrence in a collection that matches with the function criteria. If any iteration fails, it wil return "nil, INDEX_OF_ITERATION, error" ELSE if FIND OK ITEM_SELECTED, INDEX_OF_ITEM , nil ELSE nil, -1, nil
func (list ListType) FilterFirst(fn FnFilterListType) (interface{}, int, error) {
	for index := 0; index < len(list); index++ {
		if flag, err := callbackFilter(index, list[index], fn); err != nil {
			return nil, index, err
		} else if flag {
			return list[index], index, nil
		}
	}
	return nil, -1, nil
}

func callbackFilter(index int, value interface{}, fnInternal FnFilterListType) (flag bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			fmt.Printf("ERROR: %v", err)
		}
	}()
	flag = fnInternal(value)
	return flag, err
}

//FilterLast method finds the first ocurrence in a collection that matches with the function criteria. If any iteration fails, it wil return "nil, INDEX_OF_ITERATION, error" ELSE if FIND OK ITEM_SELECTED, INDEX_OF_ITEM , nil ELSE nil, -1, nil
func (list ListType) FilterLast(fn FnFilterListType) (interface{}, int, error) {
	for index := len(list) - 1; index >= 0; index-- {
		if flag, err := callbackFilter(index, list[index], fn); err != nil {
			return nil, index, err
		} else if flag {
			return list[index], index, nil
		}
	}
	return nil, -1, nil
}
