package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const ApiVersion = "1.2.14"

// Etherpad
type Etherpad struct {
	apiKey     string
	apiVersion string
	url        string
	Client     *http.Client
}

// NewEtherpadClient returns a instance of Etherpad
func NewEtherpadClient(url, apiKey string) *Etherpad {
	return &Etherpad{
		apiVersion: ApiVersion,
		apiKey:     apiKey,
		url:        url,
		Client:     &http.Client{},
	}
}

// ListAllPads returns a list of all pads.
// See: https://etherpad.org/doc/v1.8.4/#index_listallpads
func (ep *Etherpad) ListAllPads() ([]string, error) {
	res, err := ep.sendRequest("listAllPads", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			PadIDs []string `json:"padIDs"`
		} `json:"data"`
	}

	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}

	if body.Code != 0 {
		return nil, errors.New(fmt.Sprintf("error: %s (code: %d)", body.Message, body.Code))
	}

	return body.Data.PadIDs, nil
}

// GetLastEdited returns the time of the last modification of a Pad.
// See: https://etherpad.org/doc/v1.8.4/#index_getlastedited_padid
func (ep *Etherpad) GetLastEdited(padID string) (time.Time, error) {
	params := map[string]interface{}{"padID": padID}
	res, err := ep.sendRequest("getLastEdited", params)
	if err != nil {
		return time.Unix(0, 0), err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			LastEdited int64 `json:"lastEdited"`
		} `json:"data"`
	}
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return time.Unix(0, 0), err
	}

	if body.Code != 0 {
		return time.Unix(0, 0), errors.New(fmt.Sprintf("error: %s (code: %d)", body.Message, body.Code))
	}

	return time.Unix(body.Data.LastEdited/1000, 0), nil
}

// DeletePad removes a Pad.
// See: https://etherpad.org/doc/v1.8.4/#index_deletepad_padid
func (ep *Etherpad) DeletePad(padID string) error {
	params := map[string]interface{}{"padID": padID}
	res, err := ep.sendRequest("deletePad", params)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return err
	}

	if body.Code != 0 {
		return errors.New(fmt.Sprintf("error: %s (code: %d)", body.Message, body.Code))
	}

	return nil
}

func (ep *Etherpad) sendRequest(path string, params map[string]interface{}) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/api/%s/%s", ep.url, ep.apiVersion, path))
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("apikey", ep.apiKey)
	for key, value := range params {
		parameters.Add(key, fmt.Sprintf("%v", value))
	}
	uri.RawQuery = parameters.Encode()

	return ep.Client.Get(uri.String())
}
