package booksservice

type InvalidAuthorError struct {
	Message string
}

func (e *InvalidAuthorError) Error() string {
	return e.Message
}

type InvalidTitleError struct {
	Message string
}

func (e *InvalidTitleError) Error() string {
	return e.Message
}

type InvalidIDError struct {
	Message string
}

func (e *InvalidIDError) Error() string {
	return e.Message
}

type InvalidCreationTimeError struct {
	Message string
}

func (e *InvalidCreationTimeError) Error() string {
	return e.Message
}
