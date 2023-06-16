package errors

type errorString struct {
	s string
}

type joinError struct {
	errs []error
}
