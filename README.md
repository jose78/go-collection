[![Go Report Card](https://goreportcard.com/badge/github.com/jose78/go-collection)](https://goreportcard.com/report/github.com/jose78/go-collection)
[![Coverage Status](https://coveralls.io/repos/github/jose78/go-collection/badge.svg?branch=master)](https://coveralls.io/github/jose78/go-collection?branch=master)
[![CircleCI](https://circleci.com/gh/jose78/go-collection.svg?style=shield)](https://circleci.com/gh/jose78/go-collection)
[![Go Reference](https://pkg.go.dev/badge/github.com/jose78/go-collection/v2.svg)](https://pkg.go.dev/github.com/jose78/go-collection/v2)

# go-collection <img align="right" width="80" height="100" src="resources/gopher.png">




The `go-collection` package provides a set of utility functions for working with collections (slices and maps) in Go. It includes functions for mapping, filtering, zipping, and performing actions on collections.

Installation
------------

To install the package, use:

```bash
    go get github.com/jose78/go-collection
```

Usage
-----

### Types

#### Touple

Represents a key-value pair with a generic key and value.

```go
    type Touple struct {
        Key   any // Key is the key of the key-value pair.
        Value any // Value is the value of the key-value pair.
    }
```

#### Comparator
```go
type Comparator[T any] func(T, T) int 
```

Defines a comparison function for elements of type T. The function takes two arguments of type T and returns an integer:
- A negative value indicates that the first argument is less than the second.
- A zero value indicates that both arguments are equal.
- A positive value indicates that the first argument is greater than the second.



#### Mapper

Represents a function that takes a value of any type `T` and returns a value of any type.

```go
    type Mapper[T any] func(T) any
```

#### Predicate

Represents a function that takes a value of any type `T` and returns a boolean. It is used to test whether the input value satisfies a certain condition.

```go
    type Predicate[T any] func(T) bool
```

#### Action

Represents a function that takes an index and a value of any type `T`. It is intended to be used in an iteration context, such as a forEach function, where it will be executed for each element in a collection.

```go
    type Action[T any] func(int, T)
```

#### KeySelector

Represents a function that takes a key of type `K` and returns a `Touple` with the key and a value of type `V`. The key must be of a comparable type.

```go
    type KeySelector[K comparable, V any] func(K) Touple
```

#### Builder

A struct with an error and the item that caused the error.
```go
    type Builder[T any] struct {
        err  error
        item T
    }
    
    func (b *Builder[T]) Error() error {
        return b.err
    }
    
    type ErrorFormatter[T any] func(T) error // New type for error formatting
    
    func (b *Builder[T]) WithErrorMessage(fn ErrorFormatter[T]) *Builder[T] {
        if fn != nil {
            b.err = fn(b.item)
        }
        return b
    }
```
### Functions

#### Map

Applies a `Mapper` function to each element in the source collection and stores the result in the dest collection.

```go
    func Map[T any](mapper Mapper[T], source []T, dest *[]any) *Builder[T]
```

#### ForEach

Applies an `Action` function to each element in the source collection.

```go
    func ForEach[K any](action Action[K], src any) *Builder[K]
```

#### SortBy


 Comparator is a generic type that defines a comparison function for elements of type T.
 The function takes two arguments of type T and returns an integer:
 - A negative value indicates that the first argument is less than the second.
 - A zero value indicates that both arguments are equal.
 - A positive value indicates that the first argument is greater than the second.
type Comparator[T any] func(T, T) int

 SortBy sorts the elements in the `source` using the provided `comparator`.

 Parameters:
 - comparator: A comparison function of type Comparator[T], where T is the type of the elements in the `source`.
 - source: A value of type `any` that must be a map or a pointer to a list (slice or array).

 Returns:
 - An error if `source` is not a supported type or if an issue occurs during sorting.

 Description:
 The `SortBy` function accepts a `comparator` to determine the order of elements.
 The `source` can be:
 - A map whose keys will be sorted according to the `comparator`.
 - A pointer to a list (slice or array) of elements of type T.

 Usage Example:
 ```go

	comparator := func(a, b int) int {
	    return a - b
	}

 slice := []int{3, 1, 2}
 err := SortBy(comparator, &slice)

	if err != nil {
	    log.Fatal(err)
	}

  slice is now []int{1, 2, 3}
 ```

 If `source` is not a map or a pointer to a list, the function returns an error.
func SortBy[T any](comparator Comparator[T], source any) {
	lst := source.(*[]T)
	sort.Slice(*lst, func(i, j int) bool {
		return comparator((*lst)[i], (*lst)[j]) < 0
	})

}


#### Zip

Combines two slices into a map, using elements from the keys slice as keys and elements from the values slice as values.

```go
    func Zip[K comparable, V any](keys []K, values []V, result map[K]V) *Builder[K]
```

#### isMap

Checks if the given element is of map type.

```go
    func isMap(elements any) bool
```



#### Filter

Applies a `Predicate` function to each element in the source collection and stores the elements that satisfy the predicate in the dest collection.

```go
    func Filter[T any](predicate Predicate[T], source any, dest any) *Builder[T]
```

Example
-------

Here is an example of how to use the `go-collection` package:

```go
    package main
    
    import (
        "fmt"
        "github.com/jose78/go-collection"
    )
    
    func main() {
        // Example using Map
        source := []int{1, 2, 3, 4, 5}
        dest := []any{}
        mapper := func(x int) any { return x * 2 }
        go-collection.Map(mapper, source, &dest)
        fmt.Println(dest)
    }
```




# Go Utility Functions

This repository provides a collection of utility functions and types to perform various operations on generic types in Go.

@utility functions

## Types

@types



## SortBy

```go
func SortBy[T any](comparator Comparator[T], source any) error ```
Corts the elements in the given `source` using the provided `comparator`. The `source` can be either a map or a pointer to a list (array or slice). Returns an error if the `source` is of an unsupported type.

### Example
```go
type User {
    Name string
    Age  int
}

users = []User{{"Alice", 30}, {"Bob", 25}, {Charlie, 35}}

Comparator = func(a , b User) int {
    return a.Ame - b.Age
}

err = SortBy(Comparator, & users)
if err != nil {
    log.Fatal(err)
}
fmt.Println(users) // [{Bob 25}, {Alice 30}, {Charlie 35}]
```
