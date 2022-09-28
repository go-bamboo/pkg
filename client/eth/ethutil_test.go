package eth

import (
	"fmt"
	"testing"
)

func TestCheckBlockNumber(t *testing.T) {
	uri := fmt.Sprintf("http://%s:%d", "localhost", 8545)
	cli := New(uri)
	h, err := cli.EthBlockNumber()
	if err != nil {
		return
	}
	t.Errorf("h = %v", h)
}
