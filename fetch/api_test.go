package fetch

import (
	"testing"
)

func TestHashCheaker(t *testing.T) {
	res, err := Search(&SearchOpt{
		Query:   "test",
		Results: 1,
	})
}
