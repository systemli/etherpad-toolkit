package cmd

import (
	"bytes"
	"io/ioutil"
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
	assert.Equal(t, cmd.Long, strings.TrimRight(string(out), "\n"))
}
