package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
)

func prepareUrl(baseUrl, endpoint string, params url.Values) (*url.URL, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	u.Path = endpoint
	u.RawQuery = params.Encode()
	return u, nil
}

func responseErrorHandler(resp *http.Response, err error) error {
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
    	if err != nil {	
			return err
		}
		bodyString := string(bodyBytes)
		return fmt.Errorf("error: bad status code %d %s", resp.StatusCode, bodyString)
	}

	return nil
}

type FlappyBirdArcadeData struct {
	Data        string `json:"data"`
	SessionHash string `json:"session_hash"`
	Error       bool   `json:"error"`
	Msg         string `json:"msg"`
}

func flappyBirdEncryptedData(gameSessionId string, coins, score int) (*FlappyBirdArcadeData, error) {
	out, err := exec.Command("python", "_bcrypt.py", "-i", gameSessionId, "-c", fmt.Sprint(coins), "-s", fmt.Sprint(score)).Output()
	if err != nil {
		return nil, err
	}

	data := &FlappyBirdArcadeData{}
	err = json.Unmarshal(out, data)
	if err != nil {
		return nil, err
	}

	if data.Error {
		return nil, fmt.Errorf("error: %s", data.Msg)
	
	}

	return data, nil
}
