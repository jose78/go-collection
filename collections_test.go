package gocollection

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var (
	john  = testUser{name: "John", secondName: "Connor", mails: []string{}, age: 10, male: true}
	sarah = testUser{name: "Sarah", secondName: "Connor", mails: []string{}, age: 43}
	kyle  = testUser{name: "Kyle", secondName: "Risk", male: true, mails: []string{}, age: 43}
)

type testUser struct {
	name       string
	secondName string
	mails      []string
	age        int
	male       bool
}

func generateTestCaseList() []testUser {
	return []testUser{john, sarah, kyle}
}

func generateTestCaseMap() map[string]testUser {
	result := map[string]testUser{}
	for _, item := range generateTestCaseList() {
		result[item.name] = item
	}
	return result
}

func errorFmtAsString(user string) error {
	return fmt.Errorf("KO")
}

func errorFmtAsTouple(user Touple) error {
	return fmt.Errorf("KO")
}

func errorFmt(user testUser) error {
	return fmt.Errorf("KO")
}

func TestZip(t *testing.T) {
	keys := []string{"a", "b", "c"}
	values := []int{1, 2, 3}
	result := make(map[string]int)
	builder := Zip(keys, values, result)

	if err := builder.Error(); err != nil {
		t.Fatalf("Zip failed: %v", err)
	}

	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Zip result mismatch for key %v: got %v, want %v", k, result[k], v)
		}
	}
}

func TestWithErrorMessage(t *testing.T) {
	source := []int{1, 2, 3, 4}
	var dest []any

	mapper := func(n int) any {
		if n == 3 {
			panic("unexpected value")
		}
		return n * 2
	}
	builder := Map(mapper, source, &dest)

	customErrFunc := func(item int) error {
		return fmt.Errorf("custom error for item %d", item)
	}

	err := builder.WithErrorMessage(customErrFunc).Error()
	if err == nil || err.Error() != "custom error for item 3" {
		t.Fatalf("expected custom error for item 3, got %v", err)
	}
}

