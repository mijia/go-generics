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
