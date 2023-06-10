package impt

// not reliable
func GetUsedImportsStr(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if importUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}
