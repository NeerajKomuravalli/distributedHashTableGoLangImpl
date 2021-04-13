package model

import "sync"

type Hashtable struct {
	Item map[string]int
	Lock sync.RWMutex
}
