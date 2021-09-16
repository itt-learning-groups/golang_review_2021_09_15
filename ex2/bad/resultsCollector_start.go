package bad

type collector struct {
	generators    []func() []result
	results       []result
	resultsStream chan result
}

type result struct{}

func Run() []result {
	c := newCollector()

	c.generateResults()
	c.consumeResults()

	return c.results
}

func newCollector() *collector {
	resultsStream := make(chan result)
	return &collector{
		resultsStream: resultsStream,
		generators: []func() []result{
			func() []result { return []result{{}} },             // generator1 returns 1 result
			func() []result { return []result{{}, {}} },         // generator2 returns 2 results
			func() []result { return []result{{}, {}, {}} },     // generator3 returns 3 results
			func() []result { return []result{{}, {}, {}, {}} }, // generator4 returns 4 results
		},
	}
}

func (c *collector) generateResults() {
	// Todo: implement me
}

func (c *collector) consumeResults() {
	// Todo: implement me
}
