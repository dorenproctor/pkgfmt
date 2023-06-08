package fn

import (
	"fmt"
	"pkgsplit/internal/impt"
)

func (f *Fn) GetFnFileOutput() string {
	return fmt.Sprintf(`package %s
%s
%s`, f.PackageName, impt.GetImportStr(f.Imports), f.Body)
}
