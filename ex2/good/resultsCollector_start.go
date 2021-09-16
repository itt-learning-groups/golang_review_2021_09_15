package good

type result struct{}

func Run() []result {
	resultsGenerators := []func() []result{
		func() []result { return []result{{}} },             // generator1 returns 1 result
		func() []result { return []result{{}, {}} },         // generator2 returns 2 results
		func() []result { return []result{{}, {}, {}} },     // generator3 returns 3 results
		func() []result { return []result{{}, {}, {}, {}} }, // generator4 returns 4 results
	}

	resultsStream := resultsStreamOwner(resultsGenerators)
	return resultsStreamConsumer(resultsStream)
}

func resultsStreamOwner(resultsGenerators []func() []result) <-chan result {
	resultsStream := make(chan result)

	// Todo: implement me

	return resultsStream
}

func resultsStreamConsumer(resultsStream <-chan result) []result {
	var results []result

	// Todo: implement me

	return results
}
