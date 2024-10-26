package apperror

import "fmt"

type AppError struct {
	Err        error
	StructName string
	MethodName string
	RootTrace  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf(
		"%T error on %s.%s\n%s : %s",
		e.Err, e.StructName, e.MethodName, e.RootTrace, e.Err.Error(),
	)
}

func NewAppError(e error, structName, methodName, rootTrace string) *AppError {
	return &AppError{e, structName, methodName, rootTrace}
}
