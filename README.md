# Konc

> **Concurrency, reimagined for Kopexa.**

Konc is an internal concurrency library, built on the robust foundations of [Pond](https://github.com/alitto/pond), and meticulously adapted for the unique needs of Kopexa.

## Why Konc?

- **Battle-tested Foundations:** Built on Pond, Konc brings modern, safe, and efficient worker pool patterns to the Kopexa ecosystem.
- **Tailored for Kopexa:** Enhanced and fine-tuned to meet the demands of Kopexa's infrastructure, with a focus on reliability, maintainability, and developer happiness.
- **Future-proof:** Designed to evolve with Kopexa's needs, leveraging the latest Go features and best practices.

## Features

- 🏗️ **Worker Pool Management:** Efficient goroutine pooling with configurable limits.
- 🛡️ **Safe Concurrency:** Built-in safeguards against goroutine leaks and resource exhaustion.
- 🚦 **Flexible Execution:** Support for both individual task submission and batch execution with wait semantics.
- 📈 **Performance:** Minimal overhead, maximum throughput through intelligent worker reuse.
- 🔒 **Graceful Shutdown:** Clean pool termination with optional wait for completion.

## Usage

> **Note:** Konc is an internal library. Usage outside Kopexa is not supported.

Add Konc to your Go module:

```sh
go get github.com/kopexa-grc/konc
```

### Basic Usage

```go
import "github.com/kopexa-grc/konc"

// Create a pool with default settings (1000 max goroutines)
pool := konc.NewPool()

// Submit a single task
pool.Go(func() {
    // Your concurrent work here
    fmt.Println("Task executed")
})

// Submit multiple tasks and wait for completion
pool.GoMultipleAndWait(
    func() { fmt.Println("Task 1") },
    func() { fmt.Println("Task 2") },
    func() { fmt.Println("Task 3") },
)

// Gracefully shutdown the pool
pool.StopAndWait()
```

### Advanced Configuration

```go
// Create a pool with custom settings
pool := konc.NewPool(
    konc.WithMaxGoroutines(500),
)

// Use the pool for concurrent operations
pool.Go(func() {
    // Heavy computation
})

// Stop without waiting for completion
pool.Stop()
```

## API Reference

### Pool

```go
// Pool provides a managed worker pool for executing concurrent tasks.
type Pool struct {
    MaxGoroutines int // Maximum number of concurrent goroutines
}

// NewPool creates a new Pool instance with the given options.
func NewPool(opts ...PoolOptions) *Pool

// Go submits a function to be executed by the pool.
func (p *Pool) Go(f func())

// GoMultipleAndWait submits multiple functions and waits for all to complete.
func (p *Pool) GoMultipleAndWait(f ...func())

// StopAndWait stops the pool and waits for all running tasks to complete.
func (p *Pool) StopAndWait()

// Stop stops the pool without waiting for running tasks to complete.
func (p *Pool) Stop()
```

### PoolOptions

```go
// PoolOptions configures a Pool before creation.
type PoolOptions func(*Pool)

// WithMaxGoroutines sets the maximum number of concurrent goroutines.
func WithMaxGoroutines(maxGoroutines int) PoolOptions
```

## License

Konc is licensed under the [Business Source License 1.1 (BUSL-1.1)](./LICENSE) by Kopexa GmbH.

- **Non-commercial use only.**
- For commercial licensing, contact [hello@kopexa.com](mailto:hello@kopexa.com).

## Contributing

This repository is maintained by the Kopexa team. External contributions are not accepted at this time.

## Security

If you discover a security vulnerability, please report it to [security@kopexa.com](mailto:security@kopexa.com) as described in [SECURITY.md](./SECURITY.md).

---

> _Konc: Concurrency, the Kopexa way._
