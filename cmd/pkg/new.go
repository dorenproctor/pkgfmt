package pkg

import "github.com/dorenproctor/pkgfmt/cmd/impt"

func New(filePath string) (*Pkg, error) {
	p := Pkg{FilePath: filePath}
	if err := p.LoadBody(filePath); err != nil {
		return &p, err
	}
	if err := p.LoadAst(filePath); err != nil {
		return &p, err
	}
	p.Imports = impt.LoadImports(p.Ast)
	p.LoadFromAst()
	return &p, nil
}
