package crdt

import "time"

type Item struct {
	value     interface{}
	timestamp time.Time
}

func (i Item) Get() Item {
	return Item{i.value, i.timestamp}
}

func (i Item) NewItem(value interface{}, timestamp time.Time) Item {
	return Item{value, timestamp}
}
