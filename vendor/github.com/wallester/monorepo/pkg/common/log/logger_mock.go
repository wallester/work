package log

import (
	"github.com/go-openapi/strfmt"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	logapi "github.com/wallester/monorepo/pkg/common/log/api"
)

type LoggerMock struct {
	mock.Mock
}

var _ logapi.ILogger = (*LoggerMock)(nil)

func (m *LoggerMock) Bytes(name string, value []byte) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Int(name string, value int) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Int64(name string, value int64) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) String(name, value string) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Any(name string, value any) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) UUID4(name string, value strfmt.UUID4) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Strings(name string, value []string) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Field(name string, value any) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Map(fields map[string]any) logapi.ILogger {
	m.Called(fields)
	return m
}

func (m *LoggerMock) Bool(name string, value bool) logapi.ILogger {
	m.Called(name, value)
	return m
}

func (m *LoggerMock) Write(p []byte) (int, error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *LoggerMock) Error(err error, fieldArgs ...map[string]any) {
	m.Called(err, fieldArgs)
}

func (m *LoggerMock) Fatal(err error, fieldArgs ...map[string]any) {
	m.Called(err, fieldArgs)
}

func (m *LoggerMock) Panic(err error, fieldArgs ...map[string]any) {
	m.Called(err, fieldArgs)
}

func (m *LoggerMock) Warn(err error, fieldArgs ...map[string]any) {
	m.Called(err, fieldArgs)
}

func (m *LoggerMock) Debug(msg string, fieldArgs ...map[string]any) {
	m.Called(msg, fieldArgs)
}

func (m *LoggerMock) Trace(msg string, fieldArgs ...map[string]any) {
	m.Called(msg, fieldArgs)
}

func (m *LoggerMock) Info(msg string, fieldArgs ...map[string]any) {
	m.Called(msg, fieldArgs)
}

func (m *LoggerMock) NewEntry() logapi.IEntry {
	args := m.Called()
	return args.Get(0).(logapi.IEntry)
}

func (m *LoggerMock) ZeroLogger() zerolog.Logger {
	args := m.Called()
	return args.Get(0).(zerolog.Logger)
}

func (m *LoggerMock) NewChildLogger() logapi.ILogger {
	args := m.Called()
	return args.Get(0).(logapi.ILogger)
}

func (m *LoggerMock) IsChildLogger() bool {
	args := m.Called()
	return args.Bool(0)
}
