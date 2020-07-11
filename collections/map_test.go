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

func filterEmptyName(key interface{}, value interface{}) bool {
	user := value.(testUser)
	return user.name != "empty"
}

func TestMapType_FilterAll(t *testing.T) {
	type args struct {
		fn func(interface{}, interface{}) bool
	}
	tests := []struct {
		name    string
		mapType MapType
		args    args
		want    MapType
	}{
		{"Fimd the last Odd", generateMapTest(), args{filterEmptyName}, generateResultMapTest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mapType.FilterAll(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapType.FilterAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printEachItem(key, value interface{}, index int) {
	fmt.Printf("%d .- key: %v - value: %v", index, key, value)
}

func TestMapType_Foreach(t *testing.T) {
	type args struct {
		fn func(interface{}, interface{}, int)
	}
	tests := []struct {
		name    string
		mapType MapType
		args    args
	}{
		{"Should print each value of the map", generateMapTest(), args{printEachItem}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mapType.Foreach(tt.args.fn)
		})
	}
}

func extracNames(key, value interface{}, index int) interface{} {
	user := value.(testUser)
	return fmt.Sprintf("%s" ,user.name)
}

func TestMapType_Map(t *testing.T) {
	type args struct {
		fn func(interface{}, interface{}, int) interface{}
	}
	tests := []struct {
		name    string
		mapType MapType
		args    args
		want    ListType
	}{
		{"Should return a list with the nams of each value" , generateMapTest(), args{extracNames}, GenerateList("Alvaro","Sofia","empty")},
	}
	for _, tt := range tests {
		fmt.Println(tt.want)
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mapType.Map(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapType.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

