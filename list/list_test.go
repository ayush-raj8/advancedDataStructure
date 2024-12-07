package list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	if l == nil {
		t.Fatal("Expected a new List, got nil")
	}
	if len(l.data) != 0 {
		t.Fatalf("Expected an empty list, got: %v", l.data)
	}
}

func TestAppend(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append("hello")
	l.Append(true)

	if len(l.data) != 3 {
		t.Fatalf("Expected list of length 2, got: %d", len(l.data))
	}

	if l.data[0] != 1 || l.data[1] != "hello" || l.data[2] != true {
		t.Fatalf("Expected [1, 'hello'], got: %v", l.data)
	}
}

func TestExtend(t *testing.T) {
	l1 := New()
	l1.Append(1)
	l1.Append(1.1)

	l2 := New()
	l2.Append(1)
	l2.Append(true)
	l2.Append("A")

	l1.Extend(l2.data)

	// Check if the length of l1 is correct
	if len(l1.data) != 5 {
		t.Fatalf("Expected list of length 5, got: %d", len(l1.data))
	}

	// Check if the values in the list are correct
	expected := []any{1, 1.1, 1, true, "A"}
	if !reflect.DeepEqual(l1.data, expected) {
		t.Fatalf("Expected list %v, got: %v", expected, l1.data)
	}
}

func TestInsert(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(3)

	err := l.Insert(1, 2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l.data, []any{1, 2, 3}) {
		t.Fatalf("Expected [1, 2, 3], got: %v", l.data)
	}

	err = l.Insert(5, 4) // Invalid index
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestRemove(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append('a')

	err := l.Remove(2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l.data, []any{1, 'a'}) {
		t.Fatalf("Expected [1, 3], got: %v", l.data)
	}

	err = l.Remove(5) // Non-existing element
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestPop(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	elem, err := l.Pop(1)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if elem != 2 {
		t.Fatalf("Expected popped element to be 2, got: %v", elem)
	}

	if !reflect.DeepEqual(l.data, []any{1, 3}) {
		t.Fatalf("Expected [1, 3], got: %v", l.data)
	}

	_, err = l.Pop(5) // Invalid index
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestClear(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)

	l.Clear()

	if len(l.data) != 0 {
		t.Fatalf("Expected empty list, got: %v", l.data)
	}
}

func TestLen(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)

	if l.Len() != 2 {
		t.Fatalf("Expected length 2, got: %d", l.Len())
	}

	l.Clear()
	if l.Len() != 0 {
		t.Fatalf("Expected length 0, got: %d", l.Len())
	}
}

func TestReverse(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	l.Reverse()

	if !reflect.DeepEqual(l.data, []any{3, 2, 1}) {
		t.Fatalf("Expected [3, 2, 1], got: %v", l.data)
	}
}

func TestSlice(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	slicedList, err := l.Slice(1, 3)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(slicedList.data, []any{2, 3}) {
		t.Fatalf("Expected [2, 3], got: %v", slicedList.data)
	}

	// Test negative index handling
	slicedList, err = l.Slice(-3, -1)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(slicedList.data, []any{2, 3}) {
		t.Fatalf("Expected [2, 3], got: %v", slicedList.data)
	}

	// Test step functionality
	slicedList, err = l.Slice(0, 4, 2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(slicedList.data, []any{1, 3}) {
		t.Fatalf("Expected [1, 3], got: %v", slicedList.data)
	}

	slicedList, err = l.Slice(0)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(slicedList.data, []any{1, 2, 3, 4}) {
		t.Fatalf("Expected [1, 2, 3, 4], got: %v", slicedList.data)
	}
	_, err = l.Slice(0, 4, 0)

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if err.Error() != "step cannot be zero" {
		t.Fatalf("Expected error 'step cannot be zero', got: %v", err)
	}

}

func TestSort(t *testing.T) {
	l := New()
	l.Append(3)
	l.Append(1)
	l.Append(2)
	l.Append(2.5)
	l.Append(true)
	l.Append(false)
	//l.Append('a')

	err := l.Sort("asc")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l.data, []any{false, 1, true, 2, 2.5, 3}) {
		t.Fatalf("Expected [false, 1, true, 2, 2.5, 3], got: %v", l.data)
	}

	err = l.Sort("desc")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l.data, []any{3, 2.5, 2, 1, true, false}) {
		t.Fatalf("Expected [3, 2.5, 2, true, 1, false], got: %v", l.data)
	}

	sliced, err := l.Slice(-1)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(sliced.data, []any{false}) {
		t.Fatalf("Expected [false], got: %v", sliced.data)
	}

	sliced, err = l.Slice(3, 0, -1)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := []any{1, 2, 2.5}
	if !reflect.DeepEqual(sliced.data, expected) {
		t.Fatalf("Expected %v, got: %v", expected, sliced.data)
	}

	l2 := New()
	l2.Append("a")
	l2.Append("c")
	l2.Append("b")
	l2.Append("1") // Note: "1" is a string, not a character

	err = l2.Sort()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l2.data, []any{"1", "a", "b", "c"}) {
		t.Fatalf("Expected [1 a b c], got: %v", l2.data)
	}

	err = l2.Sort("desc")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !reflect.DeepEqual(l2.data, []any{"c", "b", "a", "1"}) {
		t.Fatalf("Expected [c, b, a, 1], got: %v", l2.data)
	}

	l3 := New()
	l3.Append(1)
	l3.Append("ABC")
	l3.Append(3)
	err = l3.Sort()

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if err.Error() != "cannot sort a list with mixed data types" {
		t.Fatalf("Expected error 'cannot sort a list with mixed data types', got: %v", err)
	}

	l4 := New()
	err = l4.Sort()
	if err != nil {
		t.Fatalf("Expected nil, got error %v", err)
	}

	type Person struct {
		Name string
		Age  int
	}

	l5 := New()
	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Bob", Age: 30}
	l5.Append(person1)
	l5.Append(person2)
	err = l5.Sort()

	if err == nil {
		t.Fatalf("Expected eror, got nil")
	}

	if err.Error() != "unsupported type for sorting" {
		t.Fatalf("Expected error 'unsupported type for sorting', got: %v", err)
	}
}

func TestIterator(t *testing.T) {
	l := New()
	l.Append(3)
	l.Append(1)
	l.Append(2)
	l.Append(2.5)
	l.Append(true)
	l.Append(false)

	var result []any
	for v := range l.Iterator() {
		result = append(result, v)
	}

	expected := []any{3, 1, 2, 2.5, true, false}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected result %v, got: %v", expected, result)
	}
}
