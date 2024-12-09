package log

import (
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
	logapi "github.com/wallester/monorepo/pkg/common/log/api"
	"github.com/wallester/monorepo/pkg/common/runtime"
)

func (e *Entry) Response(res *http.Response, body []byte) logapi.IEntry {
	return e.
		String(FieldNameStatus, res.Status).
		Int(FieldNameStatusCode, res.StatusCode).
		Bytes(FieldNameBody, body)
}

func (e *Entry) AuthorizationID(value string) logapi.IEntry {
	return e.String(FieldNameAuthorizationID, value)
}

func (e *Entry) AccountID(value string) logapi.IEntry { return e.String(FieldNameAccountID, value) }
func (e *Entry) AWSMessageID(value string) logapi.IEntry {
	return e.String(FieldNameAWSMessageID, value)
}
func (e *Entry) FeeID(value string) logapi.IEntry   { return e.String(FieldNameFeeID, value) }
func (e *Entry) FeeType(value string) logapi.IEntry { return e.String(FieldNameFeeType, value) }
func (e *Entry) AdjustmentID(value string) logapi.IEntry {
	return e.String(FieldNameAdjustmentID, value)
}
func (e *Entry) Address(value string) logapi.IEntry { return e.String(FieldNameAddress, value) }
func (e *Entry) CardID(value string) logapi.IEntry  { return e.String(FieldNameCardID, value) }
func (e *Entry) CardMetadataProfileID(value string) logapi.IEntry {
	return e.String(FieldNameCardMetadataProfileID, value)
}

func (e *Entry) CardProcessorCardID(value string) logapi.IEntry {
	return e.String(FieldNameCardProcessorCardID, value)
}

func (e *Entry) ClientID(value string) logapi.IEntry   { return e.String(FieldNameClientID, value) }
func (e *Entry) ContractID(value string) logapi.IEntry { return e.String(FieldNameContractID, value) }
func (e *Entry) Directory(value string) logapi.IEntry  { return e.String(FieldNameDirectory, value) }
func (e *Entry) NotificationChannel(value string) logapi.IEntry {
	return e.String(FieldNameNotificationChannel, value)
}

func (e *Entry) DatabaseNotificationMessage(value string) logapi.IEntry {
	return e.String(FieldNameDatabaseNotificationMessage, value)
}

func (e *Entry) Email(value string) logapi.IEntry    { return e.String(FieldNameEmail, value) }
func (e *Entry) Filename(value string) logapi.IEntry { return e.String(FieldNameFilename, value) }
func (e *Entry) Host(value string) logapi.IEntry     { return e.String(FieldNameHost, value) }
func (e *Entry) UserIPID(value string) logapi.IEntry { return e.String(FieldNameIPID, value) }
func (e *Entry) Locale(value string) logapi.IEntry   { return e.String(FieldNameLocale, value) }
func (e *Entry) PersonalNumber(value string) logapi.IEntry {
	return e.String(FieldNamePersonalNumber, value)
}
func (e *Entry) PersonID(value string) logapi.IEntry  { return e.String(FieldNamePersonID, value) }
func (e *Entry) Product(value string) logapi.IEntry   { return e.String(FieldNameProduct, value) }
func (e *Entry) ProductID(value string) logapi.IEntry { return e.String(FieldNameProductID, value) }
func (e *Entry) ReferenceNumber(value string) logapi.IEntry {
	return e.String(FieldNameReferenceNumber, value)
}

func (e *Entry) RequestID(value string) logapi.IEntry   { return e.String(FieldNameRequestID, value) }
func (e *Entry) XRequestID(value string) logapi.IEntry  { return e.String(FieldNameXRequestID, value) }
func (e *Entry) SearchValue(value string) logapi.IEntry { return e.String(FieldNameSearchValue, value) }
func (e *Entry) ServiceName(value string) logapi.IEntry { return e.String(FieldNameServiceName, value) }
func (e *Entry) SQSMessage(value string) logapi.IEntry  { return e.String(FieldNameSQSMessage, value) }
func (e *Entry) SQSQueue(value string) logapi.IEntry    { return e.String(FieldNameSQSQueue, value) }
func (e *Entry) Template(value string) logapi.IEntry    { return e.String(FieldNameTemplate, value) }
func (e *Entry) UserID(value string) logapi.IEntry      { return e.String(FieldNameUserID, value) }
func (e *Entry) CorrelationID(value string) logapi.IEntry {
	return e.String(FieldNameCorrelationID, value)
}

