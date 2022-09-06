package eth

import (
	"fmt"
	"testing"

	pkglog "bls/pkg/log"
	"bls/pkg/log/tee"

	"github.com/go-kratos/kratos/v2/log"
)

func TestCheckBlockNumber(t *testing.T) {
	uri := fmt.Sprintf("http://%s:%d", "localhost", 8545)
	logger, err := pkglog.NewZapLogger(tee.Level(log.LevelDebug), tee.Stdout(true))
	if err != nil {
		t.Errorf("h = %v", err)
		return
	}
	defer logger.Close()
	cli := New(logger, uri)
	h, err := cli.EthBlockNumber()
	if err != nil {
		return
	}
	t.Errorf("h = %v", h)
}
