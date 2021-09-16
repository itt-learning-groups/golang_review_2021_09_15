package good

//import (
//	"context"
//	"sync"
//)
//
//type result struct{}
//
//func Run() []result {
//	resultsGenerators := []func() []result{
//		func() []result { return []result{{}} },             // generator1 returns 1 result
//		func() []result { return []result{{}, {}} },         // generator2 returns 2 results
//		func() []result { return []result{{}, {}, {}} },     // generator3 returns 3 results
//		func() []result { return []result{{}, {}, {}, {}} }, // generator4 returns 4 results
//	}
//
//	done := make(chan struct{})
//	defer close(done)
//
//	resultsStream := resultsStreamOwner(done, resultsGenerators)
//	return resultsStreamConsumer(resultsStream)
//}
//
//func resultsStreamOwner(done <-chan struct{}, resultsGenerators []func() []result) <-chan result {
//	resultsStream := make(chan result)
//
//	var wg sync.WaitGroup
//	wg.Add(len(resultsGenerators))
//
//	for _, generator := range resultsGenerators {
//		go func(generateResults func() []result) {
//			defer wg.Done()
//
//			for _, r := range generateResults() {
//				select {
//				case <-done:
//					return
//				case resultsStream <- r:
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
//func resultsStreamConsumer(resultsStream <-chan result) []result {
//	var results []result
//
//	for r := range resultsStream {
//		results = append(results, r)
//	}
//
//	return results
//}
//
////*********************************************************************************************
//
//// RunWithCtx uses context as an alternative to the basic "done" channel pattern
//func RunWithCtx(ctx context.Context) []result {
//	resultsGenerators := []func() []result{
//		func() []result { return []result{{}} },             // generator1 returns 1 result
//		func() []result { return []result{{}, {}} },         // generator2 returns 2 results
//		func() []result { return []result{{}, {}, {}} },     // generator3 returns 3 results
//		func() []result { return []result{{}, {}, {}, {}} }, // generator4 returns 4 results
//	}
//
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	resultsStream := resultsStreamOwnerWithCtx(ctx, resultsGenerators)
//	return resultsStreamConsumer(resultsStream)
//}
//
//func resultsStreamOwnerWithCtx(ctx context.Context, resultsGenerators []func() []result) <-chan result {
//	resultsStream := make(chan result)
//
//	var wg sync.WaitGroup
//	wg.Add(len(resultsGenerators))
//
//	for _, generator := range resultsGenerators {
//		go func(generateResults func() []result) {
//			defer wg.Done()
//
//			for _, r := range generateResults() {
//				select {
//				case <-ctx.Done():
//					return
//				case resultsStream <- r:
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
