package gocollection

import (
	"errors"
	"fmt"
	"reflect"
)

// Touple represents a key-value pair with a generic key and value.
// The key must be of a comparable type, allowing it to be used in maps or
// other dest structures that require comparison operations.
// The value can be of any type.
//
// K - the type of the key, which must be comparable.
// V - the type of the value, which can be any type.
type Touple struct {
	Key   any // Key is the key of the key-value pair.
	Value any // Value is the value of the key-value pair.
}

// Mapper represents a function that takes a value of any type T and returns a value of any type.
// T - the type of the input value.
type Mapper[T any] func(T) any

// Predicate represents a function that takes a value of any type T and returns a boolean.
// It is used to test whether the input value satisfies a certain condition.
// T - the type of the input value.
type Predicate[T any] func(T) bool

// Action represents a function that takes an index and a value of any type T.
// It is intended to be used in an iteration context, such as a forEach function,
// where it will be executed for each element in a collection.
// T - the type of the input value.
type Action[T any] func(int, T)

// KeySelector represents a function that takes a key of type K and returns a Touple with the key and a value of type V.
// The key must be of a comparable type.
// K - the type of the key, which must be comparable.
// V - the type of the value.
type KeySelector[K comparable, V any] func(K) Touple

// Builder struct with an error and the item that caused the error
type Builder[T any] struct {
	err  error
	item T
}

// Method to retrieve the error from the builder
func (b *Builder[T]) Error() error {

	flag := b.err == nil

	fmt.Println(flag)

	return b.err
}

type ErrorFormatter[T any] func(T) error // New type for error formatting

// Method to set a custom error message in the builder using a function
func (b *Builder[T]) WithErrorMessage(fn ErrorFormatter[T]) *Builder[T] {
	if fn != nil {
		b.err = fn(b.item)
	}
	return b
}

// Map applies a Mapper function to each element in the source collection and stores the result in the dest collection.
// T - the type of the elements in the source collection.
// mapper - the function that transforms each element.
// source - the input collection of elements.
// dest - the output collection where the mapped elements are stored.
// Returns an error if the operation fails.
func Map[T any](mapper Mapper[T], source []T, dest *[]any) *Builder[T] {
	b := &Builder[T]{}
	for _, item := range source {
		defer func(i T) {
			if r := recover(); r != nil {
				b.err = fmt.Errorf("error mapping item: %v", r)
				b.item = i
			}
		}(item)
		*dest = append(*dest, mapper(item))
	}
	return b
}

// Filter applies a Predicate function to each element in the source collection and stores the elements that satisfy the predicate in the dest collection.
// T - the type of the elements in the source collection.
// predicate - the function that tests each element.
// source - the input collection of elements.
// dest - the output collection where the filtered elements are stored.
// Returns an error if the operation fails.
func Filter[T any](predicate Predicate[T], source []T, dest *[]T) *Builder[T] {
	b := &Builder[T]{}
	for _, item := range source {
		defer func(i T) {
			if r := recover(); r != nil {
				b.err = fmt.Errorf("error filtering item: %v", r)
				b.item = i
			}
		}(item)
		if predicate(item) {
			*dest = append(*dest, item)
		}
	}
	return b
}

// ForEach applies an Action function to each element in the source collection.
// T - the type of the elements in the source collection.
// action - the function that is executed for each element, taking the index and the element as parameters.
// source - the input collection of elements.
// Returns an error if the operation fails.
func ForEach[K any](action Action[K], src any) *Builder[K] {
	var errBuilder *Builder[K]
	evaluate := func(index int, internaParam any) {
		defer func(item any) {
			if err := recover(); err != nil {
				valueParametrized := item.(K)
				errBuilder = &Builder[K]{
					item: valueParametrized,
					err:  fmt.Errorf("Error: iterating within the forEach item:%v:  -->  %v", item, err),
				}
			}
		}(internaParam)
		action(index, internaParam.(K))
	}
	if isMap(src) {
		val := reflect.ValueOf(src)
		count := -1
		for _, key := range val.MapKeys() {
			count++
			value := val.MapIndex(key)
			touple := Touple{key, value}
			evaluate(count, touple)
		}
	} else {
		for index, item := range src.([]K) {
			if !reflect.ValueOf(errBuilder).IsZero() {
				break
			}
			if !reflect.ValueOf(errBuilder).IsNil() {
				break
			}
			evaluate(index, item)
		}
	}
	return errBuilder
}

// Zip combines two slices into a map, using elements from the keys slice as keys and elements from the values slice as values.
// K - the type of the keys, which must be comparable.
// V - the type of the values.
// keys - the slice of keys.
// values - the slice of values.
// result - the map where the keys and values are combined.
// Returns an error if the operation fails, such as when the lengths of keys and values do not match.
func Zip[K comparable, V any](keys []K, values []V, result map[K]V) *Builder[K] {
	b := &Builder[K]{}
	if len(keys) != len(values) {
		b.err = errors.New("keys and values slices must have the same length")
		return b
	}
	for i, key := range keys {
		defer func(k K, v V) {
			if r := recover(); r != nil {
				b.err = fmt.Errorf("error zipping item: %v", r)
				b.item = k
			}
		}(key, values[i])
		result[key] = values[i]
	}
	return b
}

func isMap(elements any) bool {
	t := reflect.TypeOf(elements)
	return reflect.Map == t.Kind()
}

func store(data any, dest any) {
	if isMap(dest) {
		val := reflect.ValueOf(dest)
		keyVal := reflect.ValueOf(data.(Touple).Key)
		valueVal := reflect.ValueOf(data.(Touple).Value)
		val.SetMapIndex(keyVal, valueVal)
	} else {
		sliceVal := reflect.ValueOf(dest).Elem()
		elemVal := reflect.ValueOf(data)
		result := reflect.Append(sliceVal, elemVal)
		sliceVal.Set(result)
	}
}

func Filter2[T any](predicate Predicate[T], source any, dest any) *Builder[T] {
	var action Action[T] = func(index int, item T) {
		if predicate(item) {
			fmt.Println(dest)
			store(item, dest)
			fmt.Println(dest)
		}
	}
	return ForEach[T](action, source)
}
