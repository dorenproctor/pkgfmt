package errors

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

type joinError struct {
	errs []error
}
