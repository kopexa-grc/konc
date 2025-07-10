// Copyright (c) Kopexa GmbH
// SPDX-License-Identifier: BUSL-1.1

package konc

import (
	"log"
	"sync"

	"github.com/alitto/pond/v2"
)

// Pool provides a managed worker pool for executing concurrent tasks.
// It wraps the pond library to provide a simplified interface for goroutine management.
type Pool struct {
	// pool is the underlying pond worker pool
	pool pond.Pool
	// MaxGoroutines defines the maximum number of concurrent goroutines
	MaxGoroutines int
	// opts are the configuration options for the pool
	opts []pond.Option
}

// NewPool creates a new Pool instance with the given options.
// By default, the pool allows up to 1000 concurrent goroutines.
func NewPool(opts ...PoolOptions) *Pool {
	p := &Pool{
		MaxGoroutines: DefaultMaxGoroutines,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.pool = pond.NewPool(p.MaxGoroutines, p.opts...)

	return p
}

// Go submits a function to be executed by the pool.
// The function will be executed asynchronously by one of the pool's workers.
func (p *Pool) Go(f func()) {
	err := p.pool.Go(f)
	if err != nil {
		log.Println(ErrPoolGoFailed, err)
	}
}

// GoMultipleAndWait submits multiple functions to the pool and waits for all to complete.
// All functions are executed concurrently, and the method blocks until all finish.
func (p *Pool) GoMultipleAndWait(f ...func()) {
	wg := new(sync.WaitGroup)

	for _, t := range f {
		wg.Add(1)

		p.Go(func() {
			defer wg.Done()
			t()
		})
	}

	wg.Wait()
}

// StopAndWait stops the pool and waits for all running tasks to complete.
// No new tasks can be submitted after calling this method.
func (p *Pool) StopAndWait() {
	p.pool.StopAndWait()
}

// Stop stops the pool without waiting for running tasks to complete.
// No new tasks can be submitted after calling this method.
func (p *Pool) Stop() {
	p.pool.Stop()
}

// PoolOptions configures a Pool before creation.
type PoolOptions func(*Pool)

// WithMaxGoroutines sets the maximum number of concurrent goroutines for the pool.
func WithMaxGoroutines(maxGoroutines int) PoolOptions {
	return func(p *Pool) {
		p.MaxGoroutines = maxGoroutines
	}
}

// WithOptions sets the underlying pond library options for the pool.
// This allows advanced configuration of the pond worker pool behavior.
func WithOptions(opts ...pond.Option) PoolOptions {
	return func(p *Pool) {
		p.opts = opts
	}
}
