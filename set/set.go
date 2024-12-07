package set

import "sync"

// Set is a thread-safe implementation of a set data structure.
type Set[T comparable] struct {
	elements map[T]struct{} // struct is 0 sized
	mu       sync.RWMutex
}

// New creates and returns a new instance of Set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]struct{}),
	}
}

// Add inserts an element into the set.
func (s *Set[T]) Add(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements[elem] = struct{}{}
}

// Remove deletes an element from the set.
func (s *Set[T]) Remove(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.elements, elem)
}

// Contains checks if an element is in the set.
func (s *Set[T]) Contains(elem T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.elements[elem]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.elements)
}

// ToSlice returns the elements of the set as a slice.
func (s *Set[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	slice := make([]T, 0, len(s.elements))
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// Union returns a new set that is the union of s and another set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for _, elem := range s.ToSlice() {
		result.Add(elem)
	}
	for _, elem := range other.ToSlice() {
		result.Add(elem)
	}
	return result
}

// Intersection returns a new set that is the intersection of s and another set.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	for _, elem := range s.ToSlice() {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new set that is the difference of s and another set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for _, elem := range s.ToSlice() {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// SymmetricDifference returns a new set with elements in either set but not in both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := New[T]()
	for _, elem := range s.ToSlice() {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	for _, elem := range other.ToSlice() {
		if !s.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// IsSubset checks if the current set is a subset of another set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for _, elem := range s.ToSlice() {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the current set is a superset of another set.
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		s.mu.RLock()
		defer s.mu.RUnlock()
		for elem := range s.elements {
			ch <- elem
		}
		close(ch)
	}()
	return ch
}

// ForEach applies the provided function to each element in the set.
func (s *Set[T]) ForEach(f func(T)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for elem := range s.elements {
		f(elem)
	}
}

// Map applies the function to each element and returns a new set with the results.
func (s *Set[T]) Map(f func(T) T) *Set[T] {
	result := New[T]()
	s.ForEach(func(elem T) {
		result.Add(f(elem))
	})
	return result
}

// Filter returns a new set with elements that satisfy the predicate function.
func (s *Set[T]) Filter(f func(T) bool) *Set[T] {
	result := New[T]()
	s.ForEach(func(elem T) {
		if f(elem) {
			result.Add(elem)
		}
	})
	return result
}

// Count returns the number of elements that satisfy the predicate.
func (s *Set[T]) Count(f func(T) bool) int {
	count := 0
	s.ForEach(func(elem T) {
		if f(elem) {
			count++
		}
	})
	return count
}

// Copy returns a new set that is a copy of the current set.
func (s *Set[T]) Copy() *Set[T] {
	result := New[T]()
	s.ForEach(func(elem T) {
		result.Add(elem)
	})
	return result
}

// Equal checks if the current set is equal to another set.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for _, elem := range s.ToSlice() {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Pop removes and returns a random element from the set.
func (s *Set[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for elem := range s.elements {
		delete(s.elements, elem)
		return elem, true
	}
	var zero T
	return zero, false
}

// Reverse returns a new set with elements in reverse order.
func (s *Set[T]) Reverse() *Set[T] {
	slice := s.ToSlice()
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	result := New[T]()
	for _, elem := range slice {
		result.Add(elem)
	}
	return result
}

// DifferenceCount returns the number of elements in the first set but not in the second.
func (s *Set[T]) DifferenceCount(other any) int {
	otherSet, ok := other.(*Set[T])
	if !ok {
		return s.Size()
	}

	count := 0
	s.ForEach(func(elem T) {
		if !otherSet.Contains(elem) {
			count++
		}
	})

	return count
}

// Clone creates a new set with the same elements but a different internal map.
func (s *Set[T]) Clone() *Set[T] {
	clone := New[T]()
	s.ForEach(func(elem T) {
		clone.Add(elem)
	})
	return clone
}

// IsDisjoint checks if the current set and the other set have no elements in common.
func (s *Set[T]) IsDisjoint(other any) bool {
	if otherSet, ok := other.(*Set[T]); ok {
		for _, elem := range s.ToSlice() {
			if otherSet.Contains(elem) {
				return false
			}
		}
		return true
	}
	return true
}
