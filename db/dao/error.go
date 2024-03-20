package dao

type DAOError struct {
	message    string
	statusCode int
}

func NewDAOError(message string, statusCode int) *DAOError {
	return &DAOError{message: message, statusCode: statusCode}
}
