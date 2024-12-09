package errors

import (
	"github.com/juju/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IErrorWithCustomData is interface that any CustomData error must satisfy.
type IErrorWithCustomData interface {
	error
	gRPCStatus
	CustomData() *CustomData
	Source() error
}

// This is a compile-time assertion to ensure that errorWithCustomData satisfies IErrorWithCustomData interface.
var _ IErrorWithCustomData = &errorWithCustomData{}

// NewErrorWithCustomData is method to create new error from message with custom data specified.
// In case you do not need it consider using errors.New(m string) or errors.Errorf(m string) method from juju/errors package.
func NewErrorWithCustomData(message string, customData *CustomData) error {
	return &errorWithCustomData{
		customData: customData,
		source:     errors.New(message),
	}
}

// ErrorWithCustomData is method to create error with custom data specified.
// Consider using errors.New(m string) or errors.Annotatef(other error, m string, args ...any) method from juju/errors package if you do not need custom data.
func ErrorWithCustomData(err error, customData *CustomData) error {
	return &errorWithCustomData{
		customData: customData,
		source:     err,
	}
}

// AnnotateWithCustomData is method to annotate another error with a message.
// In case other error is a custom data error, then it will attempt to merge custom data from both errors.
// Consider using errors.Annotate(other error, m string) or errors.Annotatef(other error, m string, args ...any) method from juju/errors package if you do not need custom data.
func AnnotateWithCustomData(other error, message string, customData *CustomData) error {
	if other == nil {
		return nil
	}

	if e, ok := other.(IErrorWithCustomData); ok {
		customData = e.CustomData().Merge(customData)
	}

	return &errorWithCustomData{
		source:     errors.Annotate(other, message),
		customData: customData,
	}
}

// Source is a method that returns source error from custom data error.
func (err errorWithCustomData) Source() error {
	return err.source
}

// Error is a method that returns string representation for error.
// It is needed to implement Error() interface of standard golang error library.
func (err errorWithCustomData) Error() string {
	return err.source.Error()
}

// CustomData is a method that returns custom data.
func (err errorWithCustomData) CustomData() *CustomData {
	return err.customData
}

// GRPCStatus is a method that returns gRPC status error based on the source error.
// It is needed to implement GRPCStatus() interface of gRPC library.
func (err errorWithCustomData) GRPCStatus() *status.Status {
	s, ok := err.source.(gRPCStatus)
	if ok {
		return s.GRPCStatus()
	}

	return status.New(codes.Unknown, err.source.Error())
}

// private

type gRPCStatus interface {
	GRPCStatus() *status.Status
}

type errorWithCustomData struct {
	customData *CustomData
	source     error
}
