package knowledgebase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/anondigriz/mogan-mini/internal/config"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	"github.com/anondigriz/mogan-mini/internal/utility/exchange/kbimport"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Import struct {
	lg      *zap.Logger
	vp      *viper.Viper
	cfg     *config.Config
	Cmd     *cobra.Command
	xmlPath string
}

func NewImport(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Import {
	im := &Import{
		lg:  lg,
		vp:  vp,
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
	im.Cmd.PersistentFlags().StringVarP(&im.xmlPath, "path", "p", "", "path to the imported xml file")
	cobra.OnInitialize(im.initConfig)
}

func (im *Import) initConfig() {
}

func (im *Import) runE(cmd *cobra.Command, args []string) error {
	if im.xmlPath == "" {
		cmd.Help()
		err := fmt.Errorf("The path to the imported xml file was not specified. Please pass it through the command line arguments.")
		return err
	}

	f, err := os.Open(im.xmlPath)
	if err != nil {
		im.lg.Error("Fail to open the XML file", zap.Error(err))
		return err
	}

	defer f.Close()

	kbim, err := kbimport.New(im.lg, *im.cfg)
	if err != nil {
		im.lg.Error("Fail to init knowledge base importer", zap.Error(err))
		return err
	}
	uuid := uuid.New().String()
	arg := argsCore.ImportKnowledgeBase{
		KnowledgeBaseUUID: uuid,
		XMLFile:           f,
		FileName:          filepath.Base(im.xmlPath),
	}
	cont, err := kbim.Parse(cmd.Context(), arg)
	if err != nil {
		im.lg.Error("Fail to parse xml file", zap.Error(err))
		return err
	}
	_ = cont

	return nil
}
