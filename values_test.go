package values

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNew(t *testing.T) {
	values := New()
	assert.Equal(t, values.Get(""), "")
}

func TestFetch(t *testing.T) {
	values := New()
	values.Set("golang", "great")
	values.Set("node", "js")
	assert.Equal(t, values.Fetch("node", "io"), "js")
	assert.Equal(t, values.Fetch("ruby", "rails"), "rails")
	assert.Equal(t, values.Fetch("hello", "world"), "world")
}
