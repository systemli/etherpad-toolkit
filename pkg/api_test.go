package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	etherpadApiKey = "1"
)

func TestEtherpad_ListAllPads_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": {"padIDs": ["pad1", "pad2"]}}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	pads, err := etherpad.ListAllPads()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(pads))
}

func TestEtherpad_ListAllPads_WrongApiKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":4,"message":"no or wrong API Key","data":null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	pads, err := etherpad.ListAllPads()
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(pads))
}

func TestEtherpad_GetLastEdited_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": {"lastEdited": 1340815946602}}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	edited, err := etherpad.GetLastEdited("pad")
	assert.Nil(t, err)
	assert.Equal(t, int64(1340815946), edited.Unix())
}

func TestEtherpad_GetLastEdited_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	edited, err := etherpad.GetLastEdited("pad")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), edited.Unix())
}

func TestEtherpad_DeletePad_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.DeletePad("pad")
	assert.Nil(t, err)
}

func TestEtherpad_DeletePad_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.DeletePad("pad")
	assert.NotNil(t, err)
}

func TestEtherpad_MovePad_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.MovePad("pad1", "pad2", false)
	assert.Nil(t, err)
}

func TestEtherpad_MovePad_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.MovePad("pad1", "pad2", false)
	assert.NotNil(t, err)
}
