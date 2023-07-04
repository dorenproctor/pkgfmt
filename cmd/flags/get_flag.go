package flags

// returns the value of the flag or nil if there is no value
//
// must be inside the flags var, not just declared with flag package
//
// eg. flags.GetFlag[bool]("verbose")
func GetFlag[T comparable](flagName string) *T {
	s, ok := flags[flagName].(*T)
	if !ok {
		return nil
	}
	return s
}
