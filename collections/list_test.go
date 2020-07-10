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

func TestListType_Append(t *testing.T) {
	tests := []struct {
		name string
		list ListType
		args interface{}
		want ListType
	}{
		{"Append element to list", GenerateList(1, 2), 3, GenerateList(1, 2, 3)},
		{"Append element to empty list", GenerateList(), 3, GenerateList(3)},
		{"Append duplicate element to list", GenerateList(3), 3, GenerateList(3, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Append(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}


func mapperInt(item interface{}, index int) interface{} {
	value := item.(int)
	return value * 10
}

func mapperUser(item interface{}, index int) interface{} {
	user := item.(testUser)
	return user.name
}

func mapperUserWithFails(item interface{}, index int) interface{} {
	panic(fmt.Errorf("This is a Dummy fail -> %v", item))
}

func TestListType_Map(t *testing.T) {
	type args struct {
		mapper func(interface{}, int) interface{}
	}
	tests := []struct {
		name string
		list ListType
		args args
		want ListType
		want1  string
	}{
		{"Should retrive the name of each testUser", GenerateList(testUser{"Alvaro", 6}, testUser{"Sofi", 3}), args{mapperUser}, GenerateList("Alvaro", "Sofi"), ""},
		{"Should retrive a list with each number *10", GenerateList(3, 4, 5, 6), args{mapperInt}, GenerateList(30, 40, 50, 60) , ""},
		{"Should fail", GenerateList(testUser{"Alvaro", 6}, testUser{"Sofi", 3}), args{mapperUserWithFails}, nil, "This is a Dummy fail"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.list.Map(tt.args.mapper)
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.Map() = %v, want %v", got, tt.want)
			}
			if  got1 !=  nil && got1.Error() ==  tt.want1 {
				t.Errorf("ListType.Map() = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func doSomething(item interface{}, index int) {
	fmt.Printf("%d - value:%v", index, item)

}

func factorListType() ListType {
	list := GenerateList()
	for index := 1; index <= 1000000; index++ {
		list.Append(index)
	}
	return list
}

func TestListType_Foreach(t *testing.T) {
	type args struct {
		fn func(interface{}, int)
	}
	tests := []struct {
		name string
		list ListType
		args args
	}{
		{"Should  execute for each item the same operation", factorListType(), args{doSomething}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Foreach(tt.args.fn)
		})
	}
}

func TestListType_Join(t *testing.T) {
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
			if got := tt.list.Join(tt.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.Join() = %v, want %v", got, tt.want)
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
		fn func(interface{}) bool
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
		fn func(interface{}) bool
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