package konc

import "errors"

// ErrPoolGoFailed is returned when submitting a task to the pool fails.
var ErrPoolGoFailed = errors.New("failed to submit task to pool")
