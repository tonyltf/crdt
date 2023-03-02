package crdt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Add one values
	d := NewDictionary()
	d.Add("hello", "world")
	element, _ := d.Lookup("hello")
	assert.Equal(t, "world", (*element).GetValue(), "Should be the same value")

	// Add multiple values
	d.Add("foo", "bar")
	element, _ = d.Lookup("hello")
	element2, _ := d.Lookup("foo")
	assert.Equal(t, "world", (*element).GetValue(), "Should be the same value")
	assert.Equal(t, "bar", (*element2).GetValue(), "Should be the same value")
}

func TestAddAndRemove(t *testing.T) {
	d := NewDictionary()
	d.Add("hello", "world")
	d.Rmove("hello")
	element, _ := d.Lookup("hello")
	assert.Nil(t, element, "Should be nil value")

	d.Add("hello", "world")
	d.Rmove("hello")
	d.Add("hello", "world")
	element, _ = d.Lookup("hello")
	assert.Equal(t, "world", (*element).GetValue(), "Should be the same value")
}

func TestMerge(t *testing.T) {
	d := NewDictionary()
	d2 := NewDictionary()
	d.Add("hello", "world")
	d2.Add("hello", "new world")
	d3, _ := d.Merge(*d2)
	element, _ := d3.Lookup("hello")
	assert.Equal(t, "new world", (*element).GetValue(), "Should be the new value")
}
