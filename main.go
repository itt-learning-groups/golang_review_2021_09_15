package main

import (
	"github.com/itt-learning-groups/golang_review_2021_09_15/ex2/bad"
	"log"
)

func main() {
	// Ex. 2: Separating generator/consumer concerns; Practicing clear channel ownership; Avoiding mutability issues; Identifyig Goroutine leaks

	//   First, we try separating generator/consumer concerns, but without clear channel ownership, leading to problems if
	//   we e.g. uncomment lines 22 or 66 in `ex2/bad/resultsCollector_final.go`
	log.Printf("combined results: %v", bad.Run())

	//   Next, we solve the channel ownership and mutability problems
	//log.Printf("combined results: %v", good.Run())

	//   Finally, we ensure we're protected against goroutine leaks by using the "done" channel pattern
	//log.Printf("combined results: %v", good.RunWithCtx(context.Background()))

	// *************************************************************************

	// Ex. 4: Error handling

	// Try updating the ex4 starter code to doSomethingInteresting if we get an Interesting error from one of the results generators
	//log.Printf("combined results: %v", ex4.RunWithCtx(context.Background()))
}
