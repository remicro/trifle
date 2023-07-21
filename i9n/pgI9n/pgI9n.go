package pgI9n

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PGI9Suite struct {
	ctx            context.Context
	cancel         context.CancelFunc
	container      *postgres.PostgresContainer
	cfg            *pgxpool.Config
	DB             *pgxpool.Pool
	migrations     []string
	allocatedDB    int64
	migrationsPath string
}

func New(migrationPath string) *PGI9Suite {
	return &PGI9Suite{
		migrationsPath: migrationPath,
	}
}

func (suite *PGI9Suite) TearUp() error {
	suite.ctx, suite.cancel = context.WithCancel(context.Background())

	pc, err := postgres.RunContainer(suite.ctx,
		testcontainers.WithImage("docker.io/postgres:13"),
		postgres.WithDatabase("test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return err
	}
	suite.container = pc

	connString, err := pc.ConnectionString(suite.ctx)
	if err != nil {
		return err
	}

	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return err
	}

	cfg.MinConns = 2
	cfg.MaxConnIdleTime = time.Minute
	cfg.MaxConns = 10

	suite.cfg = cfg

	suite.DB, err = pgxpool.NewWithConfig(suite.ctx, cfg)
	if err != nil {
		return err
	}
	err = suite.DB.Ping(suite.ctx)
	if err != nil {
		return err
	}

	err = filepath.Walk(suite.migrationsPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		suite.migrations = append(suite.migrations, string(data))
		return nil
	})
	return err
}

func (suite *PGI9Suite) TearDown() error {
	return suite.container.Terminate(context.Background())
}
