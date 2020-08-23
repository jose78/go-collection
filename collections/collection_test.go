package collections

import (
	"reflect"
	"testing"
)

func generateValues() []interface{} {
	lst := []interface{}{"one - 1", "two - 2", "three - 3", "another - xxxx"}
	return lst
}

func generateKeys() []interface{} {
	lst := []interface{}{"1", "2", "3", "another"}
	return lst
}

func generateMapResult() (container MapType) {
	container = MapType{}
	container["1"] = "one - 1"
	container["2"] = "two - 2"
	container["3"] = "three - 3"
	container["another"] = "another - xxxx"
	return container
}

func generateListTuples() []Tuple {
	result := []Tuple{Tuple{"1", "one - 1"}, Tuple{"2", "two - 2"}, Tuple{"3", "three - 3"}, Tuple{"another", "another - xxxx"}}
	return result
}

func TestGenerateMapFromZip(t *testing.T) {
	tests := []struct {
		name   string
		keys   []interface{}
		values []interface{}
		want   MapType
	}{
		{"generate a simple MapType", generateKeys(), generateValues(), generateMapResult()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMapFromZip(tt.keys, tt.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMapFromZip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		keys    []interface{}
		values  []interface{}
		want    []Tuple
		wantErr bool
	}{
		{"Generate a ListOfTouples from two Slices", generateKeys(), generateValues(), generateListTuples(), false},
		{"Should generate a fail", []interface{}{1, 2}, generateValues(), generateListTuples(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Zip(tt.keys, tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseListOfTupleToMap(t *testing.T) {
	tests := []struct {
		name       string
		tuples     []Tuple
		wantMapped MapType
		wantErr    bool
	}{
		{"Will convert to list of Tuples to map", generateListTuples(), generateMapResult(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMapped, err := ParseListOfTupleToMap(tt.tuples)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseListOfTupleToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMapped, tt.wantMapped) {
				t.Errorf("ParseListOfTupleToMap() = %v, want %v", gotMapped, tt.wantMapped)
			}
		})
	}
}

func Test_compareObjects(t *testing.T) {
	tests := []struct {
		name           string
		o1             interface{}
		o2             interface{}
		wantFlagEquals bool
	}{
		{"Should be the same", ListType{1, 2, 3, 4, 5}, ListType{1, 2, 3, 4, 5}, true},
		{"Should be false", generateMapTest(), generateMapResult(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlagEquals := compareObjects(tt.o1, tt.o2); gotFlagEquals != tt.wantFlagEquals {
				t.Errorf("compareObjects() = %v, want %v", gotFlagEquals, tt.wantFlagEquals)
			}
		})
	}
}
