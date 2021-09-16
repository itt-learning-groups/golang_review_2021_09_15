package ex4

//import (
//	"context"
//	"errors"
//	"fmt"
//	"sync"
//)
//
//type result struct{}
//
//type reportedResult struct {
//	result
//	err error
//}
//
//type InterestingError struct {}
//
//func (e InterestingError) Error() string {
//	return "interesting error"
//}
//
//func RunWithCtx(ctx context.Context) []result {
//	resultsGenerators := []func() ([]result, error) {
//		func() ([]result, error) { return []result{{}}, nil },                            // generator1 returns 1 result
//		func() ([]result, error) { return []result{{}, {}}, nil },                        // generator2 returns 2 results
//		func() ([]result, error) { return []result{}, fmt.Errorf("error") },     // generator3 returns an error
//		func() ([]result, error) { return []result{{}, {}, {}, {}}, nil },                // generator4 returns 4 results
//	}
//
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	resultsStream := resultsStreamOwnerWithCtx(ctx, resultsGenerators)
//	return resultsStreamConsumer(resultsStream)
//}
//
//func resultsStreamOwnerWithCtx(ctx context.Context, resultsGenerators []func() ([]result, error)) <-chan reportedResult {
//	resultsStream := make(chan reportedResult)
//
//	var wg sync.WaitGroup
//	wg.Add(len(resultsGenerators))
//
//	for _, generator := range resultsGenerators {
//		go func(generateResults func() ([]result, error)) {
//			defer wg.Done()
//
//			results, err := generateResults()
//
//			for _, r := range results {
//				select {
//				case <-ctx.Done():
//					return
//				case resultsStream <- reportedResult{ result: r, err: err}:
//				}
//			}
//		}(generator)
//	}
//
//	go func() {
//		wg.Wait()
//		close(resultsStream)
//	}()
//
//	return resultsStream
//}
//
//func resultsStreamConsumer(resultsStream <-chan reportedResult) []result {
//	var results []result
//
//	for rr := range resultsStream {
//		// check for interesting error
//		if rr.err != nil && errors.Is(rr.err, InterestingError{}) {
//			doSomethingInteresting()
//			continue
//		}
//
//		results = append(results, rr.result)
//	}
//
//	return results
//}
//
//func doSomethingInteresting() {}
