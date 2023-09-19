package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func Request(url string, method string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ja,en-US;q=0.9,en;q=0.8,ja-JP;q=0.7")
	req.Header.Add("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}

func RequestJSON(data any, url string, method string, body io.Reader) error {
	res, err := Request(url, method, body)
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	result := string(bodyBytes)
	json.Unmarshal([]byte(result), &data)
	return nil
}
