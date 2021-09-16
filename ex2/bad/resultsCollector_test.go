package bad

import (
	"strconv"
	"testing"
)

func TestRun(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expectedResultsCount := 1 + 2 + 3 + 4

			results := Run()

			if expectedResultsCount != len(results) {
				t.Errorf("expected %d results, got %d", expectedResultsCount, results)
			}
		})
	}
}
