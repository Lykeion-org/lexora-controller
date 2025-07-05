package grpc

import (
	"sort"
)

func EnumNamesOrdered(enumNameMap map[int32]string) []string {
	// Collect keys
	keys := make([]int, 0, len(enumNameMap))
	for k := range enumNameMap {
		keys = append(keys, int(k))
	}

	// Sort keys
	sort.Ints(keys)

	// Collect values in order
	names := make([]string, len(keys))
	for i, k := range keys {
		names[i] = enumNameMap[int32(k)]
	}
	return names
}