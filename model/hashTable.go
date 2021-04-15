package model

import "sync"

type Hashtable struct {
	Item map[string]interface{}
	Lock sync.RWMutex
}
