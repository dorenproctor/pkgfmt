package sliceutils

// Remove duplicates from slices of any comparable type
//
// https://gosamples.dev/generics-remove-duplicates-slice/
func RemoveDuplicates[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
