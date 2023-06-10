package group

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/checks"
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
	grUUID string
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
	s.Cmd.PersistentFlags().StringVar(&s.grUUID, "uuid", "", "UUID of the group")
	cobra.OnInitialize(s.initConfig)
}

func (s *Show) initConfig() {
}

func (s *Show) runE(cmd *cobra.Command, args []string) error {
	if err := checks.IsKnowledgeBaseChosen(*s.lg.Zap, s.cfg.CurrentKnowledgeBase.UUID); err != nil {
		return err
	}

	kbsu := kbsUC.New(s.lg.Zap,
		kbsSt.New(s.lg.Zap, s.cfg.WorkspaceDir))
	if s.grUUID == "" {
		uuid, err := chooseGroup(s.lg.Zap, kbsu, s.cfg.CurrentKnowledgeBase.UUID)
		if err != nil {
			s.lg.Zap.Error(errMsgs.ChooseGroupFail, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseGroupFail)
			return err
		}
		s.grUUID = uuid
	}

	gr, err := kbsu.GetGroup(s.cfg.CurrentKnowledgeBase.UUID, s.grUUID)
	if err != nil {
		s.lg.Zap.Error(errMsgs.GetGroupFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetGroupFail)
		return err
	}

	return showGroup(gr)
}

func showGroup(gr kbEnt.Group) error {
	messages.PrintGroupInfo()
	fmt.Printf("UUID: %s\n", gr.UUID)
	fmt.Printf("ID: %s\n", gr.ID)
	fmt.Printf("Short name: %s\n", gr.ShortName)
	fmt.Printf("Description:\n%s\n", gr.Description)
	fmt.Printf("CreatedDate: %s\n", gr.CreatedDate.Local().Format(timeFormat))
	fmt.Printf("ModifiedDate: %s\n", gr.ModifiedDate.Local().Format(timeFormat))
	return nil
}
