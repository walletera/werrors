package werrors

import "fmt"

type InternalError struct {
    wError
}

func newInternalError(msg string, retryable bool) InternalError {
    return InternalError{
        wError{
            retryable: retryable,
            code:      InternalErrorCode,
            message:   fmt.Sprintf("internal error: %s", msg),
        },
    }
}

func NewRetryableInternalError(msg string) InternalError {
    return newInternalError(msg, true)
}

func NewNonRetryableInternalError(msg string) InternalError {
    return newInternalError(msg, false)
}

type AggregateNotFoundError struct {
    wError
}

func NewAggregateNotFoundError(msg string) AggregateNotFoundError {
    return AggregateNotFoundError{
        wError{
            retryable: false,
            code:      AggregateNotFoundErrorCode,
            message:   fmt.Sprintf("Aggregate not found error: %s", msg),
        },
    }
}

type TimeoutError struct {
    wError
}

func NewTimeoutError(msg string) TimeoutError {
    return TimeoutError{
        wError{
            retryable: true,
            code:      TimeoutErrorCode,
            message:   fmt.Sprintf("timeout error: %s", msg),
        },
    }
}

type ValidationError struct {
    wError
}

func NewValidationError(msg string) ValidationError {
    return ValidationError{
        wError{
            retryable: true,
            code:      ValidationErrorCode,
            message:   fmt.Sprintf("validation error: %s", msg),
        },
    }
}

type AggregateAlreadyExistError struct {
    wError
}

func NewAggregateAlreadyExistError(msg string) AggregateAlreadyExistError {
    return AggregateAlreadyExistError{
        wError{
            retryable: false,
            code:      AggregateAlreadyExistErrorCode,
            message:   fmt.Sprintf("Aggregate already exist error: %s", msg),
        },
    }
}

type WrongAggregateVersionError struct {
    wError
}

func NewWrongAggregateVersionError(msg string) WrongAggregateVersionError {
    return WrongAggregateVersionError{
        wError{
            retryable: false,
            code:      WrongAggregateVersionErrorCode,
            message:   fmt.Sprintf("wrong Aggregate version error: %s", msg),
        },
    }
}
