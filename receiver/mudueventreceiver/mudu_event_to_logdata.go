// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"fmt"
	"reflect"
	"time"

	"github.com/SigNoz/signoz-otel-collector/receiver/mudueventreceiver/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	semconv "go.opentelemetry.io/collector/semconv/v1.6.1"
)

const (
	// Number of log attributes to add to the plog.LogRecordSlice.
	totalLogAttributes = 7

	// Number of resource attributes to add to the plog.ResourceLogs.
	totalResourceAttributes = 1
)

// k8sEventToLogRecord converts Kubernetes event to plog.LogRecordSlice and adds the resource attributes.
func k8sEventToLogData(ev *metadata.Event) plog.Logs {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()
	lr := sl.LogRecords().AppendEmpty()

	resourceAttrs := rl.Resource().Attributes()
	resourceAttrs.EnsureCapacity(totalResourceAttributes)

	resourceAttrs.PutStr(semconv.AttributeServiceName, ev.ServiceName)

	lr.SetTimestamp(pcommon.NewTimestampFromTime(time.Unix(ev.StartTime, 0)))

	// The Message field contains description about the event,
	// which is best suited for the "Body" of the LogRecordSlice.
	lr.Body().SetStr(ev.ExceptionMessage)

	attrs := lr.Attributes()

	// 获取event的tags的长度
	tagsLen := len(ev.Tags)
	attrs.EnsureCapacity(tagsLen)
	if ev.Tags != nil {
		for _, v := range ev.Tags {
			switch value := v.Value.(type) {
			case string:
				attrs.PutStr(v.Key, value)
			case int64:
				attrs.PutInt(v.Key, value)
			case bool:
				attrs.PutBool(v.Key, value)
			case float64:
				attrs.PutDouble(v.Key, value)
			default:
				fmt.Println("unsupported type", reflect.TypeOf(v.Value))
			}
		}
	}
	return ld
}
