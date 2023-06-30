package fn

import "github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"

// combines funcs that have identical names into groups
func GroupDuplicatesFns(fns []gofile.PkgPart) map[string][]gofile.PkgPart {
	m := make(map[string][]gofile.PkgPart)
	for _, fn := range fns {
		if _, exists := m[fn.Name]; !exists {
			m[fn.Name] = []gofile.PkgPart{}
		}
		m[fn.Name] = append(m[fn.Name], fn)
	}
	return m
}