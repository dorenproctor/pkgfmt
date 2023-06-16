package impt

import "fmt"

func (i Impt) String() string {
	if i.Alias == "" {
		return i.NameWithQuotes
	}
	return fmt.Sprintf("%s %s", i.Alias, i.NameWithQuotes)
}