func TestForEach(t *testing.T) {

	var actionOk Action[testUser] = func(i int, tu testUser) {
		fmt.Printf("index: %d. User: %v", i, tu)
	}

	var actionKO Action[testUser] = func(i int, tu testUser) {
		if i == 0 {
			fmt.Print(tu.mails[3])
		}
		fmt.Printf("index: %d. User: %v", i, tu)
	}

	src := generateTestCaseList()

	type args struct {
		action   Action[testUser]
		errorFmt ErrorFormatter[testUser]
		src      any
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"Iterate over list of testUser", args{action: actionOk, src: src}, nil},
		{"Iterate and generate and customizable error", args{action: actionKO, errorFmt: errorFmt, src: src}, fmt.Errorf("KO")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ForEach(tt.args.action, tt.args.src)
			if tt.want != nil && !reflect.ValueOf(got).IsZero() {
				err := got.WithErrorMessage(tt.args.errorFmt).Error()
				if err.Error() != tt.want.Error() {
					t.Errorf("Each() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

type argsFilter[T any] struct {
	predicate Predicate[T]
	errorFmt  ErrorFormatter[T]
	source    any
	dest      any
}
type testsTypeFilter[T any] struct {
	name      string
	args      argsFilter[T]
	want      any
	wantError bool
	err       error
}

func (tt testsTypeFilter[T]) runTestFilter(testRunner *testing.T) {
	testRunner.Run(tt.name, func(t *testing.T) {
		got := Filter(tt.args.predicate, tt.args.source, tt.args.dest)
		if !tt.wantError && got != nil {
			t.Errorf("filter() = %v, wantError %v", got, tt.wantError)
		}
		if !tt.wantError && !reflect.DeepEqual(tt.want, tt.args.dest) {
			t.Errorf("Filter() = %v, want %v", tt.args.dest, tt.want)
		}
	},
	)
}

func isMale(tu testUser) bool {
	return tu.male
}

func isMaleAsTouple(tu Touple) bool {
	testUser := tu.Value.(testUser)
	flag := testUser.male
	return flag
}

func TestFilter2(t *testing.T) {

	resultListFiltered := []testUser{
		john,
		kyle,
	}
	resultMapFiltered := map[string]testUser{"John": john, "Kyle": kyle}

	lstUsers := []testUser{}
	mapUsers := map[string]testUser{}

	testsTypeFilter[testUser]{
		name:      "Filter male from list of test user",
		args:      argsFilter[testUser]{isMale, errorFmt, generateTestCaseList(), &lstUsers},
		want:      &resultListFiltered,
		wantError: false,
		err:       nil}.runTestFilter(t)

	testsTypeFilter[Touple]{
		name:      "Filter male from map of test user",
		args:      argsFilter[Touple]{isMaleAsTouple, errorFmtAsTouple, generateTestCaseMap(), mapUsers},
		want:      resultMapFiltered,
		wantError: false,
		err:       nil,
	}.runTestFilter(t)
}

type argsMap[T any] struct {
	mapper   Mapper[T]
	errorFmt ErrorFormatter[T]
	source   any
	dest     any
}
type testsTypeMap[T any] struct {
	name      string
	args      argsMap[T]
	want      any
	wantError bool
	err       error
}

func (tt testsTypeMap[T]) runTestMap(testRunner *testing.T) {
	testRunner.Run(tt.name, func(t *testing.T) {
		got := Map(tt.args.mapper, tt.args.source, tt.args.dest)
		if !tt.wantError && got != nil {
			t.Errorf("Map() KO = %v, wantError %v", got, tt.wantError)
		}
		if !tt.wantError && !reflect.DeepEqual(tt.want, tt.args.dest) {
			t.Errorf("Map() = %v, want %v", tt.args.dest, tt.want)
		}
	},
	)
}

var mapperToNamesFromMap Mapper[Touple] = func(s Touple) any {
	user := s.Value.(testUser)
	return fmt.Sprintf("%s %s", user.name, user.secondName)
}

var mapperToNamesFromList Mapper[testUser] = func(s testUser) any {
	return fmt.Sprintf("%s %s", s.name, s.secondName)
}

var mapperSplitName Mapper[string] = func(s string) any {
	nameSplited := strings.Split(s, " ")
	return Touple{nameSplited[0], nameSplited[1]}
}

// Test functions
func TestMap(t *testing.T) {
	result := []string{"John Connor", "Sarah Connor", "Kyle Risk"}
	lstUsersFromList := []string{}
	lstUsersFromMap := []string{}
	names := map[string]string{}

	resulNames := map[string]string{"Sarah": "Connor", "Kyle": "Risk", "John": "Connor"}

	testsTypeMap[testUser]{
		name:      "Filter male from list of test user",
		args:      argsMap[testUser]{mapperToNamesFromList, errorFmt, generateTestCaseList(), &lstUsersFromList},
		want:      &result,
		wantError: false,
		err:       nil}.runTestMap(t)

	testsTypeMap[Touple]{
		name:      "Filter male from map of test user",
		args:      argsMap[Touple]{mapperToNamesFromMap, errorFmtAsTouple, generateTestCaseMap(), &lstUsersFromMap},
		want:      &result,
		wantError: false,
		err:       nil,
	}.runTestMap(t)

	testsTypeMap[string]{
		name:      "Map list of names and seconds names to map ",
		args:      argsMap[string]{mapperSplitName, errorFmtAsString, lstUsersFromMap, names},
		want:      resulNames,
		wantError: false,
		err:       nil,
	}.runTestMap(t)

}

type argsGroupBy[T any] struct {
	groupBy  KeySelector[T]
	errorFmt ErrorFormatter[T]
	source   any
	dest     any
}
type testsTypeGroupBy[T any] struct {
	name      string
	args      argsGroupBy[T]
	want      any
	wantError bool
	err       error
}

func (tt testsTypeGroupBy[T]) runTestGroupBy(testRunner *testing.T) {
	testRunner.Run(tt.name, func(t *testing.T) {
		got := GroupBy(tt.args.groupBy, tt.args.source, tt.args.dest)
		if !tt.wantError && got != nil {
			t.Errorf("Map() KO = %v, wantError %v", got, tt.wantError)
		}
		if !tt.wantError && !reflect.DeepEqual(tt.want, tt.args.dest) {
			t.Errorf("Map() = %v, want %v", tt.args.dest, tt.want)
		}
	},
	)
}

var keySelectorBySex KeySelector[testUser] = func(tu testUser) any {
	result := "female"
	if tu.male {
		result = "male"
	}
	return result
}

func TestGroupBy(t *testing.T) {

	result := map[string][]testUser{"male": {john, kyle}, "female": {sarah}}

	testsTypeGroupBy[testUser]{
		name:      "Filter male from list of test user",
		args:      argsGroupBy[testUser]{keySelectorBySex, errorFmt, generateTestCaseList(), map[string][]testUser{}},
		want:      result,
		wantError: false,
		err:       nil}.runTestGroupBy(t)

}

