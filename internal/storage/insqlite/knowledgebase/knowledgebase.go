package knowledgebase

import (
	"context"
	"database/sql"
	"embed"

	"github.com/anondigriz/mogan-core/pkg/loglevel"
	"github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/storage/insqlite/errors"
	"github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase/mappers"
	"github.com/anondigriz/mogan-core/pkg/insqlite/migrator"

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

func (st *Storage) CreateKnowledgeBase(ctx context.Context, knowledgeBase knowledgebase.KnowledgeBase) error {
	script := `INSERT INTO "KnowledgeBase"("UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", "ExtraData")
	VALUES ($1, $2, $3, $4, $5, $6);`

	var row mappers.KnowledgeBaseRow
	err := row.Fill(knowledgeBase)
	if err != nil {
		st.lg.Error("fail to map knowledgeBase", zap.Error(err))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	_, err = st.db.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.CreatedDate, row.ModifiedDate, row.ExtraData)

	if err != nil {
		st.lg.Error("unexpected error when trying to insert the knowledge base", zap.Error(err))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	return nil
}

func (st *Storage) GetKnowledgeBase(ctx context.Context) (knowledgebase.KnowledgeBase, error) {
	script := `SELECT "UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", "ExtraData"
	FROM "KnowledgeBase" LIMIT 1;`

	var row mappers.KnowledgeBaseRow
	err := st.db.QueryRowContext(ctx,
		script).Scan(&row.UUID, &row.ID, &row.ShortName, &row.CreatedDate, &row.ModifiedDate, &row.ExtraData)

	if err != nil {
		st.lg.Error("unexpected error when trying to get the knowledge base", zap.Error(err))
		return knowledgebase.KnowledgeBase{}, errors.NewUnexpectedStorageFailErr(err)
	}

	kb, err := row.Extract()
	if err != nil {
		st.lg.Error("fail to map knowledgeBase", zap.Error(err))
		return knowledgebase.KnowledgeBase{}, errors.NewUnexpectedStorageFailErr(err)
	}

	return kb, nil
}

func (st *Storage) UpdateKnowledgeBase(ctx context.Context, knowledgeBase knowledgebase.KnowledgeBase) error {
	script := `UPDATE "KnowledgeBase" SET "ID" = $1, "ShortName" = $2, "ModifiedDate" = $3, "ExtraData" = $4
	WHERE "UUID" = $5;`

	var row mappers.KnowledgeBaseRow
	err := row.Fill(knowledgeBase)
	if err != nil {
		st.lg.Error("fail to map knowledgeBase", zap.Error(err))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	_, err = st.db.ExecContext(ctx,
		script,
		row.ID, row.ShortName, row.ModifiedDate, row.ExtraData, row.UUID)
	if err != nil {
		st.lg.Error("unexpected error when trying to update the knowledge base", zap.Error(err))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	return nil
}
