# pkgfmt

`pkgfmt` is a CLI tool that formats Go packages.

---

## Using pkgfmt

`pkgfmt myfile.go` - Use it on an individual go file

`pkgfmt ./internal/mypkg` - Use it on a whole directory of go files

---

## What this does

A new directory `generated_pkgfmt/` is created inside the passed directory (or directory containing the file) containing:

- `types.go` - All declarations that use `type`, eg `struct`, `interface`, etc.
- `vars.go` - All `var` and `const` declarations
- A file per `func` name\*, converted to snakecase

The package declaration and the imports used by the code in each of these files is also included in the file contents.

**Test files are not included.**

<sub>\*There can be multiple funcs with the same name in which case they are put into the same file.</sub>

---

## Installing

You can install the latest version of this tool with:

`go install github.com/dorenproctor/pkgfmt@latest`

In order to run it, either:

1. Make sure ~/go/bin is in your PATH: `export PATH="$PATH:$HOME/go/bin"` - Probably put this somewhere like ~/.bashrc or ~/.profile
2. Move it to some other directory that is already in your path, eg `sudo mv $HOME/go/bin/pkgfmt /usr/local/bin/pkgfmt`

---

## Example

Ok, that's a lot of words. So what does this look like in practice?

(See more examples inside [tests/input/](./tests/input/))

Let's look at a snippet of code that used to be in this codebase. It was not reliable and all code detection is now done using `go/ast`, but we can see what output it generates.

`pkgfmt ./input.go`

**input.go**

```go
package impt

import "strings"

// Import
type Impt struct {
	Alias string
	Body  string
	Name  string
}

// Get imports referenced in a string - not reliable
func GetUsedImportsStr(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if importUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}

func importUsed(i Impt, body string) bool {
	if i.Alias != "" {
		if strings.Contains(body, i.Alias+".") {
			return true
		}
	} else {
		if strings.Contains(body, i.Name+".") {
			return true
		}
	}
	return false
}
```

**generated_pkgfmt/get_used_imports_str.go**

```go
package impt

// Get imports referenced in a string - not reliable
func GetUsedImportsStr(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if importUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}
```

**generated_pkgfmt/import_used.go**

```go
package impt

import (
	"strings"
)

func importUsed(i Impt, body string) bool {
	if i.Alias != "" {
		if strings.Contains(body, i.Alias+".") {
			return true
		}
	} else {
		if strings.Contains(body, i.Name+".") {
			return true
		}
	}
	return false
}
```

**generated_pkgfmt/types.go**

```go
package impt

type Impt struct {
	Alias string
	Body  string
	Name  string
}
```

---

## Known Limitations

- Docstrings for `type` declarations are lost.
- Currently any loose comments (ie not docstrings) are lost. `go/ast` simply does not include these. I've thought of a couple ways these could be detected and added to a `comments.go` file, but this is not implemented.
- Test files are not included.
- The `version` command does not work. In [run.sh](./scripts/run.sh) the ldflags variables are populating the variables as expected but the ldflags specified in [.goreleaser.yml](.goreleaser.yml) are not populating them in the published version and I have not been able to figure out why.

---

## When to use pkgfmt

Should you use this for everything? Should you enforce it in your CI/CD pipeline? If you want to, then sure! But probably not. In fact, not all of the code inside this codebase follows the pattern used by the tool.

Sometimes you might want to declare a helper function or variables only used by one function in the same file. Maybe you want to define the one method on a struct in the same file you declare it. There are various other reasons you may not want to enforce the pattern used by this tool.
