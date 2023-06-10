package group

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	editCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/baseinfo/edit"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/checks"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

type Create struct {
	lg  *logger.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewCreate(lg *logger.Logger, vp *viper.Viper, cfg *config.Config) *Create {
	create := &Create{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	create.Cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a group",
		Long:  `Create a knowledge base group`,
		RunE:  create.runE,
	}
	return create
}

func (c *Create) Init() {
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) runE(cmd *cobra.Command, args []string) error {
	if err := checks.IsKnowledgeBaseChosen(*c.lg.Zap, c.cfg.CurrentKnowledgeBase.UUID); err != nil {
		return err
	}

	ec := editCLI.New(c.lg.Zap)
	now := time.Now().UTC()
	info := kbEnt.BaseInfo{
		CreatedDate:  now,
		ModifiedDate: now,
	}
	info, err := ec.EditTUI(info)
	if err != nil {
		c.lg.Zap.Error(errMsgs.CreateTUIKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.CreateTUIKnowledgeBaseFail)
		return err
	}

	return c.createGroup(info)
}

func (c Create) createGroup(info kbEnt.BaseInfo) error {
	st := kbsSt.New(c.lg.Zap, c.cfg.WorkspaceDir)
	kbsu := kbsUC.New(c.lg.Zap, st)
	group := kbEnt.Group{
		BaseInfo: info,
	}
	uuid, err := kbsu.CreateGroup(c.cfg.CurrentKnowledgeBase.UUID, group)
	if err != nil {
		c.lg.Zap.Error(errMsgs.CreateGroupFail, zap.Error(err))
		messages.PrintFail(errMsgs.CreateGroupFail)
		return err
	}
	messages.PrintGroupCreated(uuid)
	return nil
}
