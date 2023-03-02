# CRDT

## Concepts

### Element

Basic unit of each element in CRDT, contains a text value and a timestamp

### Dictionary

Contains 2 sets of data, add set and remove set, each set is a dictionary mapping a key to an element

## Usage

To create new dictionary:

`dict := NewDictionary()`

Add new item to the dictionary:

`error := dict.Add("key", "value")`

Rmove item from the dictionary:

`error := dict.Remove("key")`

Lookup item from the dictionary:

`element, err := dict.Lookup("key")`

To merge from another diciontary:

`newDict, error := dict.Merge(dict2)`
