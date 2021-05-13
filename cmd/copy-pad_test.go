package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCopyPadCmd(t *testing.T) {
	cmd := NewCopyPadCmd()
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

	assert.Equal(t, cmd.UsageString(), string(out))

	cmd.SetArgs([]string{"pad1", "pad2"})
	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Empty(t, string(out))
}
