package crdt

import (
	"testing"
	"time"

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
	d.Add("hello", "new world")
	element, _ = d.Lookup("hello")
	element2, _ := d.Lookup("foo")
	assert.Equal(t, "new world", (*element).GetValue(), "Should be the updated value")
	assert.Equal(t, "bar", (*element2).GetValue(), "Should be the same value")

	// custom adding an older elment
	dt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	d.AddAt("foo", "old bar", dt)
	element, _ = d.Lookup("foo")
	assert.Equal(t, "bar", (*element).GetValue(), "Should be the newer value")

	// custom adding a newer elment
	d.AddAt("foo", "new bar", time.Now().Add(1000))
	d.AddAt("foo", "old bar", time.Now())
	element, _ = d.Lookup("foo")
	assert.Equal(t, "new bar", (*element).GetValue(), "Should be the newer value")
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

	// test new value exist on dictionary 2
	d.Add("hello", "world")
	d2.Add("hello", "new world")

	// test new value exist on dictionary 1
	d2.Add("foo", "bar")
	d.Add("foo", "new bar")

	// test value only exist on dictionary 1
	d.Add("clean", "code")

	// test value only exist on dictionary 2
	d2.Add("cleaner", "architecture")

	d3, _ := d.Merge(*d2)

	element, _ := d3.Lookup("hello")
	assert.Equal(t, "new world", (*element).GetValue(), "Should be the new value from dictionary 2")

	element2, _ := d3.Lookup("foo")
	assert.Equal(t, "new bar", (*element2).GetValue(), "Should be the new value from dictionary 1")

	element3, _ := d3.Lookup("clean")
	assert.Equal(t, "code", (*element3).GetValue(), "Should be the only value from dictionary 1")

	element4, _ := d3.Lookup("cleaner")
	assert.Equal(t, "architecture", (*element4).GetValue(), "Should be the only value from dictionary 2")

}
