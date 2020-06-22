package core

type Error struct {
	Code        string
	Description string
}

func (err *Error) Error() string {
	return err.Description
}

func NewError(code, description string) *Error {
	return &Error{
		Code:        code,
		Description: description,
	}
}
