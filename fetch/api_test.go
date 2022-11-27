package fetch

import (
	"testing"
)

func TestHashCheaker(t *testing.T) {
	res, err := Search(&SearchOpt{
		Query:   "test",
		Results: 1,
	})

	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	if res[0] != "2F2DBA2A621B693BB95601C16ED680F8" {
		t.Error("Cant Match MD5")
	}
}
