package metadata

// {
//     "serviceName": "mdf-stream-utils",
//     "startTime": 1711087357942616,
//     "exceptionType": "流媒体相关报警",
//     "exceptionMessage": 上传失败:ConnectionFailed",
//     "tags": [{"key": "host.ip", "value": "192.168.3.1", "type": "string"}]
// }

type Event struct {
	ServiceName      string `json:"serviceName"`
	StartTime        int64  `json:"startTime"`
	ExceptionType    string `json:"exceptionType"`
	ExceptionMessage string `json:"exceptionMessage"`
	Tags             []Tag  `json:"tags"`
}

type Tag struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

// GetStringValue returns the string value of the tag.
func (t Tag) GetStringValue() string {
	return t.Value.(string)
}

// GetIntValue returns the int value of the tag.
func (t Tag) GetIntValue() int64 {
	return t.Value.(int64)
}

// GetBoolValue returns the bool value of the tag.
func (t Tag) GetBoolValue() bool {
	return t.Value.(bool)
}

// GetDoubleValue returns the double value of the tag.
func (t Tag) GetDoubleValue() float64 {
	return t.Value.(float64)
}
