[![Go Report Card](https://goreportcard.com/badge/github.com/jose78/go-collection)](https://goreportcard.com/report/github.com/jose78/go-collection)
[![Coverage Status](https://coveralls.io/repos/github/jose78/go-collection/badge.svg?branch=master)](https://coveralls.io/github/jose78/go-collection?branch=master)
[![CircleCI](https://circleci.com/gh/jose78/go-collection.svg?style=shield)](https://circleci.com/gh/jose78/go-collection)


# go-collection <img align="right" width="80" height="100" src="resources/gopher.png">




The `go-collection` package provides a set of utility functions for working with collections (slices and maps) in Go. It includes functions for mapping, filtering, zipping, and performing actions on collections.

Installation
------------

To install the package, use:

    go get github.com/jose78/go-collection

Usage
-----

### Types

#### Touple

Represents a key-value pair with a generic key and value.

    type Touple struct {
        Key   any // Key is the key of the key-value pair.
        Value any // Value is the value of the key-value pair.
    }

#### Mapper

Represents a function that takes a value of any type `T` and returns a value of any type.

    type Mapper[T any] func(T) any

#### Predicate

Represents a function that takes a value of any type `T` and returns a boolean. It is used to test whether the input value satisfies a certain condition.

    type Predicate[T any] func(T) bool

#### Action

Represents a function that takes an index and a value of any type `T`. It is intended to be used in an iteration context, such as a forEach function, where it will be executed for each element in a collection.

    type Action[T any] func(int, T)

#### KeySelector

Represents a function that takes a key of type `K` and returns a `Touple` with the key and a value of type `V`. The key must be of a comparable type.

    type KeySelector[K comparable, V any] func(K) Touple

#### Builder

A struct with an error and the item that caused the error.

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

### Functions

#### Map

Applies a `Mapper` function to each element in the source collection and stores the result in the dest collection.

    func Map[T any](mapper Mapper[T], source []T, dest *[]any) *Builder[T]

#### ForEach

Applies an `Action` function to each element in the source collection.

    func ForEach[K any](action Action[K], src any) *Builder[K]

#### Zip

Combines two slices into a map, using elements from the keys slice as keys and elements from the values slice as values.

    func Zip[K comparable, V any](keys []K, values []V, result map[K]V) *Builder[K]

#### isMap

Checks if the given element is of map type.

    func isMap(elements any) bool

#### store

Inserts data into the destination collection, which can be either a map or a slice.

    func store(data any, dest any)

#### Filter

Applies a `Predicate` function to each element in the source collection and stores the elements that satisfy the predicate in the dest collection.

    func Filter[T any](predicate Predicate[T], source any, dest any) *Builder[T]

Example
-------

Here is an example of how to use the `go-collection` package:

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