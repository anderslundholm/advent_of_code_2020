package timer

import (
	"fmt"
	"time"
)

// ExecutionTimer calculates the execution time of the application
// Usage: defer executiuonTime()()
func ExecutionTimer(function string) func() {
	startTimer := time.Now()
	return func() {
		fmt.Printf("%s executed in: %v\n", function, time.Since(startTimer))
	}
}
