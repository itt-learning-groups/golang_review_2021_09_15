package bad

//import (
//	"sync"
//)
//
//// This is a hypothetical file we pretend is submitted for code review. The idea is to decide if it should be approved or not.
//// Note that the `resultsCollector_test.go` unit tests will pass for this code as written.
//// Overall assessment should be that, though this code is working now, and looks pretty clean/readable, it suffers from
//// some important weaknesses that expose it to high risk of accumulating bugs in future. In particular, its failure to
//// keep limited, clear ownership of the resultsStream channel, its use of mutable pointer types within concurrent goroutines,
//// and its failure to ensure that the goroutines it launches include a done-channel type circuit breaker to avoid potential
//// goroutine leaks are all red flags that indicate this code should get a 2nd draft.
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
//	defer close(c.resultsStream)
//
//	var wg sync.WaitGroup
//	wg.Add(10)
//
//	go c.consumeResults(&wg)
//	c.generateResults()
//
//	wg.Wait()
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
//func (c *collector) generateResults() {
//	for _, generator := range c.generators {
//		go func(generateResults func() []result) {
//			for _, r := range generateResults() {
//				c.resultsStream <- r
//			}
//		}(generator)
//	}
//}
//
//func (c *collector) consumeResults(wg *sync.WaitGroup) {
//	for r := range c.resultsStream {
//		c.results = append(c.results, r)
//		wg.Done()
//	}
//}
