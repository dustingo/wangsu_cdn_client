package util

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Format 格式化返回的字符串"{\"Code\":1,\"Message\":\"handle success\",\"itemId\":\"7109734569569527004\"}"
func Format(response string) string {
	newresp := strings.ReplaceAll(response, "\\", "")
	var out bytes.Buffer
	json.Indent(&out, []byte(newresp), "", "\t")
	return out.String()

}
