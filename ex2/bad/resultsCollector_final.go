package bad

//import "sync"
//
//type collector struct {
//	generators    []func() []result
//	results       []result
//	resultsStream chan result
//}
//
//type result struct{}
//
//func Run() []result {
//	c := newCollector()
//
//	done := make(chan struct{})
//	defer close(done)
//
//	c.generateResults(done)
//
//	// Uh oh!!
//	//c.generateResults(done)
//
//	c.consumeResults()
//
//	return c.results
//}
//
//func newCollector() *collector {
//	resultsStream := make(chan result)
//	return &collector{
//		resultsStream: resultsStream,
//		generators: []func() []result{
//			func() []result { return []result{{}} },             // generator1 returns 1 result
//			func() []result { return []result{{}, {}} },         // generator2 returns 2 results
//			func() []result { return []result{{}, {}, {}} },     // generator3 returns 3 results
//			func() []result { return []result{{}, {}, {}, {}} }, // generator4 returns 4 results
//		},
//	}
//}
//
//func (c *collector) generateResults(done <-chan struct{}) {
//	var wg sync.WaitGroup
//	wg.Add(len(c.generators))
//
//	for _, generator := range c.generators {
//		go func(generateResults func() []result) {
//			defer wg.Done()
//
//			for _, r := range generateResults() {
//				select {
//				case <-done:
//					return
//				case c.resultsStream <- r:
//				}
//			}
//		}(generator)
//	}
//
//	go func() {
//		wg.Wait()
//		close(c.resultsStream)
//	}()
//}
//
//func (c *collector) consumeResults() {
//	// Uh oh!!
//	//c.resultsStream <- result{}
//
//	for r := range c.resultsStream {
//		c.results = append(c.results, r)
//	}
//}
