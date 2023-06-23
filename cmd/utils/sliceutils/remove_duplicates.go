package sliceutils

// Remove duplicates from slices of any comparable type
//
// https://gosamples.dev/generics-remove-duplicates-slice/
func RemoveDuplicates[T comparable](s []T) []T {
	cache := map[T]bool{}
	var output []T
	for _, t := range s {
		if _, seen := cache[t]; !seen {
			cache[t] = true
			output = append(output, t)
		}
	}
	return output
}
