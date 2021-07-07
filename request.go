package go_basic_function

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Request(method string, requestUrl string, header map[string]string, bodyData string) (statusCode int, rspBody []byte, err error, rspHeader http.Header) {
	//var form = url.Values{"key":{"value"}, "name":{"value"}}
	//var bodyData = form.Encode()
	//var header = map[string]string{"content-type": "application/x-www-form-urlencoded"}

	req, err := http.NewRequest(method, requestUrl, bytes.NewBufferString(bodyData))
	if err != nil {
		return
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	var resp *http.Response

	client := http.DefaultClient
	resp, err = client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	statusCode = resp.StatusCode
	rspHeader = resp.Header
	rspBody, err = ioutil.ReadAll(resp.Body)
	return
}
