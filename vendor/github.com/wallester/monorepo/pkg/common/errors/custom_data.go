package errors

// CustomData is struct that hold map with string keys and interface values.
// It is used to pass custom data to Rollbar/Logger as separate fields.
// Example:
// CustomData with key = person-id and value set to fbceab5f-1d55-4fba-9c36-dcce24c04122 will produce following log in HTTP
//
//	{
//	   "caller": "/Users/oxxy/go/src/github.com/wallester/monorepo/vendor/github.com/wallester/monorepo/pkg/common/http/log.go:29",
//	   "duration": "45.421ms",
//	   "error": "getting person failed: unexpected response status code 422: 422 Unprocessable Entity",
//	   "http-status": 500,
//	   "ip": "127.0.0.1",
//	   "level": 3,
//	   "level-string": "error",
//	   "method": "GET",
//	   "msg": "getting person failed: unexpected response status code 422: 422 Unprocessable Entity",
//	   "person-id": "fbceab5f-1d55-4fba-9c36-dcce24c04122",
//	   "request-body": "",
//	}.
type CustomData struct {
	fields map[string]any
}

// NewCustomData is a constructor method used to initialize empty map.
func NewCustomData() *CustomData {
	return &CustomData{
		fields: make(map[string]any),
	}
}

// Field allows to add custom data with name (string) and value (any).
// It is used as a helper method to store values of different types in a map.
func (data *CustomData) Field(name string, value any) *CustomData {
	data.fields[name] = value
	return data
}

// Int is a helper method to store values of type int in map without need to cast them to any.
func (data *CustomData) Int(name string, value int) *CustomData {
	return data.Field(name, value)
}

// Int64 is a helper method to store values of type int64 in map without need to cast them to any.
func (data *CustomData) Int64(name string, value int64) *CustomData {
	return data.Field(name, value)
}

// String is a helper method to store values of type string in map without need to cast them to any.
func (data *CustomData) String(name string, value string) *CustomData {
	return data.Field(name, value)
}

// Merge is a helper method to merge two different custom data maps, resulting in a copy with values from both src and dst.
func (data *CustomData) Merge(customData *CustomData) *CustomData {
	if customData == nil {
		return data
	}

	return data.Map(customData.fields)
}

// Map is a helper method that allows to convert map of any values to custom data type.
func (data *CustomData) Map(fields map[string]any) *CustomData {
	for k, v := range fields {
		data.Field(k, v)
	}

	return data
}

// Fields is a helper method that returns the map, it is used in logger and to help debug code.
func (data *CustomData) Fields() map[string]any {
	return data.fields
}

// AccountID is a helper method to store string value in account-id key.
func (data *CustomData) AccountID(value string) *CustomData {
	return data.String(FieldNameAccountID, value)
}

// CardID is a helper method to store string value in card-id key.
func (data *CustomData) CardID(value string) *CustomData {
	return data.String(FieldNameCardID, value)
}

// CardProcessorCardID is a helper method to store string value in card-processor-card-id key.
func (data *CustomData) CardProcessorCardID(value string) *CustomData {
	return data.String(FieldNameCardProcessorCardID, value)
}

// ClientID is a helper method to store string value in client-id key.
func (data *CustomData) ClientID(value string) *CustomData {
	return data.String(FieldNameClientID, value)
}

// Email is a helper method to store string value in email key.
func (data *CustomData) Email(value string) *CustomData {
	return data.String(FieldNameEmail, value)
}

// Mobile is a helper method to store string value in mobile key.
func (data *CustomData) Mobile(value string) *CustomData {
	return data.String(FieldNameMobile, value)
}

// PersonalNumber is a helper method to store string value in personal-number key.
func (data *CustomData) PersonalNumber(value string) *CustomData {
	return data.String(FieldNamePersonalNumber, value)
}

// PersonID is a helper method to store string value in person-id key.
func (data *CustomData) PersonID(value string) *CustomData {
	return data.String(FieldNamePersonID, value)
}

// ProductID is a helper method to store string value in product-id key.
func (data *CustomData) ProductID(value string) *CustomData {
	return data.String(FieldNameProductID, value)
}

// RequestID is a helper method to store string value in request-id key.
func (data *CustomData) RequestID(value string) *CustomData {
	return data.String(FieldNameRequestID, value)
}

// UserID is a helper method to store string value in user-id key.
func (data *CustomData) UserID(value string) *CustomData {
	return data.String(FieldNameUserID, value)
}

// AuthorizationID is a helper method to store string value in authorization-id key.
func (data *CustomData) AuthorizationID(value string) *CustomData {
	return data.String(FieldNameAuthorizationID, value)
}

// CompanyID is a helper method to store string value in company-id key.
func (data *CustomData) CompanyID(value string) *CustomData {
	return data.String(FieldNameCompanyID, value)
}

// TransactionID is a helper method to store string value in transaction-id key.
func (data *CustomData) TransactionID(value string) *CustomData {
	return data.String(FieldNameTransactionID, value)
}

// StepUpRequestID is a helper method to store string value in step-up-request-id key.
func (data *CustomData) StepUpRequestID(value string) *CustomData {
	return data.String(FieldNameStepUpRequestID, value)
}

// MaskedCardNumber is a helper method to store string value in masked-card-number key.
func (data *CustomData) MaskedCardNumber(value string) *CustomData {
	return data.String(FieldNameMaskedCardNumber, value)
}

// NextAttemptsAfter is a helper method to store string value in next_attempts_after key
func (data *CustomData) NextAttemptsAfter(value string) *CustomData {
	return data.String(FieldNameNextAttemptsAfter, value)
}

// AttemptsLeft is a helper method to store int64 value in attempts_left key
func (data *CustomData) AttemptsLeft(value int64) *CustomData {
	return data.Int64(FieldNameAttemptsLeft, value)
}
