package set

import (
	"fmt"
	"math"
	"strings"
)

// Set is an unordered collection of distinct comparable items.
// Common uses include membership testing, removing duplicates from sequence,
// and doing boolean set operations like Intersection and Union.
type Set[T comparable] map[T]struct{}

// New creates a new set with 0 or many items.
func New[T comparable](items ...T) Set[T] {
	set := make(Set[T], len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

// Add includes the specified items (one or many) to the Set s.
// Set s is modified in place. If passed nothing it silently returns.
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

// Remove deletes the specified items (one or many) from the Set s.
// Set s is modified in place. If passed nothing it silently returns.
func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

// Equals tests whether s and other are the same size and contain the same items.
func (s Set[T]) Equals(other Set[T]) bool {
	if other == nil || len(s) != len(other) {
		return false
	}

	for key := range s {
		_, ok := other[key]
		if !ok {
			return false
		}
	}

	return true
}

// Contains tests if Set s contains all items passed.
// It returns false if nothing is passed.
func (s Set[T]) Contains(items ...T) bool {
	if len(items) == 0 {
		return false
	}

	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}

	return true
}

// ContainsAny tests if Set s contains any of the items passed.
// It returns false of nothing is passed.
func (s Set[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}

	return false
}

// IsSubSet tests if every element of s exists in the other.
func (s Set[T]) IsSubSet(other Set[T]) bool {
	for k := range s {
		if _, ok := other[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperSet tests if every element of other exists in s.
func (s Set[T]) IsSuperSet(other Set[T]) bool {
	return other.IsSubSet(s)
}

// Copy return a new Set with a copy of s.
func (s Set[T]) Copy() Set[T] {
	result := make(Set[T], len(s))

	for item := range s {
		result[item] = struct{}{}
	}

	return result
}

// String returns a string representation of s
func (s Set[T]) String() string {
	strs := make([]string, 0, len(s))
	for k := range s {
		strs = append(strs, fmt.Sprintf("%v", k))
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

// Slice returns a Slice of all items in Set s as a []T.
func (s Set[T]) Slice() []T {
	result := make([]T, 0, len(s))

	for item := range s {
		result = append(result, item)
	}

	return result
}

// Merge adds all items from Set other into Set s.
// This works just like Union, but it applies to Set s in place.
func (s Set[T]) Merge(other Set[T]) {
	if other == nil {
		return
	}

	for item := range other {
		s[item] = struct{}{}
	}
}

// Separate removes items in Set other from Set s.
// This works just like Difference but applies to the Set in place.
func (s Set[T]) Separate(other Set[T]) {
	for item := range other {
		s.Remove(item)
	}
}

// Union is the merger of multiple sets.
// It returns a new Set with all elements in all the sets passed.
func Union[T comparable](sets ...Set[T]) Set[T] {
	combinedLen := 0
	for _, set := range sets {
		if set == nil {
			continue
		}
		combinedLen += len(set)
	}

	if combinedLen == 0 {
		return New[T]()
	}

	result := make(Set[T], combinedLen)
	for _, set := range sets {
		result.Merge(set)
	}

	return result
}

// Difference returns a new set which contains items that are in the first
// Set but not in the others.
func Difference[T comparable](lhs Set[T], others ...Set[T]) Set[T] {
	s := lhs.Copy()
	for _, set := range others {
		s.Separate(set)
	}
	return s
}

// Intersection returns a new Set which contains items that only exist
// in all given sets.
func Intersection[T comparable](sets ...Set[T]) Set[T] {
	minIndex := -1
	minLength := math.MaxInt
	for i, set := range sets {
		if l := len(set); l < minLength {
			minLength = l
			minIndex = i
		}
	}

	if minLength == math.MaxInt || minLength == 0 {
		return New[T]()
	}

	result := sets[minIndex].Copy()
	for i, set := range sets {
		if i == minIndex {
			continue
		}

		for item := range result {
			if _, ok := set[item]; !ok {
				delete(result, item)
			}
		}
	}
	return result
}

// SymmetricDifference returns a new set that contains elements from either
// passed set but not both.
func SymmetricDifference[T comparable](lhs, rhs Set[T]) Set[T] {
	u := Difference(lhs, rhs)
	v := Difference(rhs, lhs)
	return Union(u, v)
}
