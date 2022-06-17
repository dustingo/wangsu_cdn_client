package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"wangsu_cdn_client/common/model"
)

func Call(requestMsg model.HttpRequestMsg) string {
	client := &http.Client{}
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		return err.Error()
	}
	for key := range requestMsg.Headers {
		req.Header.Set(key, requestMsg.Headers[key])
	}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	fmt.Println(resp)
	return string(body)
}
