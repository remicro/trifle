package i9n

type Fixture struct {
	Prev SuiteFixture
}

func (f *Fixture) TearUp() error {
	return nil
}

func (f *Fixture) TearDown() error {
	if f.Prev != nil {
		return f.Prev.TearDown()
	}
	return nil
}
