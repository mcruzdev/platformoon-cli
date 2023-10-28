package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}) (*http.Response, error) {

	j, err := json.Marshal(body)
	if err != nil {
		panic(any(err))
	}

	b := []byte(j)

	return http.Post(url, "application/json", bytes.NewBuffer(b))

}
