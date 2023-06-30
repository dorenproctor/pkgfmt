package pkgpart

// combines funcs that have identical names into groups
func GroupDuplicatesFns(fns []PkgPart) map[string][]PkgPart {
	m := make(map[string][]PkgPart)
	for _, fn := range fns {
		if _, exists := m[fn.Name]; !exists {
			m[fn.Name] = []PkgPart{}
		}
		m[fn.Name] = append(m[fn.Name], fn)
	}
	return m
}
