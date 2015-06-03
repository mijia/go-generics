package generics

func UniqueItems_Int64Slice(items []int64) []int64 {
	if len(items) == 0 {
		return items
	}
	visited := make(map[int64]struct{})
	uItems := make([]int64, 0, len(items))
	for _, item := range items {
		if _, ok := visited[item]; !ok {
			uItems = append(uItems, item)
			visited[item] = struct{}{}
		}
	}
	return uItems
}

func UniqueItems_StringSlice(items []string) []string {
	if len(items) == 0 {
		return items
	}
	visited := make(map[string]struct{})
	uItems := make([]string, 0, len(items))
	for _, item := range items {
		if _, ok := visited[item]; !ok {
			uItems = append(uItems, item)
			visited[item] = struct{}{}
		}
	}
	return uItems
}

func UniqueItems_IntSlice(items []int) []int {
	if len(items) == 0 {
		return items
	}
	visited := make(map[int]struct{})
	uItems := make([]int, 0, len(items))
	for _, item := range items {
		if _, ok := visited[item]; !ok {
			uItems = append(uItems, item)
			visited[item] = struct{}{}
		}
	}
	return uItems
}

func UniqueItems_Float64Slice(items []float64) []float64 {
	if len(items) == 0 {
		return items
	}
	visited := make(map[float64]struct{})
	uItems := make([]float64, 0, len(items))
	for _, item := range items {
		if _, ok := visited[item]; !ok {
			uItems = append(uItems, item)
			visited[item] = struct{}{}
		}
	}
	return uItems
}

func Equal_Int64Slice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Equal_StringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Equal_IntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Equal_Float64Slice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SetDiff_StringStringMap(a, b map[string]string) map[string]string {
	diff := make(map[string]string)
	for k, v := range a {
		if _, ok := b[k]; !ok {
			diff[k] = v
		}
	}
	return diff
}
