package collections

import (
	"errors"
)

// GenerateMap is the default item
func GenerateMap(a, b interface{}) MapType {
	result := MapType{}
	result[a] = b
	return result
}


// GenerateMapFromTuples is the default item
func GenerateMapFromTuples(tuples ListType) MapType {
	result := MapType{}
	for _, item := range tuples {
		tuple := item.(Tuple)
		result[tuple.a] = tuple.b
	}
	return result
}

// GenerateMapFromZip is the default item
func GenerateMapFromZip(keys, values []interface{}) MapType {
	tuples, _ := Zip(keys, values)
	return GenerateMapFromTuples(tuples)
}

func callbackMapTypeForeach(index int, key, value interface{}, fnInternal FnForeachMap) (err error) {
	defer func() {
		if r := recover(); r != nil {

			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()
	fnInternal(key, value, index)
	return err
}

//Foreach method performs the given action for each element of the map  until all elements have been processed or the action generates an exception.
func (mapType MapType) Foreach(fn FnForeachMap) error {
	index := 0

	for key, value := range mapType {
		if err := callbackMapTypeForeach(index, key, value, fn); err != nil {
			return err
		}
		index++
	}
	return nil
}

func callbackMapTypeMap(index int, fnKey, fnValue interface{}, fnInternal FnMapperMap) (key,value  interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {

			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()
	key, value  = fnInternal(fnKey, fnValue, index)
	return 
}

//Map function iterates through a ListType, converting each element into a new value using the function as the transformer.
func (mapType MapType) Map(fn FnMapperMap) (result interface{}, err error) {
	resultList := ListType{}
	resultMap := MapType{}
	flagFirstIteratioun := true
	var flagIsMap bool 
	index := 0
	for loopKey, loopValue := range mapType {
		key, value , err := callbackMapTypeMap(index, loopKey, loopValue, fn)
		if flagFirstIteratioun{
			flagFirstIteratioun = false
			flagIsMap =  key != nil
		}
		if err != nil {
			return nil, err
		} else {
			if flagIsMap{
				resultMap[key] = value
			}else {
				resultList = append(resultList, value)	
			}
		}
		index++
	}
	if flagIsMap{
		result = resultMap
	}else {
		result = resultList
	}
	return result, nil
}

//FilterAll method finds all ocurrences in a collection that matches with the function criteria.
func (mapType MapType) FilterAll(fn FnFilterMap) MapType {
	result := MapType{}
	for key, value := range mapType {
		if fn(key, value) {
			result[key] = value
		}
	}
	return result
}

//ListValues obtains a ListType of the values in this Map.
func (mapType MapType) ListValues() ListType {
	list := ListType{}
	for _, value := range mapType {
		list = append(list, value)
	}
	return list
}

//ListKeys obtains a ListType of the keys in this Map.
func (mapType MapType) ListKeys() ListType {
	list := ListType{}
	for key := range mapType {
		list = append(list, key)
	}
	return list
}
