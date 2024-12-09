package log

import (
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc/grpclog"
)

type GRPCLogger struct {
	v int
}

var _ grpclog.LoggerV2 = (*GRPCLogger)(nil)

func NewGRPCLogger() *GRPCLogger {
	var v int
	vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
	if vl, err := strconv.Atoi(vLevel); err == nil {
		v = vl
	}

	return &GRPCLogger{
		v: v,
	}
}

// Implement grpclog.LoggerV2

func (g *GRPCLogger) Info(args ...any)                 { g.Debug(args...) }
func (g *GRPCLogger) Infoln(args ...any)               { g.Debugln(args...) }
func (g *GRPCLogger) Infof(format string, args ...any) { g.Debugf(format, args...) }

func (g *GRPCLogger) Warning(args ...any)                 { g.Debug(args...) }
func (g *GRPCLogger) Warningln(args ...any)               { g.Debugln(args...) }
func (g *GRPCLogger) Warningf(format string, args ...any) { g.Debugf(format, args...) }

func (g *GRPCLogger) Error(args ...any)                 { g.Debug(args...) }
func (g *GRPCLogger) Errorln(args ...any)               { g.Debugln(args...) }
func (g *GRPCLogger) Errorf(format string, args ...any) { g.Debugf(format, args...) }

func (g *GRPCLogger) Fatal(args ...any)                 { g.Debug(args...) }
func (g *GRPCLogger) Fatalln(args ...any)               { g.Debugln(args...) }
func (g *GRPCLogger) Fatalf(format string, args ...any) { g.Debugf(format, args...) }

func (g *GRPCLogger) V(l int) bool { return l <= g.v }

// Send all logs to Debug

func (g *GRPCLogger) Debug(args ...any)                 { Debug(fmt.Sprint(args...)) }
func (g *GRPCLogger) Debugln(args ...any)               { Debug(fmt.Sprintln(args...)) }
func (g *GRPCLogger) Debugf(format string, args ...any) { Debug(fmt.Sprintf(format, args...)) }
