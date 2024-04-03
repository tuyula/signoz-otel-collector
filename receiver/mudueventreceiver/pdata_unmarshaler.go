// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver // import "github.com/SigNoz/signoz-otel-collector/receiver/signozkafkareceiver"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

type pdataLogsUnmarshaler struct {
	plog.Unmarshaler
	encoding string
}

func (p pdataLogsUnmarshaler) Unmarshal(buf []byte) (plog.Logs, error) {
	return p.Unmarshaler.UnmarshalLogs(buf)
}

func (p pdataLogsUnmarshaler) Encoding() string {
	return p.encoding
}

func NewPdataLogsUnmarshaler(unmarshaler plog.Unmarshaler, encoding string) LogsUnmarshaler {
	return pdataLogsUnmarshaler{
		Unmarshaler: unmarshaler,
		encoding:    encoding,
	}
}
