package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func RequestHttp(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http request failed, code=%v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AsyncRequestHttp(method, url string, headers map[string]string, body []byte, respCh chan any) {
	var err error
	var data []byte
	defer func() {
		if err != nil {
			respCh <- err
		} else {
			respCh <- data
		}
		close(respCh)
	}()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("http request failed, code=%v", resp.StatusCode)
		return
	}

	data, err = io.ReadAll(resp.Body)
}
