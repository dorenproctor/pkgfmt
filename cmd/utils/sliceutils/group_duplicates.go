package sliceutils

func GroupDuplicates[T comparable](s []T) map[T][]T {
	m := make(map[T][]T)
	for _, fn := range s {
		if _, exists := m[fn]; !exists {
			m[fn] = []T{}
		}
		m[fn] = append(m[fn], fn)
	}
	return m
}
