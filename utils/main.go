package utils

import "github.com/FlamesX-128/anigo-plugins/types"

func GetKeys[T types.Slice, Y any](m map[T]Y) (k []T) {
	for e := range m {
		k = append(k, e)

	}

	return
}
