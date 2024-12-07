package list

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type List struct {
	data []any
}

func New() *List {
	return &List{
		data: []any{},
	}
}

func (l *List) Append(element any) {
	l.data = append(l.data, element)
}

func (l *List) Extend(elements []any) {
	l.data = append(l.data, elements...)
}

func (l *List) Insert(index int, element any) error {
	if index < 0 || index > len(l.data) {
		return errors.New("index out of bounds")
	}
	l.data = append(l.data[:index], append([]any{element}, l.data[index:]...)...)
	return nil
}

func (l *List) Remove(element any) error {
	for i, v := range l.data {
		if reflect.DeepEqual(v, element) {
			l.data = append(l.data[:i], l.data[i+1:]...)
			return nil
		}
	}
	return errors.New("element not found")
}

func (l *List) Pop(index int) (any, error) {
	if index < 0 || index >= len(l.data) {
		return nil, errors.New("index out of bounds")
	}
	element := l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)
	return element, nil
}

func (l *List) Clear() {
	l.data = []any{}
}

func (l *List) Len() int {
	return len(l.data)
}

func (l *List) Reverse() {
	for i, j := 0, len(l.data)-1; i < j; i, j = i+1, j-1 {
		l.data[i], l.data[j] = l.data[j], l.data[i]
	}
}

func (l *List) Slice(params ...int) (*List, error) {
	start, end, step := 0, len(l.data), 1
	switch len(params) {
	case 1:
		start = params[0]
	case 2:
		start, end = params[0], params[1]
	case 3:
		start, end, step = params[0], params[1], params[2]
	}

	if step == 0 {
		return nil, errors.New("step cannot be zero")
	}

	// Handle negative indices
	n := len(l.data)
	if start < 0 {
		start += n
	}
	if end < 0 {
		end += n
	}

	// Clamp indices to valid range
	if start < 0 {
		start = 0
	}
	if end > n {
		end = n
	}
	if start >= n || end < 0 || start >= end {
		return &List{}, nil // Return an empty list for invalid ranges
	}

	// Create the sliced list
	slicedData := []any{}
	if step > 0 {
		for i := start; i < end; i += step {
			slicedData = append(slicedData, l.data[i])
		}
	} else {
		for i := start; i > end; i += step {
			slicedData = append(slicedData, l.data[i])
		}
	}

	return &List{data: slicedData}, nil
}

func (l *List) Sort(order ...string) error {
	if len(l.data) == 0 {
		return nil
	}

	if !l.isHomogeneous() {
		return errors.New("cannot sort a list with mixed data types")
	}

	isAscending := true
	if len(order) > 0 && order[0] == "desc" {
		isAscending = false
	}
	fmt.Println(l.data[0])

	switch l.data[0].(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, bool:
		sort.Slice(l.data, func(i, j int) bool {
			if isAscending {
				return compareNumbers(l.data[i], l.data[j]) < 0
			}
			return compareNumbers(l.data[i], l.data[j]) > 0
		})
	case string:
		sort.Slice(l.data, func(i, j int) bool {
			if isAscending {
				return l.data[i].(string) < l.data[j].(string)
			}
			return l.data[i].(string) > l.data[j].(string)
		})
	default:
		return errors.New("unsupported type for sorting")
	}
	return nil
}

func compareNumbers(a, b any) int {
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	if aVal.Kind() == reflect.Bool {
		if aVal.Bool() {
			a = 1
		} else {
			a = 0
		}
	}
	if bVal.Kind() == reflect.Bool {
		if bVal.Bool() {
			b = 1
		} else {
			b = 0
		}
	}

	var aValue = reflect.ValueOf(a).Convert(reflect.TypeOf(float64(0))).Float()
	var bValue = reflect.ValueOf(b).Convert(reflect.TypeOf(float64(0))).Float()

	if aValue < bValue {
		return -1
	}
	if aValue > bValue {
		return 1
	}
	return 0
}

func (l *List) Iterator() <-chan any {
	ch := make(chan any)
	go func() {
		for _, v := range l.data {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func (l *List) isHomogeneous() bool {
	if len(l.data) == 0 {
		return true
	}

	// Get the type of the first item
	firstType := reflect.TypeOf(l.data[0])
	isNumber := isNumericBoolType(firstType)
	isString := firstType.Kind() == reflect.String

	for _, v := range l.data {
		currentType := reflect.TypeOf(v)

		// Check if the type is consistent with the first element
		if isNumber && !isNumericBoolType(currentType) {
			return false
		}
		if isString && currentType.Kind() != reflect.String {
			return false
		}
	}
	return true
}

func isNumericBoolType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool:
		return true
	default:
		return false
	}
}

func (l *List) Display() {
	fmt.Println(l.data)
}
