package packet

import (
	"encoding/json"
)

type Resp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(code, message string, data interface{}) Resp {
	return Resp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (o *Resp) MarshalJson() (string, error) {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (o *Resp) UnmarshalJson(data string) error {
	return json.Unmarshal([]byte(data), o)
}
