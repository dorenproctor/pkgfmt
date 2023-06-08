package example

import "fmt"

func x() {
	fmt.Println(yzzy("x!"))
}

func yzzy(s string) string {
	return s
}
