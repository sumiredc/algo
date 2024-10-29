package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type Case struct {
		length   uint16
		search   uint16
		expected string
	}

	cases := []Case{
		{length: 10, search: 5, expected: "5"},
		{length: 3, search: 5, expected: ""},
		{length: 9999, search: 7456, expected: "7456"},
		{length: 9, search: 5, expected: "5"},
		{length: 9, search: 4, expected: "4"},
		{length: 9, search: 0, expected: "0"},
		{length: 9, search: 8, expected: "8"},
	}

	for _, c := range cases {
		bs := NewBinarySearch(c.length)
		actual := bs.Search(c.search)

		if c.expected != actual {
			t.Errorf(
				"Failed test case [length: %d, search: %d, expected: %s, actual: %s]",
				c.length,
				c.search,
				c.expected,
				actual,
			)
		}
	}
}
