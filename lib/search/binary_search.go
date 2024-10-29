package search

import (
	"strconv"
)

type Row struct {
	Key   uint16
	Value string
}

type BinarySearch struct {
	data []Row
}

func NewBinarySearch(length uint16) BinarySearch {
	var i uint16
	var data []Row

	for i = 0; i < length; i++ {
		data = append(data, Row{Key: i, Value: strconv.Itoa(int(i))})
	}

	return BinarySearch{data}
}

func (bs *BinarySearch) Search(k uint16) string {
	low := 0
	high := len(bs.data) - 1

	for low <= high {
		middle := (low + high) / 2
		row := bs.data[middle]

		if k == row.Key {
			return row.Value
		} else if k < row.Key {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return ""
}