func (e *Entry) ExceptionData(value string) logapi.IEntry {
	return e.String(FieldNameExceptionData, value)
}
func (e *Entry) ErrorField(value string) logapi.IEntry { return e.String(FieldNameErrorField, value) }
func (e *Entry) Body(value string) logapi.IEntry       { return e.String(FieldNameBody, value) }
func (e *Entry) Method(value string) logapi.IEntry     { return e.String(FieldNameMethod, value) }
func (e *Entry) URL(value string) logapi.IEntry        { return e.String(FieldNameURL, value) }
func (e *Entry) LanguageCode(value string) logapi.IEntry {
	return e.String(FieldNameLanguageCode, value)
}
func (e *Entry) MessageType(value string) logapi.IEntry { return e.String(FieldNameMessageType, value) }
func (e *Entry) CompanyID(value string) logapi.IEntry   { return e.String(FieldNameCompanyID, value) }
func (e *Entry) DeviceType(value string) logapi.IEntry  { return e.String(FieldNameDeviceType, value) }
func (e *Entry) Provider(value string) logapi.IEntry    { return e.String(FieldNameProvider, value) }
func (e *Entry) ExitCode(value string) logapi.IEntry    { return e.String(FieldNameExitCode, value) }
func (e *Entry) ShutdownError(value string) logapi.IEntry {
	return e.String(FieldNameShutdownError, value)
}
func (e *Entry) Bucket(value string) logapi.IEntry  { return e.String(FieldNameBucket, value) }
func (e *Entry) Key(value string) logapi.IEntry     { return e.String(FieldNameKey, value) }
func (e *Entry) JobName(value string) logapi.IEntry { return e.String(FieldNameJobName, value) }
func (e *Entry) StepUpRequestID(value string) logapi.IEntry {
	return e.String(FieldNameStepUpRequestID, value)
}

func (e *Entry) RepresentativeID(value string) logapi.IEntry {
	return e.String(FieldNameRepresentativeID, value)
}

func (e *Entry) Duration(value time.Duration) logapi.IEntry {
	return e.String(FieldNameDuration, value.String())
}

func (e *Entry) ProductSettings(value map[string]any) logapi.IEntry {
	return e.Field(FieldNameProductSettings, value)
}

func (e *Entry) MemoryUsage() logapi.IEntry {
	memoryUsage := runtime.NewMemoryUsageMessage()
	return e.Map(map[string]any{
		FieldNameAlloc:      memoryUsage.Alloc,
		FieldNameTotalAlloc: memoryUsage.TotalAlloc,
		FieldNameSys:        memoryUsage.Sys,
		FieldNameNumGC:      memoryUsage.NumGC,
	})
}

func (e *Entry) BatchID(batchID strfmt.UUID4) logapi.IEntry {
	return e.UUID4(FieldNameBatchID, batchID)
}

func (e *Entry) Hash(value string) logapi.IEntry     { return e.String(FieldNameHash, value) }
func (e *Entry) Path(value string) logapi.IEntry     { return e.String(FieldNamePath, value) }
func (e *Entry) FileType(value string) logapi.IEntry { return e.String(FieldNameFileType, value) }
func (e *Entry) Retry(value int) logapi.IEntry       { return e.Int(FieldNameRetry, value) }
func (e *Entry) RetryCount(value int) logapi.IEntry  { return e.Int(FieldNameRetryCount, value) }
func (e *Entry) RetrySleepTime(value time.Duration) logapi.IEntry {
	return e.String(FieldNameRetry, value.String())
}
