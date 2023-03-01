package crdt

import "time"

type set interface {
	addItem(key string, value interface{}, time time.Time) error
	removeItem(key string, value interface{}, time time.Time) error
}

type Dictonary struct {
	AddSet    set
	RemoveSet set
}

func (d Dictonary) Add(key string, value interface{}, time time.Time) error {
	return nil
}

func (d Dictonary) Rmove(key string, value interface{}, time time.Time) error {
	return nil
}

func (d Dictonary) Lookup(key string) (*Item, error) {
	return nil, nil
}

func (d Dictonary) Merge(source Dictonary) (*Dictonary, error) {
	return nil, nil
}
