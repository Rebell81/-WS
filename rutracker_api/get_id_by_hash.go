package rutracker_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type hashAndNum struct {
	Result map[string]*int `json:"result"`
}

func GetIdByHashes(hashes []string, apiKey string) (map[string]*int, error) {
	fullHashSet := make(map[string]*int)

	for _, v := range hashes {
		resp, err := makeReq(v, apiKey)
		if err != nil {
			return nil, err
		}

		hashSet, err := parse(resp)
		if err != nil {
			return nil, err
		}

		for k, num := range hashSet {
			fullHashSet[k] = num
		}
	}

	return fullHashSet, nil
}

func makeReq(hashes, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://api.rutracker.cc/v1/get_topic_id?by=hash&val=%s&api_key=%s", hashes, apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parse(data []byte) (map[string]*int, error) {
	var p hashAndNum
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return p.Result, nil
}
