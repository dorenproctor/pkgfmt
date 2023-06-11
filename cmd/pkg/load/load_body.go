package load

import "io/ioutil"

func (p *types.Pkg) LoadBody(filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	p.Body = string(fileBytes)
	return nil
}
