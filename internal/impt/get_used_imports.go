package impt

func GetUsedImports(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if ImportUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}
