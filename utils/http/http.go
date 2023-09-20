package http

import (
	"applemango/boorutan/backend/db/redis"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func PushCache(url string, body string) error {
	err := redis.Push(fmt.Sprintf("cache:%s", url), body)
	return err
}

func GetCache(url string) (string, error) {
	body, err := redis.Get(fmt.Sprintf("cache:%s", url))
	return body, err
}

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

type RequestOption struct {
	Data   any
	Url    string
	Method string
	Body   io.Reader
	Cache  bool
}

func RequestJSON(option RequestOption) error {
	if option.Cache {
		cache, err := GetCache(option.Url)
		if err == nil {
			json.Unmarshal([]byte(cache), &option.Data)
			return nil
		}
	}
	res, err := Request(option.Url, option.Method, option.Body)
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	result := string(bodyBytes)
	PushCache(option.Url, result)
	json.Unmarshal([]byte(result), &option.Data)
	return nil
}

func RequestDownloadImage(url string, path string) error {
	res, err := Request(url, "GET", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
