package collections

import (
	"reflect"
	"testing"
)

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

func filterOddNumber(item interface{}) bool {
	return item.(int)%3 == 0
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
	}{
		{"Fimd the last Odd", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumber}, 9, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.list.FilterLast(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.FilterLast() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListType.FilterLast() got1 = %v, want %v", got1, tt.want1)
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
	}{
		{"Fimd the first Odd", GenerateList(5, 1, 2, 3, 4, 7, 6, 5, 9, 67), args{filterOddNumber}, 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.list.FilterFirst(tt.args.fn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListType.FilterFirst() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListType.FilterFirst() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
