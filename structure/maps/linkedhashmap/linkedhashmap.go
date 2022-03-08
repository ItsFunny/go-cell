// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashmap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"fmt"
	"github.com/itsfunny/go-cell/structure/lists/doublylinkedlist"
	"math/rand"
	"strings"
	"github.com/emirpasic/gods/maps"
)

func assertMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map struct {
	table    map[interface{}]interface{}
	ordering *doublylinkedlist.List
}

// New instantiates a linked-hash-map.
func New() *Map {
	return &Map{
		table:    make(map[interface{}]interface{}),
		ordering: doublylinkedlist.New(),
	}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Put(key interface{}, value interface{}) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

func (m *Map) PutIfNotExist(key interface{}, value interface{}) bool {
	if _, contains := m.table[key]; contains {
		return true
	}
	m.ordering.Append(key)
	m.table[key] = value
	return false
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value = m.table[key]
	found = value != nil
	return
}
func (m *Map) randomGet() interface{} {
	l := m.ordering.Size()
	if l == 0 {
		return nil
	}
	index := rand.Intn(l)
	get, _ := m.ordering.Get(index)
	return get
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map) Remove(key interface{}) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}
func (m *Map) RemoveWithReturn(key interface{}) (interface{}, bool) {
	v, contains := m.table[key]
	if contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
	return v, contains
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys in-order
func (m *Map) Keys() []interface{} {
	return m.ordering.Values()
}

// Values returns all values in-order based on the key.
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.table = make(map[interface{}]interface{})
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "LinkedHashMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}