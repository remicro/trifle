package pgI9n

import (
	"github.com/remicro/trifle/i9n"
	"os"
	"testing"
)

var suite *i9n.Suite

func TestMain(m *testing.M) {
	suite = i9n.New(m, New("sql/correct"))
	os.Exit(suite.Run())
}
