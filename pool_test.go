// Copyright 2025 Kopexa GmbH
// SPDX-License-Identifier: BUSL-1.1

package konc

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestPool_BasicAndOptions(t *testing.T) {
	var count int32

	pool := NewPool(WithMaxGoroutines(2))

	// Test Go
	pool.Go(func() {
		atomic.AddInt32(&count, 1)
	})
	pool.Go(func() {
		atomic.AddInt32(&count, 1)
	})
	pool.Go(func() {
		atomic.AddInt32(&count, 1)
	})
	pool.StopAndWait()

	if atomic.LoadInt32(&count) != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func TestPool_GoMultipleAndWait(t *testing.T) {
	var count int32

	pool := NewPool()

	pool.GoMultipleAndWait(
		func() { atomic.AddInt32(&count, 1) },
		func() { atomic.AddInt32(&count, 1) },
		func() { atomic.AddInt32(&count, 1) },
	)
	pool.Stop()

	if atomic.LoadInt32(&count) != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func TestPool_StopAndStopAndWait(t *testing.T) {
	pool := NewPool()
	ch := make(chan struct{})

	pool.Go(func() {
		time.Sleep(10 * time.Millisecond)
		close(ch)
	})
	pool.StopAndWait()

	select {
	case <-ch:
		// ok
	case <-time.After(100 * time.Millisecond):
		t.Error("StopAndWait did not wait for goroutine")
	}

	// After Stop, no new tasks should be accepted (pond silently ignores)
	pool.Stop()
	pool.Go(func() { t.Error("should not run after Stop") })
}

func TestPool_WithOptions(_ *testing.T) {
	pool := NewPool(WithOptions())
	pool.Go(func() {})
	pool.StopAndWait()
}
