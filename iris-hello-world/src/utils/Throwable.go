package utils

type Throwable interface {
	SetMessage(message string)
	GetMessage()
	SetCode(code int)
	GetCode()
}
