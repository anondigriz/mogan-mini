package knowledgebase

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

const (
	timeFormat = "02.01.2006 15:04:05"
)

type Show struct {
	lg     *logger.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	kbUUID string
}

func NewShow(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Show {
	s := &Show{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	s.Cmd = &cobra.Command{
		Use:   "show",
		Short: "Show knowledge bases",
		Long:  `Show knowledge bases located in the working directory`,
		RunE:  s.runE,
	}
	return s
}

func (s *Show) Init() {
	s.Cmd.PersistentFlags().StringVar(&s.kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(s.initConfig)
}

func (s *Show) initConfig() {
}

func (s *Show) runE(cmd *cobra.Command, args []string) error {
	if s.kbUUID == "" {
		kbsu := kbsUC.New(s.lg.Zap,
			kbsSt.New(s.lg.Zap, s.cfg.WorkspaceDir))
		uuid, err := chooseKnowledgeBase(s.lg.Zap, kbsu)
		if err != nil {
			s.lg.Zap.Error(errMsgs.ChooseKnowledgeBaseFail, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseKnowledgeBaseFail)
			return err
		}
		s.kbUUID = uuid
	}

	kbsu := kbsUC.New(s.lg.Zap,
		kbsSt.New(s.lg.Zap, s.cfg.WorkspaceDir))
	kb, err := kbsu.GetKnowledgeBase(s.kbUUID)
	if err != nil {
		s.lg.Zap.Error(errMsgs.GetKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetKnowledgeBaseFail)
		return err
	}

	return showKnowledgeBase(kb)
}

func showKnowledgeBase(kb kbEnt.KnowledgeBase) error {
	messages.PrintKnowledgeBaseInfo()
	fmt.Printf("UUID: %s\n", kb.UUID)
	fmt.Printf("ID: %s\n", kb.ID)
	fmt.Printf("Short name: %s\n", kb.ShortName)
	fmt.Printf("Description:\n%s\n", kb.Description)
	fmt.Printf("CreatedDate: %s\n", kb.CreatedDate.Local().Format(timeFormat))
	fmt.Printf("ModifiedDate: %s\n", kb.ModifiedDate.Local().Format(timeFormat))
	return nil
}
