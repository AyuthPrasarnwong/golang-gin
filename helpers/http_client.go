package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpClient(method string, url string, body io.Reader, headers map[string]string, query string) (statusCode int, result map[string]interface{}) {

	req, err := http.NewRequest(method, url, body)

	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = query

	if err != nil {

	}

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)
	//
	if resp.StatusCode >= 400  {
		fmt.Println("err", resp.StatusCode)
		return resp.StatusCode, result
	}

	if err != nil {
		fmt.Println("err", err)
	}

	//fmt.Println("result", result)

	return resp.StatusCode, result
}
