package data_structure

import "testing"

type PushCase struct {
	Values   []int
	Expected int
}

func TestPush(t *testing.T) {
	cases := []PushCase{
		{Values: []int{1, 1, 2, 2, 3}, Expected: 5},
		{Values: []int{}, Expected: 0},
		{Values: []int{5, 4, 3, 2, 1, 0}, Expected: 6},
	}

	for _, c := range cases {
		stack := NewStack()
		for _, v := range c.Values {
			stack.Push(v)
		}

		actual := stack.Len()
		if c.Expected != actual {
			t.Errorf(
				"Failed to Push to the Stack [values: %v, expected: %d, actual: %d]",
				c.Values,
				c.Expected,
				actual,
			)
		}
	}

}

type PopCase struct {
	Values []int
}

func TestPop(t *testing.T) {

	cases := []PopCase{
		{Values: []int{1, 1, 2, 2, 3}},
		{Values: []int{100}},
		{Values: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}

	for _, c := range cases {
		stack := NewStack()
		for _, v := range c.Values {
			stack.Push(v)
		}

		// Validate the value returned by Pop
		for i := len(c.Values) - 1; 0 <= i; i-- {
			expected := c.Values[i]
			actual := stack.Pop()

			if actual == nil {
				t.Errorf(
					"Failed to Pop to the Stack [values: %v, expected: %d, actual: nil]",
					c.Values,
					expected,
				)
			}

			if expected != *actual {
				t.Errorf(
					"Failed to Pop to the Stack [values: %v, expected: %d, actual: %d]",
					c.Values,
					expected,
					*actual,
				)
			}
		}

		// Validate the value returned by Pop is nil
		actual := stack.Pop()
		if nil != actual {
			t.Errorf(
				"Failed to last Pop to the Stack [values: %v, expected: nil, actual: %d]",
				c.Values,
				*actual,
			)
		}
	}

}

func TestEmpty(t *testing.T) {
	stack := NewStack()

	if !stack.Empty() {
		t.Error("Failed to Empty to the initial Stack")
	}

	values := []int{1, 1, 1}
	for _, v := range values {
		stack.Push(v)
	}

	count := len(values)
	for count > 0 {
		count--
		stack.Pop()
	}

	if !stack.Empty() {
		t.Error("Failed to Empty to the Stack")
	}
}

func TestInit(t *testing.T) {
	stack := NewStack()

	values := []int{1, 1, 1}
	for _, v := range values {
		stack.Push(v)
	}

	stack.Init()

	if !stack.Empty() {
		t.Error("Failed to Init to the Stack")
	}
}
