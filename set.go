package generics

func SetDiff_StringStringMap(a, b map[string]string) map[string]string {
	diff := make(map[string]string)
	for k, v := range a {
		if _, ok := b[k]; !ok {
			diff[k] = v
		}
	}
	return diff
}
