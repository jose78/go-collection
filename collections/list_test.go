package collections

import (
	"fmt"
	"reflect"
	"testing"
)

type testUser struct {
	name string
	age  int
}

var mapperInt FnMapperList =  func (item interface{}, index int) (key, value interface{}) {
	value = item.(int) * 10
	return 
}

var mapperListToMap FnMapperList =  func (item interface{}, index int) (key, value interface{}) {
	user := item.(testUser)
	value = user.name
	key = index * 100
	return 
}

var mapperUserWithFails FnMapperList = func (item interface{}, index int) (key, value interface{}) {
	panic(fmt.Errorf("This is a Dummy fail -> %v", item))
}

func  buildDefaultResultMap()  MapType{
	result := MapType{}
	result[0] = "Alvaro"
	result[100] = "Sofi"
	return result
}
func TestListType_Map(t *testing.T) {
	type args struct {
		mapper FnMapperList
	}
	tests := []struct {
		name  string
		list  ListType
		args  args
		want  interface{}
		want1 bool
	}{
		{"Should retrive the name of each testUser", GenerateList(testUser{"Alvaro", 6}, testUser{"Sofi", 3}), args{mapperListToMap}, buildDefaultResultMap(), false},
		{"Should retrive a list with each number *10", GenerateList(3, 4, 5, 6), args{mapperInt}, GenerateList(30, 40, 50, 60), false },
		{"Should fail", GenerateList(testUser{"Alvaro", 6}, testUser{"Sofi", 3}), args{mapperUserWithFails}, nil,  true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.Map(tt.args.mapper)
			if err != nil &&  !tt.want1 {
				t.Errorf("ListType.Map() = %v, want %v", err, tt.want1)
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListType_JoinAsString(t *testing.T) {
	tests := []struct {
		name      string
		list      ListType
		separator string
		want      string
	}{
		{"Should retrive the name of each testUser", GenerateList("Alvaro", "Sofi"), ",", "Alvaro,Sofi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.JoinAsString(tt.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.JoinAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListType_Reverse(t *testing.T) {
	tests := []struct {
		name string
		list ListType
		want ListType
	}{
		{"1º - Should generate a new ListTypewith inverted values", GenerateList("Alvaro", "Sofi"), GenerateList("Sofi", "Alvaro")},
		{"2º - Should generate a new ListTypewith inverted values", GenerateList(testUser{"Alvaro", 6}, testUser{"Sofi", 3}), GenerateList(testUser{"Sofi", 3}, testUser{"Alvaro", 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func filterOddNumber(item interface{}) bool {
	return item.(int)%3 == 0
}

func filterOddNumberWithError(item interface{}) bool {
	if item.(int)%3 == 0 {
		panic(fmt.Errorf("This is a Dummy fail -> %v", item))
	} else {
		return false
	}
}

func TestListType_FilterLast(t *testing.T) {
	type args struct {
		fn FnFilterList
	}
	tests := []struct {
		name  string
		list  ListType
		args  args
		want  interface{}
		want1 int
		want2 string
	}{
		{"Fimd the last Odd", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumber}, 9, 8, ""},
		{"It should manage the fail", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumberWithError}, nil, 8, "This is a Dummy fail -> 9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.list.FilterLast(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.FilterLast() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListType.FilterLast() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != nil && got2.Error() != tt.want2 {
				t.Errorf("ListType.FilterLast() got2 = %v(%T), want %v(%T)", got2, got2, tt.want2, tt.want2)
			}
		})
	}
}

func TestListType_FilterFirst(t *testing.T) {
	type args struct {
		fn FnFilterList
	}
	tests := []struct {
		name  string
		list  ListType
		args  args
		want  interface{}
		want1 int
		want2 string
	}{
		{"Fimd the first Odd", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumber}, 3, 3, ""},
		{"It should manage the fail", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumberWithError}, nil, 3, "This is a Dummy fail -> 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.list.FilterFirst(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.FilterLast() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListType.FilterLast() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != nil && got2.Error() != tt.want2 {
				t.Errorf("ListType.FilterLast() got2 = %v(%T), want %v(%T)", got2, got2, tt.want2, tt.want2)
			}
		})
	}
}

func factorListType() ListType {
	list := GenerateList()
	for index := 1; index <= 100; index++ {
		list = append(list, index)
	}
	return list
}

func doSomethingWithPanic(item interface{}, index int) {
	if index%3 != 0 {
		panic(fmt.Errorf("This is a Dummy fail -> %v", item))
	}

	fmt.Printf("%d - value:%v", index, item)

}

func doSomething(item interface{}, index int) {
	fmt.Printf("%d - value:%v", index, item)
}

func TestListType_Foreach(t *testing.T) {
	type args struct {
		action func(interface{}, int)
	}
	tests := []struct {
		name    string
		list    ListType
		args    args
		wantErr bool
	}{
		{"Should  execute for each item the same operation", factorListType(), args{doSomething}, false},
		{"Should  be failed ", factorListType(), args{doSomethingWithPanic}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.list.Foreach(tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListType.Foreach() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
