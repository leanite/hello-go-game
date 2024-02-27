package errors

type OutOfManaError struct {
	Message string
}

// Implementação do método Error() para o tipo OutOfManaError.
func (e *OutOfManaError) Error() string {
	return e.Message
}
