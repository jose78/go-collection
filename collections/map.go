package collections

// GenerateMapEmpty create an empty MapType
func GenerateMapEmpty() MapType {
	result := MapType{}
	return result
}

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
		result[tuple.a] =  tuple.b
	}
	return result
}

// GenerateMapFromZip is the default item
func GenerateMapFromZip(keys, values []interface{}) MapType {
	tuples, _ := Zip(keys, values)
	return GenerateMapFromTuples(tuples)
}

//Foreach is the default
func (mapType MapType) Foreach(fn func(interface{}, interface{}, int)) {
	index := 0
	for key, value := range mapType {
		fn(key, value, index)
		index++
	}
}

//Map is the default
func (mapType MapType) Map(fn func(interface{}, interface{}, int) interface{}) ListType {
	result := ListType{}
	index := 0
	for key, value := range mapType {
		result = append(result, fn(key, value, index))
		index++
	}
	return result
}

//FilterAll is the default
func (mapType MapType) FilterAll(fn func(interface{}, interface{}) bool) MapType {
	result := GenerateMapEmpty()
	for key, value := range mapType {
		if fn(key, value) {
			result[key] =  value
		}
	}
	return result
}
