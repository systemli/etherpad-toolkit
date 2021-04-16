package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsePadExpiration(t *testing.T) {
	s := "default:720h,temp:24h,keep:262800h"
	exp, err := ParsePadExpiration(s)
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(2592000000000000), exp[DefaultSuffix])
	assert.Equal(t, time.Duration(946080000000000000), exp["keep"])
	assert.Equal(t, time.Duration(86400000000000), exp["temp"])

	s = "wrong:1d"

	_, err = ParsePadExpiration(s)
	assert.Error(t, err)
	assert.Equal(t, "missing default expiration duration", err.Error())

	s = "wrong:1h:2h"
	_, err = ParsePadExpiration(s)
	assert.Error(t, err)
	assert.Equal(t, "missing default expiration duration", err.Error())

	s = ""
	_, err = ParsePadExpiration(s)
	assert.Error(t, err)
	assert.Equal(t, "input string is empty", err.Error())
}

func TestPadExpiration_GetDuration(t *testing.T) {
	s := "default:24h"
	exp, err := ParsePadExpiration(s)
	assert.Nil(t, err)

	dur := exp.GetDuration("pad")
	assert.Equal(t, "-24h0m0s", dur.String())

	s = "default:24h,temp:10m"
	exp, err = ParsePadExpiration(s)
	assert.Nil(t, err)

	dur = exp.GetDuration("pad")
	assert.Equal(t, "-24h0m0s", dur.String())

	dur = exp.GetDuration("pad-temp")
	assert.Equal(t, "-10m0s", dur.String())
}

func TestGroupPadsByExpiration(t *testing.T) {
	s := "default:720h,temp:24h,keep:262800h"
	pads := []string{"pad", "pad2", "pad-keep", "pad-temp"}
	exp, err := ParsePadExpiration(s)
	assert.Nil(t, err)

	sorted := GroupPadsByExpiration(pads, exp)

	if _, ok := sorted[DefaultSuffix]; !ok {
		t.Fail()
	}

	if _, ok := sorted["keep"]; !ok {
		t.Fail()
	}

	if _, ok := sorted["temp"]; !ok {
		t.Fail()
	}

	assert.Equal(t, []string{"pad", "pad2"}, sorted[DefaultSuffix])
	assert.Equal(t, []string{"pad-keep"}, sorted["keep"])
	assert.Equal(t, []string{"pad-temp"}, sorted["temp"])
}
