package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}) (*http.Response, error) {

	marshal, err := json.Marshal(body)
	if err != nil {
		panic(any(err))
	}

	byteArr := []byte(marshal)

	return http.Post(url, "application/json", bytes.NewBuffer(byteArr))

}
