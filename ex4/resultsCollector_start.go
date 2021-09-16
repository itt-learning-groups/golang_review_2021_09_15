package ex4

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type result struct{}

func RunWithCtx(ctx context.Context) []result {
	resultsGenerators := []func() ([]result, error) {
		func() ([]result, error) { return []result{{}}, nil },                            // generator1 returns 1 result
		func() ([]result, error) { return []result{{}, {}}, nil },                        // generator2 returns 2 results
		func() ([]result, error) { return []result{}, fmt.Errorf("error") },     // generator3 returns an error
		func() ([]result, error) { return []result{{}, {}, {}, {}}, nil },                // generator4 returns 4 results
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	resultsStream := resultsStreamOwnerWithCtx(ctx, resultsGenerators)
	return resultsStreamConsumer(resultsStream)
}

func resultsStreamOwnerWithCtx(ctx context.Context, resultsGenerators []func() ([]result, error)) <-chan result {
	resultsStream := make(chan result)

	var wg sync.WaitGroup
	wg.Add(len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go generate(ctx, generator, &wg, resultsStream)
	}

	go func() {
		wg.Wait()
		close(resultsStream)
	}()

	return resultsStream
}

func generate(ctx context.Context, getResults func() ([]result, error), wg *sync.WaitGroup, stream chan<- result) {
	defer wg.Done()

	results, err := getResults()

	// Todo: Answer this question: What do we do with this error?
	if err != nil {
		log.Printf(err.Error())
	}

	for _, r := range results {
		select {
		case <-ctx.Done():
			return
		case stream <- r:
		}
	}
}

func resultsStreamConsumer(resultsStream <-chan result) []result {
	var results []result

	for r := range resultsStream {
		results = append(results, r)
	}

	return results
}
