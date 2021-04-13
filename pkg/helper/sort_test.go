package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPads(t *testing.T) {
	pads := []string{"pad", "pad-keep", "pad-temp"}
	sorted := SortPads(pads)

	assert.Equal(t, []string{"pad"}, sorted["none"])
	assert.Equal(t, []string{"pad-keep"}, sorted["keep"])
	assert.Equal(t, []string{"pad-temp"}, sorted["temp"])
}
