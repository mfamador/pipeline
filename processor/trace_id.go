// Package processor for benthos processors
package processor

import (
	"errors"
	"time"

	"github.com/uber/jaeger-client-go"

	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/types"
	"github.com/opentracing/opentracing-go"
)

//------------------------------------------------------------------------------

func init() {
	processor.RegisterPlugin(
		"trace_id",
		func() interface{} {
			return NewTraceIDConfig()
		},
		func(
			iconf interface{},
			mgr types.Manager,
			log log.Modular,
			stats metrics.Type,
		) (types.Processor, error) {
			conf, ok := iconf.(*TraceIDConfig)
			if !ok {
				return nil, errors.New("failed to cast config")
			}
			return NewTraceID(conf, log, stats)
		},
	)
}

// TraceIDConfig contains configuration fields for the TraceID processor.
type TraceIDConfig struct {
	MetadataKey string `json:"metadata_key" yaml:"metadata_key"`
}

// NewTraceIDConfig returns a TraceIDConfig with default values.
func NewTraceIDConfig() *TraceIDConfig {
	return &TraceIDConfig{
		MetadataKey: "trace_id",
	}
}

//------------------------------------------------------------------------------

// TraceID is a processor retrieves an TraceID token
type TraceID struct {
	parts []int

	conf  *TraceIDConfig
	log   log.Modular
	stats metrics.Type
}

// NewTraceID returns a TraceID processor.
func NewTraceID(cnfg *TraceIDConfig, logger log.Modular, stats metrics.Type) (types.Processor, error) {
	a := &TraceID{
		conf:  cnfg,
		log:   logger,
		stats: stats,
	}
	return a, nil
}

// ProcessMessage applies the processor to a message, either creating >0
// resulting messages or a response to be sent back to the message source.
func (o *TraceID) ProcessMessage(msg types.Message) ([]types.Message, types.Response) {
	newMsg := msg.Copy()

	proc := func(index int, span opentracing.Span, part types.Part) error {
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			t := sc.TraceID()
			part.Metadata().Set(o.conf.MetadataKey, t.String())
		}
		return nil
	}

	processor.IteratePartsWithSpan("trace_id", o.parts, newMsg, proc)

	msgs := [1]types.Message{newMsg}
	return msgs[:], nil
}

// CloseAsync shuts down the processor and stops processing requests.
func (o *TraceID) CloseAsync() {
}

// WaitForClose blocks until the processor has closed down.
func (o *TraceID) WaitForClose(_ time.Duration) error {
	return nil
}
