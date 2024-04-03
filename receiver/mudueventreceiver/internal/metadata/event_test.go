package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	event := Event{
		ServiceName:      "test-service",
		StartTime:        1629876543,
		ExceptionType:    "error",
		ExceptionMessage: "something went wrong",
		Tags: []Tag{
			{	Key:   "key1",
				Value: "value1",
				Type:  "string",
			},
			{	Key:	"key2",
				Value: 2,
				Type:  "int",
			},
		},
	}

	data, err := json.Marshal(event)
	assert.NoError(t, err)

	t.Logf("Event: %s", string(data))

	var decodedEvent Event
	err = json.Unmarshal(data, &decodedEvent)
	assert.NoError(t, err)

	data2, err := json.Marshal(decodedEvent)
	assert.NoError(t, err)
	assert.Equal(t, string(data), string(data2))

	eventJson := "{\"serviceName\":\"mdf-stream-utils\",\"startTime\":1711087357942616,\"exceptionType\":\"流媒体相关报警\",\"exceptionMessage\":\"上传失败:ConnectionFailed\",\"tags\":[{\"key\":\"host.ip\",\"value\":\"192.168.3.1\",\"type\":\"string\"}]}"
	var eventFromJson Event
	err = json.Unmarshal([]byte(eventJson), &eventFromJson)
	assert.NoError(t, err)

	data, err = json.Marshal(eventFromJson)
	assert.NoError(t, err)

	t.Logf("Event: %s", string(data))
}
