package set

import (
	"testing"
)

// TestAdd tests the Add method of the Set.
func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1, but it does not")
	}
}

// TestRemove tests the Remove method of the Set.
func TestRemove(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Expected set to not contain 1, but it does")
	}
}

// TestContains tests the Contains method of the Set.
func TestContains(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1, but it does not")
	}
	if s.Contains(2) {
		t.Errorf("Expected set to not contain 2, but it does")
	}
}

// TestSize tests the Size method of the Set.
func TestSize(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	if size := s.Size(); size != 2 {
		t.Errorf("Expected set size to be 2, but got %d", size)
	}
}

// TestUnion tests the Union method of the Set.
func TestUnion(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)

	unionSet := s1.Union(s2)
	if unionSet.Size() != 3 {
		t.Errorf("Expected union size to be 3, but got %d", unionSet.Size())
	}
	if !unionSet.Contains(1) || !unionSet.Contains(2) || !unionSet.Contains(3) {
		t.Errorf("Union set is missing expected elements")
	}
}

// TestIntersection tests the Intersection method of the Set.
func TestIntersection(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)

	intersectionSet := s1.Intersection(s2)
	if intersectionSet.Size() != 1 {
		t.Errorf("Expected intersection size to be 1, but got %d", intersectionSet.Size())
	}
	if !intersectionSet.Contains(2) {
		t.Errorf("Intersection set is missing element 2")
	}
}

// TestDifference tests the Difference method of the Set.
func TestDifference(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)

	differenceSet := s1.Difference(s2)
	if differenceSet.Size() != 1 {
		t.Errorf("Expected difference size to be 1, but got %d", differenceSet.Size())
	}
	if !differenceSet.Contains(1) {
		t.Errorf("Difference set is missing element 1")
	}
}

// TestSymmetricDifference tests the SymmetricDifference method of the Set.
func TestSymmetricDifference(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)

	symmetricDifferenceSet := s1.SymmetricDifference(s2)
	if symmetricDifferenceSet.Size() != 2 {
		t.Errorf("Expected symmetric difference size to be 2, but got %d", symmetricDifferenceSet.Size())
	}
	if !symmetricDifferenceSet.Contains(1) || !symmetricDifferenceSet.Contains(3) {
		t.Errorf("Symmetric difference set is missing expected elements")
	}
}

// TestIsSubset tests the IsSubset method of the Set.
func TestIsSubset(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	if !s1.IsSubset(s2) {
		t.Errorf("Expected s1 to be a subset of s2")
	}
}

// TestIsSuperset tests the IsSuperset method of the Set.
func TestIsSuperset(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s2 := New[int]()
	s2.Add(1)

	if !s1.IsSuperset(s2) {
		t.Errorf("Expected s1 to be a superset of s2")
	}
}

// TestClear tests the Clear method of the Set.
func TestClear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected set size to be 0 after clear, but got %d", s.Size())
	}
}

// TestForEach tests the ForEach method of the Set.
func TestForEach(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	count := 0
	s.ForEach(func(elem int) {
		count++
	})
	if count != 2 {
		t.Errorf("Expected ForEach to be called twice, but it was called %d times", count)
	}
}

// TestMap tests the Map method of the Set.
func TestMap(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	mappedSet := s.Map(func(elem int) int {
		return elem * 2
	})
	if mappedSet.Size() != 2 || !mappedSet.Contains(2) || !mappedSet.Contains(4) {
		t.Errorf("Mapped set does not contain expected elements")
	}
}

// TestFilter tests the Filter method of the Set.
func TestFilter(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	filteredSet := s.Filter(func(elem int) bool {
		return elem%2 == 0
	})
	if filteredSet.Size() != 1 || !filteredSet.Contains(2) {
		t.Errorf("Filtered set does not contain expected elements")
	}
}

// TestClone tests the Clone method of the Set.
func TestClone(t *testing.T) {
	s := New[int]()
	s.Add(1)
	cloneSet := s.Clone()
	if !cloneSet.Contains(1) {
		t.Errorf("Clone set does not contain expected element 1")
	}
}

// TestPop tests the Pop method of the Set.
func TestPop(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	elem, ok := s.Pop()
	if !ok {
		t.Errorf("Expected Pop to return an element, but it did not")
	}
	if !(elem == 1 || elem == 2) {
		t.Errorf("Expected Pop to return 1 or 2, but got %d", elem)
	}
	if s.Size() != 1 {
		t.Errorf("Expected set size to be 1 after Pop, but got %d", s.Size())
	}
}

// TestEqual tests the Equal method of the Set.
func TestEqual(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s2 := New[int]()
	s2.Add(1)
	if !s1.Equal(s2) {
		t.Errorf("Expected sets to be equal")
	}

	s2.Add(2)
	if s1.Equal(s2) {
		t.Errorf("Expected sets to not be equal")
	}
}
