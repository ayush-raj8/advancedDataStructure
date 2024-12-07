package set

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1, but it does not")
	}
}

func TestRemove(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Expected set to not contain 1, but it does")
	}
}

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

func TestSize(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	if size := s.Size(); size != 2 {
		t.Errorf("Expected set size to be 2, but got %d", size)
	}
}

func TestSetCount(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)

	isEven := func(n int) bool {
		return n%2 == 0
	}

	count := s.Count(isEven)

	expectedCount := 2
	if count != expectedCount {
		t.Fatalf("Expected count: %d, got: %d", expectedCount, count)
	}
}

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

// Tests SymmetricDifference method of the Set.
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

// Tests the IsSubset method of the Set.
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

	if s2.IsSubset(s1) {
		t.Errorf("Expected s2 to be a subset of s1")
	}
}

func TestReverse(t *testing.T) {

	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := s1.Reverse()
	s3 := s2.Reverse()

	// Verify the result
	if !reflect.DeepEqual(s3, s1) {
		t.Fatalf("Expected reversed order: %v, got: %v", s3, s1)
	}
}

func TestIsDisjoint(t *testing.T) {

	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := New[string]()
	s2.Add("Ab1")
	s2.Add("Bc2")
	s2.Add("Cd3")

	if !s1.IsDisjoint(s2) {
		t.Fatalf("Expected set %v and set %v to be disjoint", s1, s2)
	}

	s3 := New[int]()
	s3.Add(11)
	s3.Add(22)
	s3.Add(23)

	if !s1.IsDisjoint(s3) {
		t.Fatalf("Expected set %v and set %v to be disjoint", s1, s2)
	}
}

func TestDifferenceCount(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2 := New[int]()
	set2.Add(2)
	set2.Add(4)

	diff := set1.DifferenceCount(set2)
	expected := 2
	if diff != expected {
		t.Fatalf("Expected %d, got: %d", expected, diff)
	}

	set3 := New[int]()
	diff = set1.DifferenceCount(set3)
	expected = 3
	if diff != expected {
		t.Fatalf("Expected %d, got: %d", expected, diff)
	}

	// Test with a mismatched type (string set)
	set4 := New[string]()
	set4.Add("a")
	set4.Add("b")

	diff = set1.DifferenceCount(set4)
	expected = 3
	if diff != expected {
		t.Fatalf("Expected %d, got: %d", expected, diff)
	}

	set5 := New[int]()
	set5.Add(1)
	set5.Add(2)
	set5.Add(3)

	set6 := New[int]()
	set6.Add(1)
	set6.Add(2)
	set6.Add(3)
	// all same
	diff = set5.DifferenceCount(set6)
	expected = 0
	if diff != expected {
		t.Fatalf("Expected %d, got: %d", expected, diff)
	}
	// Both empty
	set7 := New[int]()
	set8 := New[int]()

	diff = set7.DifferenceCount(set8)
	expected = 0
	if diff != expected {
		t.Fatalf("Expected %d, got: %d", expected, diff)
	}
}

// Tests the IsSuperset method of the Set.
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

// Tests the Clear method of the Set.
func TestClear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("Expected set size to be 0 after clear, but got %d", s.Size())
	}
}

// Tests the ForEach method of the Set.
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

// Tests the Map method of the Set.
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

// Tests the Filter method of the Set.
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

// Tests the Clone method of the Set.
func TestClone(t *testing.T) {
	s := New[int]()
	s.Add(1)
	cloneSet := s.Clone()
	if !cloneSet.Contains(1) {
		t.Errorf("Clone set does not contain expected element 1")
	}
}

// Tests the Pop method of the Set.
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

// Tests the Equal method of the Set.
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

func TestIterator(t *testing.T) {
	s1 := New[int]()
	s1.Add(3)
	s1.Add(1)
	s1.Add(2)

	s2 := New[int]()
	for v := range s1.Iterator() {
		s2.Add(v)
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Fatalf("Expected result %v, got: %v", s1, s2)
	}
}
