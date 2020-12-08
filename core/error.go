package core

type Error struct {
	ErrMsg string
}

func NewError(msg string) *Error {
	return &Error{ErrMsg: msg}
}

func(err *Error) Error() string {
	return err.ErrMsg
}
