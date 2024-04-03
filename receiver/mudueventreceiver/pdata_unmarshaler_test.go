// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mudueventreceiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestNewPdataLogsUnmarshaler(t *testing.T) {
	um := NewPdataLogsUnmarshaler(&plog.ProtoUnmarshaler{}, "test")
	assert.Equal(t, "test", um.Encoding())
}
