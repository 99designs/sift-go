package sift

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	apiEndpoint = "https://api.siftscience.com/v203/events"
)

func Call(e TypedEvent) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		return nil
	case 400:
		var apiError ApiError
		err = json.Unmarshal(res, &apiError)
		if err != nil {
			return err
		}
		return &apiError
	case 500:
		return errors.New("error 500")
	default:
		return errors.New("received strange status code " + string(resp.StatusCode))
	}
}
