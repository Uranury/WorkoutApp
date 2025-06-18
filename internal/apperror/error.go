package apperror

import "net/http"

type AppError struct {
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, statuscode int) *AppError {
	return &AppError{Message: message, StatusCode: statuscode}
}

var (
	ErrEmailAlreadyInUse = &AppError{
		Message:    "email already in use",
		StatusCode: http.StatusConflict,
	}

	ErrUsernameAlreadyInUse = &AppError{
		Message:    "username already in use",
		StatusCode: http.StatusConflict,
	}

	ErrUserNotFound = &AppError{
		Message:    "user not found",
		StatusCode: http.StatusNotFound,
	}

	ErrInvalidCredentials = &AppError{
		Message:    "incorrect password",
		StatusCode: http.StatusUnauthorized,
	}

	ErrEmailPasswordRequired = &AppError{
		Message:    "email and password are required",
		StatusCode: http.StatusBadRequest,
	}

	ErrNoUsersFound = &AppError{
		Message:    "no users yet",
		StatusCode: http.StatusNotFound,
	}

	ErrExerciseAlreadyExists = &AppError{
		Message:    "exercise already exists",
		StatusCode: http.StatusConflict,
	}
)

// Generic errors
var (
	ErrInternalServer = &AppError{
		Message:    "internal server error",
		StatusCode: http.StatusInternalServerError,
	}

	ErrDatabaseError = &AppError{
		Message:    "database error",
		StatusCode: http.StatusInternalServerError,
	}

	ErrHashPassword = &AppError{
		Message:    "failed to hash password",
		StatusCode: http.StatusInternalServerError,
	}

	ErrGenerateToken = &AppError{
		Message:    "failed to generate token",
		StatusCode: http.StatusInternalServerError,
	}

	ErrBadRequest = &AppError{
		Message:    "invalid request",
		StatusCode: http.StatusBadRequest,
	}

	ErrUnauthorized = &AppError{
		Message:    "unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
)
