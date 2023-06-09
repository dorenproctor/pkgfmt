# pkgfmt

`pkgfmt` is a CLI tool that formats Go packages.

This has 0 dependencies on third-party libraries! Only libraries built into the core Go library are used.

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
	// Full import path declared in import,
	// eg "go/ast"
	FullName string

	// Name of imported package, after last / of import path,
	// eg "ast" from `import "go/ast"`
	Name string

	// Optional alias variable before the import,
	// eg myalias from `import myalias "go/ast"`
	Alias string
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

// Import
type Impt struct {
	// Full import path declared in import,
	// eg "go/ast"
	FullName string

	// Name of imported package, after last / of import path,
	// eg "ast" from `import "go/ast"`
	Name string

	// Optional alias variable before the import,
	// eg myalias from `import myalias "go/ast"`
	Alias string
}

```

---

## Known Limitations

- Test files are not included when passing a dir. You can pass them individually though.
- The `version` command does not work. In [run.sh](./scripts/run.sh) the ldflags variables are populating the variables as expected but the ldflags specified in [.goreleaser.yml](.goreleaser.yml) are not populating them in the published version and I have not been able to figure out why.

---

## Potential Future Features

- Maybe a flag to specify output dir
- Maybe a flag for including test files when formatting a whole dir
- Maybe some way to indicate whether to use () for individual variable declarations and imports

---

## Why I made pkgfmt

I like having a file for each function in Go. This pattern doesn't really work in most languages but it does in Go. I've run across large Go files that I split up using this pattern to make it easier for me to understand and see what's there. I found creating the file in snake case, copying the function over, adding the package name, and adding the imports (the last step is done by my editor at least) to be a bit tedious. Not terribly so, but enough to make me wonder if I could automate it. I added steps as they seemed necessary. And here we are.

---

## When to use pkgfmt

When you have a Go file or package that is too big for you to follow, use this to split it. If you like the code generated by pkgfmt as is, awesome. Delete the old files and replace them with your freshly formatted code. Otherwise, copy the files you want to keep, delete the duplicate code in the other files, and get rid of the rest.

Should you use this for everything? Should you enforce it in your CI/CD pipeline? If you want to, then sure! But probably not. In fact, not all of the code inside this codebase follows the pattern used by the tool.

Sometimes you might want to declare a helper function or variables only used by one function in the same file. Maybe you want to define the one method on a struct in the same file you declare it. There are various other reasons you may not want to enforce the pattern used by this tool.
