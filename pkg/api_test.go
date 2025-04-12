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

func TestEtherpad_ListSavedRevisions_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": {"savedRevisions": [2,42,1337]}}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	revs, err := etherpad.ListSavedRevisions("pad")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(revs))
}

func TestEtherpad_ListSavedRevisions_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	revs, err := etherpad.ListSavedRevisions("pad")
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(revs))
}

func TestEtherpad_GetText_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": {"text": "Hello World"}}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	text, err := etherpad.GetText("pad")
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", text)
}

func TestEtherpad_GetText_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	text, err := etherpad.GetText("pad")
	assert.NotNil(t, err)
	assert.Equal(t, "", text)
}

func TestEtherpad_RestoreRevision_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.RestoreRevision("pad", "1")
	assert.Nil(t, err)
}

func TestEtherpad_RestoreRevision_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.RestoreRevision("pad", "1")
	assert.NotNil(t, err)
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

func TestEtherpad_CopyPad_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.CopyPad("pad1", "pad2", false)
	assert.Nil(t, err)
}

func TestEtherpad_CopyPad_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	err := etherpad.CopyPad("pad1", "pad2", false)
	assert.NotNil(t, err)
}

func TestEtherpad_GetRevisionsCount_Successful(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 0, "message":"ok", "data": {"revisions": 31}}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	rev, err := etherpad.GetRevisionsCount("pad")
	assert.Nil(t, err)
	assert.Equal(t, 31, rev)
}

func TestEtherpad_GetRevisionsCount_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": 1, "message":"padID does not exist", "data": null}`))
	}))
	defer ts.Close()

	etherpad := NewEtherpadClient(ts.URL, etherpadApiKey)
	etherpad.Client = ts.Client()

	rev, err := etherpad.GetRevisionsCount("pad")
	assert.NotNil(t, err)
	assert.Equal(t, 0, rev)
}
