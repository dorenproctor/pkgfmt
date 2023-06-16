package sliceutils

// Remove duplicates from slices of any comparable type
//
// https://gosamples.dev/generics-remove-duplicates-slice/
func RemoveDuplicates[T comparable](s []T) []T {
	alreadySeen := make(map[T]bool)
	var output []T
	for _, t := range s {
		if _, ok := alreadySeen[t]; !ok {
			alreadySeen[t] = true
			output = append(output, t)
		}
	}
	return output
}
