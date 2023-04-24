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

// SyncedString is synchronized string
type SyncedString struct {
	mutex  sync.RWMutex
	string string
}

// Set string value
func (ss *SyncedString) Set(value string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.string = value
}

// Get string value
func (ss *SyncedString) Get() string {
	ss.mutex.RLock()
	defer ss.mutex.RUnlock()
	return ss.string
}

// Append string value
func (ss *SyncedString) Append(value string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.string += value
}
