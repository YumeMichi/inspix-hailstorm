package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const APP_ID = "1665027261"

type Resp struct {
	Results []struct {
		Version string `json:"version"`
	} `json:"results"`
}

func GetAppVersion() (string, error) {
	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%s&lang=ja_jp&country=jp&rnd=%d", APP_ID, time.Now().Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	header := http.Header{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0"},
	}
	req.Header = header
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("Abnormal HTTP status code: %d. Message: %s", res.StatusCode, res.Status)
	}
	var resp Resp
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	return resp.Results[0].Version, nil
}
