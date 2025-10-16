package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func PerformHttpRequest(url string, token string, httpMethod string, reqBody []byte) ([]byte, error) {
	req, _ := http.NewRequest(httpMethod, url, bytes.NewBuffer(reqBody))

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	return body, nil
}
