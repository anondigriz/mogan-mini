package creator

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	entKB "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	kbStorage "github.com/anondigriz/mogan-mini/internal/storage/insqlite/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
)

type Migrator interface {
	Migrate() error
}

type Creator struct {
	lg  *zap.Logger
	cfg config.Config
	fc  *filecreator.FileCreator
}

func New(lg *zap.Logger, cfg config.Config) *Creator {
	fc := filecreator.New(lg)

	return &Creator{
		lg:  lg,
		cfg: cfg,
		fc:  fc,
	}
}

func (c Creator) generateFilePath(fileName string) string {
	file := path.Join(c.cfg.ProjectsPath, fileName+".db")
	return file
}

func (c Creator) CreateProject(ctx context.Context, name string) (*kbStorage.Storage, error) {
	kbUUID := uuid.New().String()
	filePath := c.generateFilePath(kbUUID)
	err := c.fc.CreateFile(filePath)
	if err != nil {
		c.lg.Error("fail to create a database file for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	dsn := fmt.Sprintf("file:%s", filePath)

	st, err := kbStorage.New(ctx, c.lg, dsn, c.cfg.Databases.LogLevel)
	if err != nil {
		c.lg.Error("fail to init a new database for the project of the knowledge base", zap.Error(err))
		return nil, err
	}

	kb := buildEntity(filePath, name)

	err = st.CreateKnowledgeBase(ctx, kb)
	if err != nil {
		c.lg.Error("fail to create knowledge base int the database", zap.Error(err))
		return nil, err
	}

	return st, nil
}

func buildEntity(kbUUID string, name string) entKB.KnowledgeBase {
	now := time.Now().UTC()
	kb := entKB.KnowledgeBase{
		BaseInfo: entKB.BaseInfo{
			UUID:         kbUUID,
			ID:           kbUUID,
			ShortName:    name,
			CreatedDate:  now,
			ModifiedDate: now,
		},
	}
	return kb
}
