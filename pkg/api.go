package pkg

import (
	"encoding/json"
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
		return nil, fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
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
		return time.Unix(0, 0), fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return time.Unix(body.Data.LastEdited/1000, 0), nil
}

// ListSavedRevisions returns the list of saved revisions of a Pad.
// See: https://etherpad.org/doc/v1.8.4/#index_listsavedrevisions_padid
func (ep *Etherpad) ListSavedRevisions(padID string) ([]int, error) {
	params := map[string]interface{}{"padID": padID}
	res, err := ep.sendRequest("listSavedRevisions", params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			SavedRevisions []int `json:"savedRevisions"`
		} `json:"data"`
	}
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}
	if body.Code != 0 {
		return nil, fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return body.Data.SavedRevisions, nil
}

// GetText returns the text of a Pad.
// See: https://etherpad.org/doc/v1.8.4/index.html#index_gettext_padid_rev
func (ep *Etherpad) GetText(padID string) (string, error) {
	params := map[string]interface{}{"padID": padID}
	res, err := ep.sendRequest("getText", params)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Text string `json:"text"`
		} `json:"data"`
	}
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return "", err
	}

	if body.Code != 0 {
		return "", fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return body.Data.Text, nil
}

// RestoreRevision restores a Pad to a specific revision.
// See: https://etherpad.org/doc/v1.8.4/#index_restorerevision_padid_rev
func (ep *Etherpad) RestoreRevision(padID, rev string) error {
	params := map[string]interface{}{"padID": padID, "rev": rev}
	res, err := ep.sendRequest("restoreRevision", params)
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
		return fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return nil
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
		return fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return nil
}

// MovePad moves a pad. If force is true and the destination pad exists, it will be overwritten.
// See: https://etherpad.org/doc/v1.8.4/#index_movepad_sourceid_destinationid_force_false
func (ep *Etherpad) MovePad(sourceID, destinationID string, force bool) error {
	params := map[string]interface{}{"sourceID": sourceID, "destinationID": destinationID, "force": force}
	res, err := ep.sendRequest("movePad", params)
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
		return fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return nil
}

// CopyPad copies a pad with full history and chat. If force is true and the destination pad exists, it will be overwritten.
// See: https://etherpad.org/doc/v1.8.4/#index_copypad_sourceid_destinationid_force_false
func (ep *Etherpad) CopyPad(sourceID, destinationID string, force bool) error {
	params := map[string]interface{}{"sourceID": sourceID, "destinationID": destinationID, "force": force}
	res, err := ep.sendRequest("copyPad", params)
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
		return fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return nil
}

// GetRevisionsCount returns the number of revisions of this pad.
// See: https://etherpad.org/doc/v1.8.4/#index_getrevisionscount_padid
func (ep *Etherpad) GetRevisionsCount(padID string) (int, error) {
	params := map[string]interface{}{"padID": padID}
	res, err := ep.sendRequest("getRevisionsCount", params)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Revisions int `json:"revisions"`
		}
	}
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return 0, err
	}

	if body.Code != 0 {
		return 0, fmt.Errorf("error: %s (code: %d)", body.Message, body.Code)
	}

	return body.Data.Revisions, nil
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

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	return ep.Client.Do(req)
}
