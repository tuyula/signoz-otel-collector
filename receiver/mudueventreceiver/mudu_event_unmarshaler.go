// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver // import "github.com/SigNoz/signoz-otel-collector/receiver/signozkafkareceiver"

import (
	"encoding/json"

	"github.com/SigNoz/signoz-otel-collector/receiver/mudueventreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/plog"
)

type muduEventUnmarshaler struct {
	plog.Unmarshaler
	encoding string
}

func newMuduEventUnmarshaler() LogsUnmarshaler {
	return &muduEventUnmarshaler{}
}

// 将buf转化为protobuf events格式, 然后用hanlder处理
func (p muduEventUnmarshaler) Unmarshal(buf []byte) (plog.Logs, error) {

	var event metadata.Event

	err := json.Unmarshal(buf, &event)
	if err != nil {
		return plog.Logs{}, err
	}

	// 将protobuf格式的event转化为plog.Logs
	logs := k8sEventToLogData(&event)
	return logs, nil
}

func (p muduEventUnmarshaler) Encoding() string {
	p.encoding = "mudu-log"
	return p.encoding
}

type muduEventTraceUnmarshaler struct {
	ptrace.Unmarshaler
	encoding string
}

func newMuduTraceEventUnmarshaler() TracesUnmarshaler {
	return &muduEventTraceUnmarshaler{}
}

// 将buf转化为protobuf events格式, 然后用hanlder处理
func (p muduEventTraceUnmarshaler) Unmarshal(buf []byte) (ptrace.Traces, error) {

	var event metadata.Event

	err := json.Unmarshal(buf, &event)
	if err != nil {
		return ptrace.Traces{}, err
	}

	// 将protobuf格式的event转化为plog.Logs
	traces := k8sEventToTraceData(&event)
	return traces, nil
}

func (p muduEventTraceUnmarshaler) Encoding() string {
	p.encoding = "mudu-trace"
	return p.encoding
}