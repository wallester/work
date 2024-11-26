package log

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/juju/errors"
	rollbar "github.com/rollbar/rollbar-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger implements a logger.
type Logger struct {
	zerolog.Logger
}

func init() {
	zerolog.TimestampFieldName = FieldNameTime
	zerolog.LevelFieldName = FieldNameLevelString
	zerolog.MessageFieldName = FieldNameMsg
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackFieldName = FieldNameErrorStack
	zerolog.CallerSkipFrameCount = 4
}

// New initializes a new logger instance and returns a func to close it.
func New(cfg Configuration) (*Logger, func()) {
	writer, closer := configureWriter(cfg)

	level := zerolog.Level(cfg.Level)
	logger := &Logger{
		zerolog.New(writer).Level(level).With().Timestamp().Logger().Hook(LevelHook{}),
	}

	if cfg.RollbarToken != "" {
		rollbar.SetToken(cfg.RollbarToken)
		rollbar.SetEnvironment(cfg.Environment)
		rollbar.SetStackTracer(func(_ error) ([]runtime.Frame, bool) {
			// Empty stack trace as it is useless at the moment
			return nil, true
		})
	}

	return logger, closer
}

// private

func configureWriter(cfg Configuration) (io.Writer, func()) {
	var writer io.Writer = os.Stdout
	if cfg.LogPrettify {
		writer = Writer{
			Out:    writer,
			Indent: true,
		}
	}

	if cfg.ConsoleWriter {
		writer = zerolog.ConsoleWriter{
			Out: writer,
		}
	}

	var err error
	var logFile *os.File
	if cfg.LogToFile && cfg.LogFileName != "" {
		logFile, err = openLogFile(cfg.LogFileName)
		if err != nil {
			log.Fatal().Err(err).Msg("opening log file failed")
		}

		writer = io.MultiWriter(logFile, writer)
	}

	if cfg.MaskSensitiveFields {
		writer = Writer{
			Out:          writer,
			ProcessEvent: maskEvent,
		}
	}

	if cfg.XMLMaskSensitiveFields {
		writer = Writer{
			Out:          writer,
			ProcessEvent: xmlMaskEvent,
		}
	}

	closer := func() {
		if logFile != nil {
			log.Debug().Msgf("Closing log file %s...", logFile.Name())
			if closeErr := logFile.Close(); closeErr != nil {
				log.Error().Err(closeErr).Msg("closing log file failed")
			}
		}

		if rollbar.Token() != "" {
			rollbar.Close()
		}
	}

	return writer, closer
}

func openLogFile(filePath string) (*os.File, error) {
	directory := filepath.Dir(filePath)
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return nil, errors.Annotatef(err, "creating directory failed: directory=%s", directory)
	}

	logFile, err := os.OpenFile(filepath.Clean(filePath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return nil, errors.Annotatef(err, "opening log file failed: filePath=%s", filePath)
	}

	return logFile, nil
}
