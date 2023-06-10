package group

import (
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

type Edit struct {
	lg     *logger.Logger
	vp     *viper.Viper
	cfg    *config.Config
	Cmd    *cobra.Command
	grUUID string
}

func NewEdit(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Edit {
	edit := &Edit{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	edit.Cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit the group",
		Long:  `Edit the knowledge base group`,
		RunE:  edit.runE,
	}
	return edit
}

func (e *Edit) Init() {
	e.Cmd.PersistentFlags().StringVar(&e.grUUID, "uuid", "", "group UUID")

	cobra.OnInitialize(e.initConfig)
}

func (e *Edit) initConfig() {
}

func (e *Edit) runE(cmd *cobra.Command, args []string) error {
	if err := checks.IsKnowledgeBaseChosen(*e.lg.Zap, e.cfg.CurrentKnowledgeBase.UUID); err != nil {
		return err
	}

	kbsu := kbsUC.New(e.lg.Zap,
		kbsSt.New(e.lg.Zap, e.cfg.WorkspaceDir))

	if e.grUUID == "" {
		uuid, err := chooseGroup(e.lg.Zap, kbsu, e.cfg.CurrentKnowledgeBase.UUID)
		if err != nil {
			e.lg.Zap.Error(errMsgs.ChooseGroupFail, zap.Error(err))
			messages.PrintFail(errMsgs.ChooseGroupFail)
			return err
		}
		e.grUUID = uuid
	}

	messages.PrintChosenGroup(e.grUUID)

	return nil
}
