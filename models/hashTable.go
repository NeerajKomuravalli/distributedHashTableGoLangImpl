package models

import "sync"

type Hashtable struct {
	Item map[string]byte
	Lock sync.RWMutex
}
