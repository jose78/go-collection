package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func generateMapTest() (container MapType) {
	container = MapType{}
	container[1] = testUser{"Alvaro", 6}
	container[2] = testUser{"Sofia", 3}
	container[3] = testUser{"empty", 0}
	return container
}

func generateResultMapTest() (container MapType) {
	container = MapType{}
	container[1] = testUser{"Alvaro", 6}
	container[2] = testUser{"Sofia", 3}
	return container
}

func filterMapOddNumber(item interface{}) bool {
	return item.(int)%3 == 0
}

func filterEmptyNameWithFail(key interface{}, value interface{}) bool {
	panic("This is another dummy error")
}

func filterEmptyName(key interface{}, value interface{}) bool {
	user := value.(testUser)
	return user.name != "empty"
}

func TestMapType_FilterAll(t *testing.T) {
	type args struct {
		fn func(interface{}, interface{}) bool
	}
	tests := []struct {
		name      string
		mapType   MapType
		args      args
		want      MapType
		flagError bool
	}{
		{"Fimd the last Odd", generateMapTest(), args{filterEmptyName}, generateResultMapTest(), false},
		{"Fimd the last Odd", generateMapTest(), args{filterEmptyNameWithFail}, generateResultMapTest(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := tt.mapType.FilterAll(tt.args.fn)
			if got2 == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapType.FilterAll() = %v, want %v", got, tt.want)
			}

			if got2 != nil && !tt.flagError {
				t.Errorf("error %v in %v", got2, tt.name)
			}
			//if got2 != nil && tt.flagError
		})
	}
}

func TestMapType_ListKeys(t *testing.T) {
	tests := []struct {
		name    string
		mapType MapType
		want    ListType
	}{
		{"Should return a list with the keys of the map", generateMapTest(), generateList(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mapType.ListKeys(); !checkIfAIsContainInB(got, tt.want) {
				t.Errorf("MapType.ListKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapType_ListValues(t *testing.T) {
	tests := []struct {
		name    string
		mapType MapType
		want    ListType
	}{
		{"Should return a list with the values of the map", generateMapTest(), generateList(testUser{"Alvaro", 6}, testUser{"empty", 0}, testUser{"Sofia", 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mapType.ListValues(); !checkIfAIsContainInB(got, tt.want) {
				t.Errorf("MapType.ListValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

var mapperMapToList FnMapperMap = func(fnKey, fnValue interface{}, index int) (key, value interface{}) {
	user := fnValue.(testUser)
	value = fmt.Sprintf("%s", user.name)
	return
}

var mapperMapToMap FnMapperMap = func(fnKey, fnValue interface{}, index int) (key, value interface{}) {
	user := fnValue.(testUser)
	value = fmt.Sprintf("%s", user.name)
	key = index
	return
}
var extracNamesWithError FnMapperMap = func(fnKey, fnValue interface{}, index int) (key, value interface{}) {
	user := fnValue.(testUser)
	if user.name == "empty" {
		panic("This is a dummy error")
	}
	value = fmt.Sprintf("%s", user.name)
	return
}

func TestMapType_Map(t *testing.T) {
	type args struct {
		fn FnMapperMap
	}
	tests := []struct {
		name    string
		mapType MapType
		args    args
		want    ListType
		wantErr bool
	}{

		{"Should return a list with the nams of each value", generateMapTest(), args{mapperMapToList}, generateList("Alvaro", "Sofia", "empty"), false},
		{"Should return a Map with the nams of each value and their IDś", generateMapTest(), args{mapperMapToList}, generateList("Alvaro", "Sofia", "empty"), false},
		{"Should fail ", generateMapTest(), args{extracNamesWithError}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mapType.Map(tt.args.fn)
			if (err != nil) && !tt.wantErr {
				t.Errorf("MapType.Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && !checkIfAIsContainInB(got, tt.want) {
				t.Errorf("MapType.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printEachItem(key, value interface{}, index int) {
	fmt.Printf("%d .- key: %v - value: %v", index, key, value)
}

func printEachItemWithError(key, value interface{}, index int) {
	panic(fmt.Sprintf("ERROR %v - %v", key, value))
}

func TestMapType_Foreach(t *testing.T) {
	type args struct {
		fn func(interface{}, interface{}, int)
	}
	tests := []struct {
		name    string
		mapType MapType
		args    args
		wantErr bool
	}{
		{"Should print each value of the map", generateMapTest(), args{printEachItem}, false},
		{"Should fail", generateMapTest(), args{printEachItemWithError}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mapType.Foreach(tt.args.fn)
		})
	}
}
