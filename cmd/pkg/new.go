package pkg

func New(filePath string) (*Pkg, error) {
	p := Pkg{FilePath: filePath}
	if err := p.LoadBody(filePath); err != nil {
		return &p, err
	}
	if err := p.LoadAst(filePath); err != nil {
		return &p, err
	}
	p.LoadImports()
	p.LoadFromAst()
	return &p, nil
}
