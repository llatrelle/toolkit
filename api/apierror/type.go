package apierror

func New(err string) apiError {
	a := apiError{message: err}
	return a
}

type apiError struct {
	code    int
	message string
}

func (a *apiError) Error() string {
	return (*a).message
}

func (a *apiError) SetErrorCode(code int) {
	a.code = code
}

func (a *apiError) GetErrorCode() int {
	return a.code
}
