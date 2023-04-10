package knowledgebase

import (
	"context"
	"database/sql"

	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/storage/insqlite/errors"
	"github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase/mappers"
	"go.uber.org/zap"
)

func (st *Storage) FillFromContainer(ctx context.Context, c kbEnt.Container) error {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		st.lg.Error("Fail to start transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`DELETE FROM "Group";`)
	if err != nil {
		// TODO
		st.lg.Error("unexpected delete groups of the knowledge base",
			zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM "Pattern";`)
	if err != nil {
		// TODO
		st.lg.Error("unexpected delete patterns of the knowledge base",
			zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	for _, v := range c.Groups {
		err = insertGroup(ctx, tx, v)
		if err != nil {
			// TODO проверка уникальности ("ID") (ограничение)
			st.lg.Error("unexpected error when trying to insert the group",
				zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
			return errors.NewUnexpectedStorageFailErr(err)
		}
	}

	for _, v := range c.Parameters {
		err = insertParameter(ctx, tx, v)
		if err != nil {
			// TODO проверка уникальности ("ID") (ограничение)
			st.lg.Error("unexpected error when trying to insert the parameter",
				zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
			return errors.NewUnexpectedStorageFailErr(err)
		}
	}

	for _, v := range c.Patterns {
		err = insertPattern(ctx, tx, v)
		if err != nil {
			// TODO проверка уникальности ("ID") (ограничение)
			st.lg.Error("unexpected error when trying to insert the pattern",
				zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
			return errors.NewUnexpectedStorageFailErr(err)
		}
	}

	for _, v := range c.Rules {
		err = insertRule(ctx, tx, v)
		if err != nil {
			// TODO проверка уникальности ("ID") (ограничение)
			st.lg.Error("unexpected error when trying to insert the rule",
				zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
			return errors.NewUnexpectedStorageFailErr(err)
		}
	}

	err = updateKnowledgeBase(ctx, tx, c.KnowledgeBase)
	if err != nil {
		// TODO
		st.lg.Error("unexpected error when trying to update the knowledge base",
			zap.Error(err), zap.String("KnowledgeBaseUUID", c.KnowledgeBase.UUID))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	err = tx.Commit()
	if err != nil {
		// TODO
		st.lg.Error("commit failed", zap.Error(err))
		return errors.NewUnexpectedStorageFailErr(err)
	}

	return nil
}

func insertGroup(ctx context.Context, tx *sql.Tx, group kbEnt.Group) error {
	script := `INSERT INTO "Group"("UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", "ExtraData")
	VALUES ($1, $2, $3, $4, $5, $6);`

	var row mappers.GroupRow
	err := row.Fill(group)
	if err != nil {
		return errors.NewUnexpectedStorageFailErr(err)
	}
	_, err = tx.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.CreatedDate, row.ModifiedDate, row.ExtraData)

	if err != nil {
		return err
	}

	return nil
}

func insertParameter(ctx context.Context, tx *sql.Tx, parameter kbEnt.Parameter) error {
	script := `INSERT INTO "Parameter"("UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", "GroupUUID",
	"TypeID", "ExtraData")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	var row mappers.ParameterRow
	err := row.Fill(parameter)
	if err != nil {
		return errors.NewUnexpectedStorageFailErr(err)
	}
	_, err = tx.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.CreatedDate, row.ModifiedDate, row.GroupUUID,
		row.Type, row.ExtraData)

	if err != nil {
		return err
	}

	return nil
}

func insertPattern(ctx context.Context, tx *sql.Tx, pattern kbEnt.Pattern) error {
	script := `INSERT INTO "Pattern"("UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", 
                      "TypeID", "ExtraData")
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	var row mappers.PatternRow
	err := row.Fill(pattern)
	if err != nil {
		return errors.NewUnexpectedStorageFailErr(err)
	}
	_, err = tx.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.CreatedDate, row.ModifiedDate,
		row.Type, row.ExtraData)

	if err != nil {
		return err
	}

	return nil
}

func insertRule(ctx context.Context, tx *sql.Tx, rule kbEnt.Rule) error {
	script := `INSERT INTO "Rule"("UUID", "ID", "ShortName", "CreatedDate", "ModifiedDate", "PatternUUID", "ExtraData")
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	var row mappers.RuleRow
	err := row.Fill(rule)
	if err != nil {
		return errors.NewUnexpectedStorageFailErr(err)
	}
	_, err = tx.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.CreatedDate, row.ModifiedDate, row.PatternUUID, row.ExtraData)

	if err != nil {
		return err
	}

	return nil
}

func updateKnowledgeBase(ctx context.Context, tx *sql.Tx, knowledgeBase kbEnt.KnowledgeBase) error {
	script := `UPDATE "KnowledgeBase"
	SET "ID" = $2, "ShortName" = $3, "ModifiedDate" = $4, "ExtraData" = $5
	WHERE "UUID" = $1;`

	var row mappers.KnowledgeBaseRow
	err := row.Fill(knowledgeBase)
	if err != nil {
		return errors.NewUnexpectedStorageFailErr(err)
	}
	_, err = tx.ExecContext(ctx,
		script,
		row.UUID, row.ID, row.ShortName, row.ModifiedDate, row.ExtraData)

	if err != nil {
		return err
	}

	return nil
}
