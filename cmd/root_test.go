package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCmd(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{})
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, string(out))
	assert.Equal(t, etherpadApiKey, "")
	assert.Equal(t, etherpadUrl, "http://localhost:9001")
	assert.Equal(t, cmd.Long, strings.TrimRight(string(out), "\n"))
}

func TestNewRootCmdArgs(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"--etherpad.apikey", "1"})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, etherpadApiKey, "1")
	assert.Equal(t, etherpadUrl, "http://localhost:9001")
}

func TestNewRootCmdEnv(t *testing.T) {
	_ = os.Setenv("ETHERPAD_APIKEY", "1")

	cmd := NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, etherpadApiKey, "1")
	assert.Equal(t, etherpadUrl, "http://localhost:9001")

	_ = os.Unsetenv("ETHERPAD_APIKEY")
}
