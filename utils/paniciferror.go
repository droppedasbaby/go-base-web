package utils

import "errors"

// PanicIfError panics if the error is not nil and not in the list of ignored errors.
func PanicIfError(err error, ignoredErrs ...error) {
	if err != nil && !isErrorInList(err, ignoredErrs...) {
		panic(err)
	}
}

// isErrorInList checks if the error is in the list of errors.
func isErrorInList(err error, errs ...error) bool {
	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
