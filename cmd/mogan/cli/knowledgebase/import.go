package knowledgebase

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anondigriz/mogan-mini/internal/config"
	argsCore "github.com/anondigriz/mogan-mini/internal/core/args"
	kbEnt "github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
	"github.com/anondigriz/mogan-mini/internal/utility/exchange/kbimport"
	"github.com/anondigriz/mogan-mini/internal/utility/knowledgebase/dbcreator"
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

	f, err := im.openFile()
	if err != nil {
		return err
	}
	defer f.Close()

	cont, err := im.parseFile(cmd.Context(), f)
	if err != nil {
		return err
	}
	defer f.Close()

	err = im.createDB(cmd.Context(), cont)
	if err != nil {
		return err
	}

	return nil
}

func (im *Import) openFile() (*os.File, error) {
	f, err := os.Open(im.xmlPath)
	if err != nil {
		im.lg.Error("Fail to open the XML file", zap.Error(err))
		return nil, err
	}
	return f, nil
}

func (im *Import) parseFile(ctx context.Context, f *os.File) (kbEnt.Container, error) {
	kbim, err := kbimport.New(im.lg, *im.cfg)
	if err != nil {
		im.lg.Error("Fail to init knowledge base importer", zap.Error(err))
		return kbEnt.Container{}, err
	}
	uuid := uuid.New().String()
	arg := argsCore.ImportKnowledgeBase{
		KnowledgeBaseUUID: uuid,
		XMLFile:           f,
		FileName:          filepath.Base(im.xmlPath),
	}
	cont, err := kbim.Parse(ctx, arg)
	if err != nil {
		im.lg.Error("Fail to parse xml file", zap.Error(err))
		return kbEnt.Container{}, err
	}
	return cont, nil
}

func (im *Import) createDB(ctx context.Context, cont kbEnt.Container) error {
	dc := dbcreator.New(im.lg, *im.cfg)
	st, err := dc.Create(ctx, cont.KnowledgeBase.ShortName, dc.GenerateFilePathWithUUID(cont.KnowledgeBase.UUID))
	if err != nil {
		im.lg.Error("fail to create database for the project of the knowledge base", zap.Error(err))
		return err
	}
	defer st.Shutdown()

	err = st.FillFromContainer(ctx, cont)
	if err != nil {
		im.lg.Error("fail to fill the database of the knowledge base project by the data from the xml file", zap.Error(err))
		return err
	}
	return nil
}