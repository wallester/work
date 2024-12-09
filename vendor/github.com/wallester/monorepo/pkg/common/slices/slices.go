package slices

// slices provides a set of functions for working with slices.
//
// The functions use following naming conventions:
//
//	T - type of slice elements
//	OT - type of slice elements after applying the given function
//	K - type of map keys
//	V - type of map values

import (
	"log"

	"golang.org/x/exp/constraints"
)

// Select returns a new slice containing all elements for which the given
// function returns true.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := Select(people, func(p Person) bool {
//	  return p.Age == 20
//	})
//
//	// result:
//	// []Person{
//	//   {"Alice", 20},
//	//   {"Charlie", 20},
//	// }
func Select[T any](list []T, predicate func(T) bool) []T {
	if list == nil {
		return nil
	}

	result := make([]T, 0, len(list))
	for _, t := range list {
		if predicate(t) {
			result = append(result, t)
		}
	}

	return result
}

// SelectCollect combines Select and Collect. It returns a new slice containing the
// results of applying the given function to each element for which the given
// function returns true.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := SelectCollect(people, func(p Person) (string, bool) {
//	  if p.Age == 20 {
//	    return p.Name, true
//	  }
//
//	  return "", false
//	})
//
//	// result:
//	// []string{
//	//   "Alice",
//	//   "Charlie",
//	// }
//
//	result := SelectCollect(people, func(p Person) (string, bool) {
//	  return p.Name, p.Age > 20
//	}
//
//	// result:
//	// []string{
//	//   "Bob",
//	// }
func SelectCollect[T, OT any](list []T, filter func(T) (OT, bool)) []OT {
	if list == nil {
		return nil
	}

	result := make([]OT, 0, len(list))
	for _, t := range list {
		ot, ok := filter(t)
		if ok {
			result = append(result, ot)
		}
	}

	return result
}

// SelectNotNilDeref returns a new slice containing the results of applying the given
// function to each element. The order of elements is preserved.
// If the result of applying the given function is nil, the element is not included in the resulting slice.
// Each element in the resulting slice is dereferenced.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  *int
//	}
//
//	people := []Person{
//	  {"Alice", nil},
//	  {"Charlie", 20},
//	}
//
//	result := SelectNotNilDeref(people, func(p Person) *int {
//	  return p.Age
//	})
//
//	// result:
//	// []int{
//	//   20,
//	// }
func SelectNotNilDeref[T, OT any](list []T, collect func(T) *OT) []OT {
	if list == nil {
		return nil
	}

	result := make([]OT, 0, len(list))
	for _, t := range list {
		if ot := collect(t); ot != nil {
			result = append(result, *ot)
		}
	}

	return result
}

// SelectNotNil works like SelectNotNilDeref, but does not dereference the result.
func SelectNotNil[T, OT any](list []T, collect func(T) *OT) []*OT {
	if list == nil {
		return nil
	}

	result := make([]*OT, 0, len(list))
	for _, t := range list {
		if ot := collect(t); ot != nil {
			result = append(result, ot)
		}
	}

	return result
}

// Contains returns true if the given function returns true for any element in
// the slice.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := Contains(people, func(p Person) bool {
//	  return p.Name == "Bob"
//	})
//
//	// result:
//	// true
func Contains[T any](list []T, predicate func(T) bool) bool {
	for _, t := range list {
		if predicate(t) {
			return true
		}
	}

	return false
}

// All returns true if the given function returns true for all elements in the
// slice. If the slice is empty, true is returned.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := All(people, func(p Person) bool {
//	  return p.Age == 20
//	})
//
//	// result:
//	// false
//
//	result := All(people, func(p Person) bool {
//	  return p.Age >= 20
//	})
//
//	// result:
//	// true
func All[T any](list []T, predicate func(T) bool) bool {
	for _, t := range list {
		if !predicate(t) {
			return false
		}
	}

	return true
}

