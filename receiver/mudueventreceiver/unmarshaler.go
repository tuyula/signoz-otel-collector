// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver // import "github.com/SigNoz/signoz-otel-collector/receiver/signozkafkareceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type TracesUnmarshaler interface {
	// Unmarshal deserializes the message body into traces.
	Unmarshal([]byte) (ptrace.Traces, error)

	// Encoding of the serialized messages.
	Encoding() string
}


// LogsUnmarshaler deserializes the message body.
type LogsUnmarshaler interface {
	// Unmarshal deserializes the message body into traces.
	Unmarshal([]byte) (plog.Logs, error)

	// Encoding of the serialized messages.
	Encoding() string
}

type LogsUnmarshalerWithEnc interface {
	LogsUnmarshaler

	// WithEnc sets the character encoding (UTF-8, GBK, etc.) of the unmarshaler.
	WithEnc(string) (LogsUnmarshalerWithEnc, error)
}


func defaultLogsUnmarshalers() map[string]LogsUnmarshaler {
	mudu := newMuduEventUnmarshaler()
	return map[string]LogsUnmarshaler{
		mudu.Encoding():   mudu,

	}
}


func defaultTracesUnmarshalers() map[string]TracesUnmarshaler {
	mudu := newMuduTraceEventUnmarshaler()

	return map[string]TracesUnmarshaler{
		mudu.Encoding():       mudu,
	}
}