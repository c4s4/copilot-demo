package main

import (
	"testing"
)

func TestSyncedBool(t *testing.T) {
	var sb SyncedBool
	if sb.IsSet() {
		t.Error("Expected false")
	}
	sb.Set(true)
	if !sb.IsSet() {
		t.Error("Expected true")
	}
}

func TestSyncedString(t *testing.T) {
	var ss SyncedString
	if ss.Get() != "" {
		t.Error("Expected empty string")
	}
	ss.Set("foo")
	if ss.Get() != "foo" {
		t.Error("Expected foo")
	}
	ss.Append("bar")
	if ss.Get() != "foobar" {
		t.Error("Expected foobar")
	}
}
