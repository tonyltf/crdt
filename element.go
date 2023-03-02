package crdt

import "time"

type Element struct {
	value     string
	timestamp time.Time
}

func (i Element) NewItem(value string, timestamp time.Time) *Element {
	return &Element{value, timestamp}
}

func (i Element) GetValue() string {
	return i.value
}

func (i Element) GetTime() time.Time {
	return i.timestamp
}
