package main

import "sync"

// SyncedBool is synchronized boolean
type SyncedBool struct {
	mutex   sync.RWMutex
	boolean bool
}

// Set boolean value
func (sb *SyncedBool) Set(value bool) {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()
	sb.boolean = value
}

// IsSet boolean value
func (sb *SyncedBool) IsSet() bool {
	sb.mutex.RLock()
	defer sb.mutex.RUnlock()
	return sb.boolean
}
