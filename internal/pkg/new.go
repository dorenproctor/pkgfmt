package pkg

func New(filePath string) (*Package, error) {
	p := Package{}
	if err := p.LoadBody(filePath); err != nil {
		return &p, err
	}
	if err := p.LoadAst(filePath); err != nil {
		return &p, err
	}
	p.LoadImports()
	p.LoadFns()
	return &p, nil
}
