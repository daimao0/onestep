package slice_util

import (
	"golang.org/x/exp/rand"
	"time"
)

// RandomSampling is a function that randomly selects a specified number of elements from a slice.
func RandomSampling[T any](data []T, size int) []T {
	// func border
	if size <= 0 {
		return []T{}
	}
	if size > len(data) {
		return data
	}
	//Randomly generate sampling indexes
	rand.Seed(uint64(time.Now().UnixNano()))
	indexes := make(map[int]any, size)
	for i := 0; i < size; i++ {
		index := rand.Intn(len(data))
		if indexes[index] == nil {
			indexes[index] = 1
			continue
		}
		for indexes[index] != nil {
			index++
			if index >= len(data) {
				index = 0
			}
			if indexes[index] == nil {
				indexes[index] = 1
				break
			}
		}
	}
	result := make([]T, 0, size)
	for key := range indexes {
		result = append(result, data[key])
	}
	return result
}
