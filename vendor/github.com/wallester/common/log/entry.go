package log

import (
	"net/http"
)

type Entry struct {
	fields map[string]interface{}
}

func NewEntry() *Entry {
	return &Entry{
		fields: make(map[string]interface{}),
	}
}

// Typed parameters.
func (e *Entry) Bytes(name string, value []byte) *Entry  { return e.Field(name, value) }
func (e *Entry) Int(name string, value int) *Entry       { return e.Field(name, value) }
func (e *Entry) String(name string, value string) *Entry { return e.Field(name, value) }

// Convenience field setter.
func (e *Entry) Field(name string, value interface{}) *Entry {
	e.fields[name] = value
	return e
}

// Convenience HTTP response log entry.
func (e *Entry) Response(res *http.Response, body []byte) *Entry {
	return e.
		String(FieldNameStatus, res.Status).
		Int(FieldNameStatusCode, res.StatusCode).
		Bytes(FieldNameBody, body)
}

// Named parameters.
func (e *Entry) AuthorizationID(value string) *Entry {
	return e.String(FieldNameAuthorizationID, value)
}

func (e *Entry) AccountID(value string) *Entry { return e.String(FieldNameAccountID, value) }
func (e *Entry) Address(value string) *Entry   { return e.String(FieldNameAddress, value) }
func (e *Entry) CardID(value string) *Entry    { return e.String(FieldNameCardID, value) }
func (e *Entry) CardMetadataProfileID(value string) *Entry {
	return e.String(FieldNameCardMetadataProfileID, value)
}

func (e *Entry) CardProcessorCardID(value string) *Entry {
	return e.String(FieldNameCardProcessorCardID, value)
}

func (e *Entry) ClientID(value string) *Entry       { return e.String(FieldNameClientID, value) }
func (e *Entry) ContractID(value string) *Entry     { return e.String(FieldNameContractID, value) }
func (e *Entry) Directory(value string) *Entry      { return e.String(FieldNameDirectory, value) }
func (e *Entry) Email(value string) *Entry          { return e.String(FieldNameEmail, value) }
func (e *Entry) Filename(value string) *Entry       { return e.String(FieldNameFilename, value) }
func (e *Entry) Host(value string) *Entry           { return e.String(FieldNameHost, value) }
func (e *Entry) Locale(value string) *Entry         { return e.String(FieldNameLocale, value) }
func (e *Entry) PersonalNumber(value string) *Entry { return e.String(FieldNamePersonalNumber, value) }
func (e *Entry) PersonID(value string) *Entry       { return e.String(FieldNamePersonID, value) }
func (e *Entry) Product(value string) *Entry        { return e.String(FieldNameProduct, value) }
func (e *Entry) ProductID(value string) *Entry      { return e.String(FieldNameProductID, value) }
func (e *Entry) ReferenceNumber(value string) *Entry {
	return e.String(FieldNameReferenceNumber, value)
}

func (e *Entry) RequestID(value string) *Entry     { return e.String(FieldNameRequestID, value) }
func (e *Entry) SearchValue(value string) *Entry   { return e.String(FieldNameSearchValue, value) }
func (e *Entry) ServiceName(value string) *Entry   { return e.String(FieldNameServiceName, value) }
func (e *Entry) SQSMessage(value string) *Entry    { return e.String(FieldNameSQSMessage, value) }
func (e *Entry) SQSQueue(value string) *Entry      { return e.String(FieldNameSQSQueue, value) }
func (e *Entry) Template(value string) *Entry      { return e.String(FieldNameTemplate, value) }
func (e *Entry) UserID(value string) *Entry        { return e.String(FieldNameUserID, value) }
func (e *Entry) CorrelationID(value string) *Entry { return e.String(FieldNameCorrelationID, value) }
func (e *Entry) ExceptionData(value string) *Entry { return e.String(FieldNameExceptionData, value) }
func (e *Entry) ErrorField(value string) *Entry    { return e.String(FieldNameErrorField, value) }
func (e *Entry) Body(value string) *Entry          { return e.String(FieldNameBody, value) }
func (e *Entry) Method(value string) *Entry        { return e.String(FieldNameMethod, value) }
func (e *Entry) URL(value string) *Entry           { return e.String(FieldNameURL, value) }
func (e *Entry) LanguageCode(value string) *Entry  { return e.String(FieldNameLanguageCode, value) }
func (e *Entry) MessageType(value string) *Entry   { return e.String(FieldNameMessageType, value) }
func (e *Entry) CompanyID(value string) *Entry     { return e.String(FieldNameCompanyID, value) }
func (e *Entry) DeviceType(value string) *Entry    { return e.String(FieldNameDeviceType, value) }
func (e *Entry) Provider(value string) *Entry      { return e.String(FieldNameProvider, value) }

// Call any of the following to write the log entry.
func (e *Entry) Debug(msg string) { Debug(msg, e.fields) }
func (e *Entry) Error(err error)  { Error(err, e.fields) }
func (e *Entry) Fatal(err error)  { Fatal(err, e.fields) }
func (e *Entry) Info(msg string)  { Info(msg, e.fields) }
func (e *Entry) Panic(err error)  { Panic(err, e.fields) }
func (e *Entry) Warn(err error)   { Warn(err, e.fields) }
