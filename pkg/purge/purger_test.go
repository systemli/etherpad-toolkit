package purge

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/helper"
)

var pads = map[string]struct {
	Revisions  int
	LastEdited time.Time
}{
	"pad": {
		Revisions:  30,
		LastEdited: time.Now().Add(-1 * time.Hour),
	},
	"pad+empty": {
		Revisions:  0,
		LastEdited: time.Now().Add(-1 * time.Hour),
	},
	"pad+expired": {
		Revisions:  1,
		LastEdited: time.Now().Add(-999 * time.Hour),
	},
}

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.String(), "listAllPads") {
		var body struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				PadIDs []string `json:"padIDs"`
			} `json:"data"`
		}

		var padIDs []string

		for padId, _ := range pads {
			padIDs = append(padIDs, padId)
		}

		body.Code = 0
		body.Message = "ok"
		body.Data.PadIDs = padIDs

		b, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
		return
	}

	if strings.Contains(r.URL.String(), "getRevisionsCount") {
		padID := r.URL.Query().Get("padID")
		revisions := pads[padID].Revisions

		var body struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Revisions int `json:"revisions"`
			}
		}

		body.Code = 0
		body.Message = "ok"
		body.Data.Revisions = revisions

		b, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
		return
	}

	if strings.Contains(r.URL.String(), "getLastEdited") {
		padID := r.URL.Query().Get("padID")
		lastEdited := pads[padID].LastEdited

		var body struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				LastEdited int64 `json:"lastEdited"`
			} `json:"data"`
		}

		body.Code = 0
		body.Message = "ok"
		body.Data.LastEdited = lastEdited.Unix() * 1000

		b, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
		return
	}

	if strings.Contains(r.URL.String(), "deletePad") {
		padID := r.URL.Query().Get("padID")
		delete(pads, padID)

		var body struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}

		body.Code = 0
		body.Message = "ok"

		b, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
		return
	}
})

func TestPurger_PurgePads_DryRun(t *testing.T) {
	rec := httptest.NewServer(handler)
	etherpad := pkg.NewEtherpadClient(rec.URL, "")
	expiration, err := helper.ParsePadExpiration("default:720h")
	if err != nil {
		t.Fail()
	}
	purger := NewPurger(etherpad, expiration, true)

	assert.Equal(t, 3, len(pads))

	purger.PurgePads(1)

	assert.Equal(t, 3, len(pads))
}

func TestPurger_PurgePads(t *testing.T) {
	rec := httptest.NewServer(handler)
	etherpad := pkg.NewEtherpadClient(rec.URL, "")
	expiration, err := helper.ParsePadExpiration("default:720h")
	if err != nil {
		t.Fail()
	}
	purger := NewPurger(etherpad, expiration, false)

	assert.Equal(t, 3, len(pads))

	purger.PurgePads(1)

	assert.Equal(t, 1, len(pads))
}
