package knowledgebase

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

type Import struct {
	lg      *logger.Logger
	cfg     *config.Config
	Cmd     *cobra.Command
	xmlPath string
}

func NewImport(lg *logger.Logger, cfg *config.Config) *Import {
	im := &Import{
		lg:  lg,
		cfg: cfg,
	}

	im.Cmd = &cobra.Command{
		Use:   "import",
		Short: "Import the knowledge base",
		Long:  `Import the knowledge base to the local workspace`,
		RunE:  im.runE,
	}
	return im
}

func (im *Import) Init() {
	im.Cmd.PersistentFlags().StringVarP(&im.xmlPath, "path", "p", "", "path to the xml file to import")
	cobra.OnInitialize(im.initConfig)
}

func (im *Import) initConfig() {
}

func (im *Import) runE(cmd *cobra.Command, args []string) error {
	if im.xmlPath == "" {
		err := fmt.Errorf(errMsgs.XMLFilePathIsEmpty)
		im.lg.Zap.Error(err.Error())
		messages.PrintFail(errMsgs.XMLFilePathIsEmpty)
		return err
	}

	f, err := os.Open(im.xmlPath)
	if err != nil {
		im.lg.Zap.Error(err.Error(), zap.Error(err))
		messages.PrintFail(errMsgs.XMLFileOpen)
		return err
	}
	defer f.Close()

	st := kbsSt.New(im.lg.Zap, *im.cfg)
	kbsu := kbsUC.New(im.lg.Zap, st)

	iArgs := argsCore.ImportKnowledgeBase{
		XMLFile:  f,
		FileName: f.Name(),
	}

	uuid, err := kbsu.ImportKnowledgeBase(iArgs)
	if err != nil {
		im.lg.Zap.Error(errMsgs.ImportProject, zap.Error(err))
		messages.PrintFail(errMsgs.ImportProject)
		return err
	}

	messages.PrintCreatedKnowledgeBase(uuid)
	return nil
}
