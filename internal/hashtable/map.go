package hashtable

import "errors"

type Value interface {
	any
}

type HashTable struct {
	m map[string]string
}

func NewHashTable() HashTable {
	return HashTable{
		m: make(map[string]string),
	}
}

func (t *HashTable) Set(key string, val string) {
	t.m[key] = val
}

func (t *HashTable) Get(key string) (string, error) {
	v, ok := t.m[key]
	if ok {
		return v, nil
	} else {
		return "", errors.New("not found")
	}
}

func (t *HashTable) Ttl(key string) {
}
