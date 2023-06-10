package knowledgebase

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	editCLI "github.com/anondigriz/mogan-mini/cmd/mogan/cli/baseinfo/edit"
	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	"github.com/anondigriz/mogan-mini/internal/config"
	"github.com/anondigriz/mogan-mini/internal/logger"
	kbsSt "github.com/anondigriz/mogan-mini/internal/storage/knowledgebases"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

type Create struct {
	lg  *logger.Logger
	cfg *config.Config
	Cmd *cobra.Command
}

func NewCreate(lg *logger.Logger, cfg *config.Config) *Create {
	c := &Create{
		lg:  lg,
		cfg: cfg,
	}

	c.Cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a local knowledge base",
		Long:  `Create a local knowledge base in the base project directory`,
		RunE:  c.runE,
	}
	return c
}

func (c *Create) Init() {
	cobra.OnInitialize(c.initConfig)
}

func (c *Create) initConfig() {
}

func (c *Create) runE(cmd *cobra.Command, args []string) error {
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

	return c.createKnowledgeBase(info)
}

func (c Create) createKnowledgeBase(info kbEnt.BaseInfo) error {
	st := kbsSt.New(c.lg.Zap, c.cfg.WorkspaceDir)
	kbsu := kbsUC.New(c.lg.Zap, st)
	knowledgeBase := kbEnt.KnowledgeBase{
		BaseInfo: info,
	}
	uuid, err := kbsu.CreateKnowledgeBase(knowledgeBase)
	if err != nil {
		c.lg.Zap.Error(errMsgs.CreateKnowledgeBaseProjectFail, zap.Error(err))
		messages.PrintFail(errMsgs.CreateKnowledgeBaseProjectFail)
		return err
	}
	messages.PrintCreatedKnowledgeBase(uuid)
	return nil
}
