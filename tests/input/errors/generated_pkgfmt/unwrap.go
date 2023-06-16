package errors

func (e *joinError) Unwrap() []error {
	return e.errs
}
