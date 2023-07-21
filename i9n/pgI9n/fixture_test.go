package pgI9n

import (
	"context"
	"fmt"
	"github.com/remicro/trifle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type address struct {
	ID       int
	Street   string
	Building int
}

func testFunc(t *testing.T, exp address) {
	fix := suite.With(t, WithMigrations).(*Fixture)
	defer fix.TearDown()

	_, err := fix.DB.Exec(context.Background(), "INSERT INTO public.address (street, building) VALUES ($1, $2)", exp.Street, exp.Building)
	require.NoError(t, err)
	results := make([]address, 0, 1)
	rows, err := fix.DB.Query(context.Background(), "SELECT id, street, building FROM public.address")
	require.NoError(t, err)
	for rows.Next() {
		var res address
		require.NoError(t, rows.Scan(&res.ID, &res.Street, &res.Building))
		results = append(results, res)
	}
	require.Len(t, results, 1)
	assert.Equal(t, exp, results[0])
}

func TestWithMigrations(t *testing.T) {
	exp := address{
		ID:       1,
		Street:   trifle.String(),
		Building: 42,
	}

	for i := 0; i < 10; i++ {
		testName := fmt.Sprintf("should be able to correct save data in %d concurrent test", i)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			testFunc(t, exp)
		})
	}

}
