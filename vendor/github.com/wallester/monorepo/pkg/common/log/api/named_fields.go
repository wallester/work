package api

import (
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
)

type INamedFields interface {
	AWSMessageID(value string) IEntry
	AccountID(value string) IEntry
	Address(value string) IEntry
	AdjustmentID(value string) IEntry
	AuthorizationID(value string) IEntry
	BatchID(batchID strfmt.UUID4) IEntry
	Body(value string) IEntry
	Bucket(value string) IEntry
	CardID(value string) IEntry
	CardMetadataProfileID(value string) IEntry
	CardProcessorCardID(value string) IEntry
	ClientID(value string) IEntry
	CompanyID(value string) IEntry
	ContractID(value string) IEntry
	CorrelationID(value string) IEntry
	DatabaseNotificationMessage(value string) IEntry
	DeviceType(value string) IEntry
	Directory(value string) IEntry
	Duration(value time.Duration) IEntry
	Email(value string) IEntry
	ErrorField(value string) IEntry
	ExceptionData(value string) IEntry
	ExitCode(value string) IEntry
	FeeID(value string) IEntry
	FeeType(value string) IEntry
	FileType(value string) IEntry
	Filename(value string) IEntry
	Hash(value string) IEntry
	Host(value string) IEntry
	JobName(value string) IEntry
	Key(value string) IEntry
	LanguageCode(value string) IEntry
	Locale(value string) IEntry
	MemoryUsage() IEntry
	MessageType(value string) IEntry
	Method(value string) IEntry
	NotificationChannel(value string) IEntry
	Path(value string) IEntry
	PersonID(value string) IEntry
	PersonalNumber(value string) IEntry
	Product(value string) IEntry
	ProductID(value string) IEntry
	Provider(value string) IEntry
	ReferenceNumber(value string) IEntry
	RepresentativeID(value string) IEntry
	RequestID(value string) IEntry
	Response(res *http.Response, body []byte) IEntry
	Retry(value int) IEntry
	RetryCount(value int) IEntry
	RetrySleepTime(value time.Duration) IEntry
	SQSMessage(value string) IEntry
	SQSQueue(value string) IEntry
	SearchValue(value string) IEntry
	ServiceName(value string) IEntry
	ShutdownError(value string) IEntry
	StepUpRequestID(value string) IEntry
	Template(value string) IEntry
	URL(value string) IEntry
	UserID(value string) IEntry
	UserIPID(value string) IEntry
	XRequestID(value string) IEntry
	ProductSettings(value map[string]any) IEntry
}