// Reject returns a new slice containing all elements for which the given
// function returns false.
//
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := Reject(people, func(p Person) bool {
//	  return p.Age == 20
//	})
//
//	// result:
//	// []Person{
//	//   {"Bob", 30},
//	// }
func Reject[T any](list []T, predicate func(T) bool) []T {
	if list == nil {
		return nil
	}

	result := make([]T, 0, len(list))
	for _, t := range list {
		if !predicate(t) {
			result = append(result, t)
		}
	}

	return result
}

// Collect returns a new slice containing the results of applying the given function
// to each element. The order of elements is preserved.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := Collect(people, func(p Person) string {
//	  return p.Name
//	})
//
//	// result:
//	// []string{
//	//   "Alice",
//	//   "Bob",
//	//   "Charlie",
//	// }
func Collect[T, OT any](list []T, collect func(T) OT) []OT {
	if list == nil {
		return nil
	}

	result := make([]OT, 0, len(list))
	for _, t := range list {
		result = append(result, collect(t))
	}

	return result
}

// Map function is an alias for Collect function.
// The name Map is a common name for such functionality.
func Map[T any, OT any](slice []T, mapper func(T) OT) []OT {
	return Collect(slice, func(t T) OT { return mapper(t) })
}

// Unique returns a new slice containing only the unique elements from the given
// slice. The order of elements is preserved.
// Example:
//
//	result := Unique([]int{1, 2, 3, 1})
//
//	// result:
//	// []int{
//	//   1,
//	//   2,
//	//   3,
//	// }
func Unique[T comparable](list []T) []T {
	if list == nil {
		return nil
	}

	result := make([]T, 0, len(list))
	seen := make(map[T]bool, len(list))
	for _, t := range list {
		if !seen[t] {
			seen[t] = true
			result = append(result, t)
		}
	}

	return result
}

// SliceOfAny returns a new slice containing the same elements as the given
// slice, but with the type of each element set to any.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := SliceOfAny(people)
//
//	// result:
//	// []any{
//	//   Person{"Alice", 20},
//	//   Person{"Bob", 30},
//	//   Person{"Charlie", 20},
//	// }
func SliceOfAny[T any](list []T) []any {
	if list == nil {
		return nil
	}

	result := make([]any, len(list))
	for i, t := range list {
		result[i] = t
	}

	return result
}

// BatchesOfAny returns batches of any. batchSize must be greater than 0,
// otherwise it will be set to the length of the given slice.
// Example:
//
//	result := BatchesOfAny([]int{1, 2, 3, 4, 5}, 2)
//
//	// result:
//	// [][]any{
//	//   []any{1, 2},
//	//   []any{3, 4},
//	//   []any{5},
//	// }
//
//	result := BatchesOfAny([]int{1, 2, 3, 4, 5}, 3)
//
//	// result:
//	// [][]any{
//	//   []any{1, 2, 3},
//	//   []any{4, 5},
//	// }
//
//	result := BatchesOfAny([]int{1, 2, 3, 4, 5}, 0)
//
//	// result:
//	// [][]any{
//	//   []any{1, 2, 3, 4, 5},
//	// }
func BatchesOfAny[T any](list []T, batchSize int) [][]any {
	return Batches(SliceOfAny(list), batchSize)
}

// Batches returns batches of the given slice. batchSize must be greater than 0,
// otherwise it will be set to the length of the given slice.
// Example:
//
//	people := []string{"Alice", "Bob", "Charlie"}
//
//	result := Batches(people, 2)
//
//	// result:
//	// [][]string{
//	//   []string{"Alice", "Bob"},
//	//   []string{"Charlie"},
//	// }
func Batches[T any](list []T, batchSize int) [][]T {
	if list == nil {
		return nil
	}

	if batchSize <= 0 {
		batchSize = len(list)
	}

	var batches [][]T
	for i := 0; i < len(list); i += batchSize {
		end := i + batchSize
		if end > len(list) {
			end = len(list)
		}

		batches = append(batches, list[i:end])
	}

	return batches
}

// Find returns first element from slice for which the given function returns
// true. If no such element is found, the zero value of the slice's element type
// is returned. The second return value indicates whether a matching element was
// found.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result, ok := Find(people, func(p Person) bool {
//	  return p.Name == "Bob"
//	})
//
//	// result:
//	// Person{"Bob", 30}, true
//
//	result, ok := Find(people, func(p Person) bool {
//	  return p.Name == "Dave"
//	})
//
//	// result:
//	// Person{}, false
func Find[T any](list []T, predicate func(T) bool) (T, bool) {
	for _, t := range list {
		if predicate(t) {
			return t, true
		}
	}

	var zero T
	return zero, false
}

// FindFirst returns first element from slice for which the given function returns
// true. If no such element is found, the zero value of the slice's element type
// is returned.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result, ok := FindFirst(people, func(p Person) bool {
//	  return p.Name == "Bob"
//	})
//
//	// result:
//	// Person{"Bob", 30}
//
//	result, ok := FindFirst(people, func(p Person) bool {
//	  return p.Name == "Dave"
//	})
//
//	// result:
//	// Person{}
func FindFirst[T any](list []T, predicate func(T) bool) T {
	res, _ := Find(list, predicate)
	return res
}

// Group returns a map of slices, grouped by a key.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := Group(people, func(p Person) int {
//	  return p.Age
//	})
//
//	// result:
//	// map[int][]Person{
//	//   20: []Person{
//	//     {"Alice", 20},
//	//     {"Charlie", 20},
//	//   },
//	//   30: []Person{
//	//     {"Bob", 30},
//	//   },
//	// }
func Group[K comparable, T any](list []T, key func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, t := range list {
		k := key(t)
		result[k] = append(result[k], t)
	}

	return result
}

// GroupCollect combines Group and Collect. It returns a map of slices, grouped by a
// key. The values of the map are the results of applying the given function to
// each element.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	  {"Charlie", 20},
//	}
//
//	result := GroupCollect(people, func(p Person) (int, string) {
//	  return p.Age, p.Name
//	})
//
//	// result:
//	// map[int][]string{
//	//   20: []string{
//	//     "Alice",
//	//     "Charlie",
//	//   },
//	//   30: []string{
//	//     "Bob",
//	//   },
//	// }
func GroupCollect[K comparable, T, OT any](list []T, mapper func(T) (K, OT)) map[K][]OT {
	result := make(map[K][]OT)
	for _, t := range list {
		k, ot := mapper(t)
		result[k] = append(result[k], ot)
	}

	return result
}

// Apply applies the given methods to the target object. The methods are applied
// in the order they are given. The target object is returned
// for convenience. This allows for method chaining.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	var p Person
//
//	Apply(&p,
//	  func(p *Person) {
//	    p.Name = "Alice"
//	  },
//	  func(p *Person) {
//	    p.Age = 20
//	  },
//	)
//
//	// p:
//	// Person{
//	//   Name: "Alice",
//	//   Age:  20,
//	// }
func Apply[T any](target T, methods ...func(T)) T {
	for _, method := range methods {
		if method != nil {
			method(target)
		}
	}

	return target
}

// Sum returns the sum of all elements in the slice. If the slice is empty, zero is
// returned.
// Example:
//
//	result := Sum([]int{1, 2, 3, 4, 5}, func(n int) int { return n })
//
//	// result:
//	// 15
//
//	result := Sum([]float64{1.1, 2.2}, func(f float64) float64 { return f })
//
//	// result:
//	// 3.3
func Sum[OT Number, T any](list []T, collect func(T) OT) OT {
	var sum OT
	for _, t := range list {
		sum += collect(t)
	}

	return sum
}

type Number interface {
	constraints.Integer | constraints.Float
}

// ToMap converts a slice to a map. The key and value types are determined by
// the given converter function.
// Example:
//
//	type Person struct {
//	  Name string
//	  Age  int
//	}
//
//	people := []Person{
//	  {"Alice", 20},
//	  {"Bob", 30},
//	}
//
//	result := ToMap(people, func(p Person) (string, Person) {
//	  return p.Name, p
//	})
//
//	// result:
//	// map[string]Person{
//	//   "Alice": Person{"Alice", 20},
//	//   "Bob":   Person{"Bob", 30},
//	// }
//
//	result := ToMap(people, func(p Person) (int, string) {
//	  return p.Age, p.Name
//	})
//
//	// result:
//	// map[int]string{
//	//   20: "Alice",
//	//   30: "Bob",
//	// }
func ToMap[T any, K comparable, V any](list []T, mapper func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(list))
	for _, t := range list {
		k, v := mapper(t)
		result[k] = v
	}

	return result
}

// ToBoolMap converts a slice into a map where each element of the slice
// is a key in the map, and the value is always true. This can be useful
// for quickly checking the existence of an element in the slice.
//
// Parameters:
//
//	list - A slice of elements of any comparable type.
//
// Returns:
//
//	A map where each key is an element from the slice and the value is true.
func ToBoolMap[T comparable](list []T) map[T]bool {
	return ToMap(list, func(t T) (T, bool) { return t, true })
}

// Each applies the given function to each element in the slice.
// Example:
//
//	Each([]int{1, 2, 3}, func(n int) {
//	  fmt.Println(n)
//	})
//
//	// Output:
//	// 1
//	// 2
//	// 3
func Each[T any](list []T, visit func(T)) {
	for _, t := range list {
		visit(t)
	}
}

// Shift removes the first element from the given slice and returns it. If the
// slice is empty, nil is returned. The second return value is the slice without
// the first element.
// Example:
//
//	result, list := Shift([]int{1, 2, 3})
//
//	// result:
//	// 1
//
//	// list:
//	// []int{
//	//   2,
//	//   3,
//	// }
func Shift[T any](list []T) (T, []T) {
	if len(list) == 0 {
		var zero T
		return zero, nil
	}

	return list[0], list[1:]
}

// Pop removes the last element from the given slice and returns it. If the slice
// is empty, nil is returned. The second return value is the slice without the
// last element.
// Example:
//
//	result, list := Pop([]int{1, 2, 3})
//
//	// result:
//	// 3
//
//	// list:
//	// []int{
//	//   1,
//	//   2,
//	// }
func Pop[T any](list []T) (T, []T) {
	if len(list) == 0 {
		var zero T
		return zero, nil
	}

	return list[len(list)-1], list[:len(list)-1]
}

// Reverse returns a new slice with the elements in reverse order.
// Example:
//
//	result := Reverse([]int{1, 2, 3})
//
//	// result:
//	// []int{
//	//   3,
//	//   2,
//	//   1,
//	// }
func Reverse[T any](list []T) []T {
	if list == nil {
		return nil
	}

	result := make([]T, len(list))
	for i, j := 0, len(list)-1; i < len(list); i, j = i+1, j-1 {
		result[i] = list[j]
	}

	return result
}

// Prepend returns a new slice with the given element added at the beginning.
// The slice where the element is added is the first argument, similar to the
// append function.
// Example:
//
//	result := Prepend([]int{2, 3}, 1,
//
//	// result:
//	// []int{
//	//   1,
//	//   2,
//	//   3,
//	// }
func Prepend[T any](list []T, elements ...T) []T {
	return append(elements, list...)
}

// SelectByType returns a new slice containing all elements of the given type.
// Example:
//
//	result := SelectByType([]any{1, "two", 3.0}, int)
//
//	// result:
//	// []int{
//	//   1,
//	// }
//
//	result := SelectByType([]any{1, "two", 3.0}, string)
//
//	// result:
//	// []string{
//	//   "two",
//	// }
func SelectByType[T any](list []any) []T {
	return SelectCollect(list, func(v any) (T, bool) {
		value, ok := v.(T)
		return value, ok
	})
}

// RequireUnique panics if the given slice contains duplicate elements.
// It is meant to check for programming errors and should not be used for
// validation of user input!
func RequireUnique[T comparable](list []T) []T {
	seen := make(map[T]bool, len(list))
	for _, t := range list {
		if seen[t] {
			log.Panicf("duplicate element found: %v", t)
		}

		seen[t] = true
	}

	return list
}
