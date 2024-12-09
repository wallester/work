package errors

const (
	// bank-service related error codes.
	ErrorCodeVibanNotFound                ErrorCode = 101001
	ErrorCodeVibanAlreadyLinkedOnAccount  ErrorCode = 101002
	ErrorCodeVibanFunctionalityIsDisabled ErrorCode = 101003
	ErrorCodeVibanProviderNotEnabled      ErrorCode = 101004
)
