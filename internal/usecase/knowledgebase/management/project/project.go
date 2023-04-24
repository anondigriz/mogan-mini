package project

import (
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/manager"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/management/pathmaker"
	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
)

type settings struct {
	ProjectsPath string
}

type Project struct {
	lg       *zap.Logger
	settings settings
	fc       *filecreator.FileCreator
	con      *connection.Connection
	pm       *pathmaker.PathMaker
	man      *manager.Manager
}

func New(lg *zap.Logger, cfg config.Config, con *connection.Connection, pm *pathmaker.PathMaker,
	man *manager.Manager, fc *filecreator.FileCreator) *Project {
	p := &Project{
		lg:  lg,
		con: con,
		pm:  pm,
		fc:  fc,
		man: man,
		settings: settings{
			ProjectsPath: cfg.ProjectsPath,
		},
	}
	return p
}
