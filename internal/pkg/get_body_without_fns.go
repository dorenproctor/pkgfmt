package pkg

func (p *Package) GetBodyWithoutFns() string {
	s := p.Body
	for i := len(p.Fns) - 1; i >= 0; i-- {
		fn := p.Fns[i]
		s = s[0:fn.LPos] + s[fn.RPos:]
	}
	return s
}
