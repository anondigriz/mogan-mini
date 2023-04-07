package knowledgebases

import (
	"context"
	"database/sql"
	"embed"

	"github.com/anondigriz/mogan-core/pkg/loglevel"
	"github.com/anondigriz/mogan-editor-cli/internal/storage/insqlite/errors"
	"github.com/anondigriz/mogan-editor-cli/pkg/insqlite/migrator"

	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Storage struct {
	lg  *zap.Logger
	db  *sql.DB
	dsn string
}

func New(ctx context.Context, lg *zap.Logger, dsn string, logLevel loglevel.LogLevel) (*Storage, error) {
	m := migrator.New(lg, dsn, &embedMigrations)
	if err := m.Migrate(); err != nil {
		lg.Error("migration fails", zap.Error(err))
		return nil, err
	}
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		lg.Error("unable to create pool", zap.Error(err))
		return nil, err
	}

	st := &Storage{lg: lg, db: db, dsn: dsn}

	return st, nil
}

func (st *Storage) GetDSN() string {
	return st.dsn
}

func (st *Storage) Shutdown() {
	st.db.Close()
}

func (st *Storage) Ping(ctx context.Context) error {
	err := st.db.PingContext(ctx)
	if err != nil {
		st.lg.Error("storage ping fail", zap.Error(err))
		return errors.NewPingFailErr(err)
	}
	return nil
}
