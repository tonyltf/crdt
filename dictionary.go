package crdt

import (
	"time"
)

type Dictonary struct {
	addSet    map[string]Element
	removeSet map[string]Element
}

func NewDictionary() *Dictonary {
	return &Dictonary{
		make(map[string]Element),
		make(map[string]Element),
	}
}

func (d Dictonary) Add(key string, value string) error {
	return d.AddAt(key, value, time.Now())
}

func (d Dictonary) AddAt(key string, value string, time time.Time) error {
	element, exists := d.addSet[key]
	if !exists || exists && time.After(element.GetTime()) {
		d.addSet[key] = *element.NewItem(value, time)
	}
	return nil
}

func (d Dictonary) Rmove(key string) error {
	return d.RemoveAt(key, time.Now())
}

func (d Dictonary) RemoveAt(key string, time time.Time) error {
	element, exists := d.addSet[key]
	if !exists || exists && time.After(element.GetTime()) {
		d.removeSet[key] = *element.NewItem("", time)
	}
	return nil
}

func (d Dictonary) Lookup(key string) (*Element, error) {
	addElement, exists := d.addSet[key]
	if !exists {
		return nil, nil
	}
	removeElement, exists := d.removeSet[key]
	if exists {
		if removeElement.GetTime().After(addElement.GetTime()) {
			return nil, nil
		}
	}
	return &addElement, nil
}

func (d Dictonary) Merge(source Dictonary) (*Dictonary, error) {
	dest := &d
	for key, element := range source.addSet {
		dest.AddAt(key, element.GetValue(), element.GetTime())
	}
	for key, element := range source.removeSet {
		dest.RemoveAt(key, element.GetTime())
	}
	return dest, nil
}
