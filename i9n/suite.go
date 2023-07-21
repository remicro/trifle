package i9n

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

type CallBackFunction func() error

type SuiteFixture interface {
	TearUp() error
	TearDown() error
}

type ForkFixture func(t *testing.T, prev SuiteFixture) (SuiteFixture, error)

func New(m *testing.M, fix SuiteFixture) *Suite {
	return &Suite{
		fixture: fix,
		m:       m,
	}
}

type Suite struct {
	fixture SuiteFixture
	m       *testing.M
}

func (tst *Suite) TearDown() {
	tst.NoError(tst.fixture.TearDown())
}

func (tst *Suite) NoError(err error) {
	if err != nil {
		tst.Fatal(err.Error())
	}
}

func (tst *Suite) With(t *testing.T, forks ...ForkFixture) SuiteFixture {
	prev := tst.fixture
	for _, f := range forks {
		res, err := f(t, prev)
		require.NoError(t, err)
		require.NoError(t, res.TearUp())
		prev = res
	}
	return prev
}

func (tst *Suite) Run() int {
	tst.NoError(tst.fixture.TearUp())
	defer tst.TearDown()
	return tst.m.Run()
}

func (tst *Suite) Fatal(v ...any) {
	tst.TearDown()
	log.Fatal(v...)
}
