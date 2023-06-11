package load

import (
	"strings"

	"github.com/dorenproctor/github.com/dorenproctor/pkgfmt/types"
	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *types.Pkg) LoadImports() {
	p.Imports = []impt.Impt{}
	for _, importSpec := range p.Ast.Imports {
		i := impt.Impt{Name: importSpec.Path.Value}
		i.NameWithQuotes = importSpec.Path.Value
		i.Name = strings.Replace(importSpec.Path.Value, "\"", "", 2)
		if importSpec.Name != nil {
			i.Alias = importSpec.Name.Name
		}
		p.Imports = append(p.Imports, i)
	}
}
