package hook

import (
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

type Tracing struct{}

func NewTracing() Tracing {
	return Tracing{}
}

func (h Tracing) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()
	span := trace.SpanFromContext(ctx)
	spanID := span.SpanContext().SpanID()
	traceID := span.SpanContext().TraceID()
	// Do not bother getting values if either of the values is not valid.
	if !spanID.IsValid() || !traceID.IsValid() {
		return
	}

	// Add span-id and trace-id fields to the logger
	e.Str(spanIDEvent, spanID.String()).
		Str(traceIDEvent, traceID.String())
}

// private

const (
	// spanIDEvent is a constant which represents string field span-id in zerolog event.
	spanIDEvent = "span-id"
	// traceIDEvent is a constant which represents string field trace-id in zerolog event.
	traceIDEvent = "trace-id"
)
