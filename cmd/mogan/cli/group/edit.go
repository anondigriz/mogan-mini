package group

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	chooseCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/group/choose"
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
	if e.cfg.CurrentKnowledgeBase.UUID == "" {
		err := fmt.Errorf(errMsgs.KnowledgeBaseNotChosen)
		e.lg.Zap.Error(err.Error())
		messages.PrintFail(errMsgs.KnowledgeBaseNotChosen)
		return err
	}

	kbsu := kbsUC.New(e.lg.Zap,
		kbsSt.New(e.lg.Zap, e.cfg.WorkspaceDir))

	if e.grUUID == "" {
		uuid, err := e.chooseGroup(kbsu)
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

func (e *Edit) chooseGroup(kbsu *kbsUC.KnowledgeBases) (string, error) {
	ch := chooseCLI.New(e.lg.Zap)
	if err := ch.Init(kbsu, e.cfg.CurrentKnowledgeBase.UUID); err != nil {
		e.lg.Zap.Error(errMsgs.InitGroupChooserFail, zap.Error(err))
		messages.PrintFail(errMsgs.InitGroupChooserFail)
		return "", err
	}
	uuid, err := ch.ChooseGroup()
	if err != nil {
		e.lg.Zap.Error(errMsgs.ChooseGroupFail, zap.Error(err))
		return "", err
	}
	return uuid, nil
}
