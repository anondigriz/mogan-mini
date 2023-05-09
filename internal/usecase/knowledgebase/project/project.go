package project

import (
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/storage/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/connection"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/editor"
	"github.com/anondigriz/mogan-mini/internal/usecase/knowledgebase/pathmaker"
	"github.com/anondigriz/mogan-mini/internal/utility/filecreator"
)

type settings struct {
	ProjectsPath string
}

type NewProjectArgs struct {
	Lg     *zap.Logger
	Cfg    config.Config
	Con    *connection.Connection
	Pm     *pathmaker.PathMaker
	Editor *editor.Editor
	Fc     *filecreator.FileCreator
	Parser *parser.Parser
	Strc   *knowledgebase.StorageCreator
}

type Project struct {
	lg       *zap.Logger
	settings settings
	fc       *filecreator.FileCreator
	con      *connection.Connection
	pm       *pathmaker.PathMaker
	editor   *editor.Editor
	parser   *parser.Parser
	strc     *knowledgebase.StorageCreator
}

func New(args NewProjectArgs) *Project {
	p := &Project{
		lg:     args.Lg,
		con:    args.Con,
		pm:     args.Pm,
		fc:     args.Fc,
		editor: args.Editor,
		parser: args.Parser,
		strc:   args.Strc,
		settings: settings{
			ProjectsPath: args.Cfg.ProjectsPath,
		},
	}
	return p
}
