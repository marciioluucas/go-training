package utils

type Exception struct {
	message string
	code    string
}

func (Exception) New(params ...string) Exception {
	if len(params) > 2 {
		panic("Exception have only 2 parameters, the first is 'message' and the second is 'code'.")
	}
	code := "0"
	if len(params[1]) > 0 {
		code = params[1]
	}
	return Exception{params[0], code}
}

func (e *Exception) SetMessage(message string) *Exception {
	e.message = message
	return e
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) SetCode(code string) {
	e.code = code
}

func (e *Exception) GetCode() string {
	return e.code
}
