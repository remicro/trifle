package pgI9n

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/remicro/trifle/i9n"
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
)

func WithMigrations(t *testing.T, parent i9n.SuiteFixture) (i9n.SuiteFixture, error) {
	prev, ok := parent.(*PGI9Suite)
	if !ok {
		return nil, errors.New("unexpected parent type")
	}

	cnt := atomic.AddInt64(&prev.allocatedDB, 1)
	cfg := *prev.cfg
	connCfg := *cfg.ConnConfig
	connCfg.Database = fmt.Sprintf("%s_%d", prev.cfg.ConnConfig.Database, cnt)

	_, err := prev.DB.Exec(context.Background(),
		fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s",
			connCfg.Database,
			connCfg.User,
		))
	require.NoError(t, err)
	cfg.ConnConfig = &connCfg

	fxt := Fixture{}

	fxt.DB, err = pgxpool.NewWithConfig(prev.ctx, &cfg)
	require.NoError(t, err)
	require.NoError(t, fxt.DB.Ping(prev.ctx))

	for _, migration := range prev.migrations {
		_, err := fxt.DB.Exec(context.Background(), migration)
		if err != nil {
			return nil, err
		}
	}
	return &fxt, nil
}

type Fixture struct {
	i9n.Fixture
	DB *pgxpool.Pool
}
