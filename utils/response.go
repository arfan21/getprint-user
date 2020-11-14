package utils

import (
	"encoding/json"
	"fmt"
)

type ResponseStruct struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(status, message string, data interface{}) ResponseStruct {
	dataType := fmt.Sprintf("%T", data)

	if dataType == "error" {
		dataJSON, _ := json.Marshal(data)
		return ResponseStruct{status, message, dataJSON}
	}

	return ResponseStruct{status, message, data}
}
