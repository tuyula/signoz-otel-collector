// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"time"

	"github.com/SigNoz/signoz-otel-collector/receiver/mudueventreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	semconv "go.opentelemetry.io/collector/semconv/v1.6.1"
)

const (
	totalEventAttributes = 3
)

// k8sEventToLogRecord converts Kubernetes event to plog.LogRecordSlice and adds the resource attributes.
func k8sEventToTraceData(ev *metadata.Event) ptrace.Traces {
	td := ptrace.NewTraces()
	rs := td.ResourceSpans().AppendEmpty()
	ss := rs.ScopeSpans().AppendEmpty()
	sss := ss.Spans().AppendEmpty()

	sourceAttrs := rs.Resource().Attributes()
	sourceAttrs.EnsureCapacity(1)
	sourceAttrs.PutStr(semconv.AttributeServiceName, ev.ServiceName)

	// Set the trace ID and span ID.
	hexStr := "b05558a6fbbe68f9fc4b9c0588888888"
	spanHexStr := "439af60a88888888"
	traceId, _ := hex.DecodeString(hexStr)
	spanId, _ := hex.DecodeString(spanHexStr)
	sss.SetTraceID(pcommon.TraceID(traceId))
	sss.SetSpanID(pcommon.SpanID(spanId))
	sss.SetKind(ptrace.SpanKindInternal)

	spanAttrs := sss.Attributes()
	spanAttrs.EnsureCapacity(len(ev.Tags))

	sss.SetName(ev.ExceptionType)

	event := sss.Events().AppendEmpty()
	event.SetName("exception")
	event.SetTimestamp(pcommon.NewTimestampFromTime(nanoToTime(ev.StartTime)))
	eventAttrs := event.Attributes()
	eventAttrs.EnsureCapacity(totalEventAttributes)
	eventAttrs.PutStr("exception.message", ev.ExceptionMessage)
	eventAttrs.PutStr("exception.type", ev.ExceptionType)

	sss.SetStartTimestamp(pcommon.NewTimestampFromTime(nanoToTime(ev.StartTime)))
	sss.SetEndTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	sss.Status().SetCode(ptrace.StatusCodeError)

	if ev.Tags != nil {
		for _, v := range ev.Tags {
			switch value := v.Value.(type) {
			case string:
				spanAttrs.PutStr(v.Key, value)
			case int64:
				spanAttrs.PutInt(v.Key, value)
			case bool:
				spanAttrs.PutBool(v.Key, value)
			case float64:
				spanAttrs.PutDouble(v.Key, value)
			default:
				fmt.Println("unsupported type", reflect.TypeOf(v.Value))
			}
		}
	}
	return td
}


// 将纳秒转化为time.time的格式
func nanoToTime(nano int64) time.Time {
	return time.Unix(nano/1e9, nano%1e9)
}
