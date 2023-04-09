package knowledgebase

import (
	"context"
	"fmt"

	kbEnt "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	baseinfoKBTui "github.com/anondigriz/mogan-editor-cli/internal/tui/baseinfo/edit"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/anondigriz/mogan-editor-cli/cmd/mogan/cli/errors"
	"github.com/anondigriz/mogan-editor-cli/internal/config"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/connection"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Edit struct {
	lg  *zap.Logger
	vp  *viper.Viper
	cfg *config.Config
	Cmd *cobra.Command
}

func NewEdit(lg *zap.Logger, vp *viper.Viper, cfg *config.Config) *Edit {
	e := &Edit{
		lg:  lg,
		vp:  vp,
		cfg: cfg,
	}

	e.Cmd = &cobra.Command{
		Use:   "edit",
		Short: "Editing the database",
		Long:  `Editing basic information about the knowledge base`,
		Run:   e.run,
	}
	return e
}

func (e *Edit) Init() {
	cobra.OnInitialize(e.initConfig)
}

func (e *Edit) initConfig() {
}

func (e *Edit) run(cmd *cobra.Command, args []string) {
	if e.cfg.CurrentKnowledgeBase.UUID == "" {
		err := errors.KnowledgeBaseNotChosenErr
		e.lg.Error(err.Error(), zap.Error(err))
		fmt.Print(err.Error())
		return
	}
	con := connection.New(e.lg, *e.cfg)
	st, err := con.GetByUUID(cmd.Context(), e.cfg.CurrentKnowledgeBase.UUID)
	if err != nil {
		e.lg.Error("Error to get connection with database connection", zap.Error(err))
		fmt.Printf("An unexpected error occurred when opening a knowledge base project: %v\n", err)
		return
	}
	defer st.Shutdown()

	kb, err := st.GetKnowledgeBase(cmd.Context())
	if err != nil {
		e.lg.Error("Error getting knowledge base information", zap.Error(err))
		fmt.Printf("\n---\nError getting knowledge base information: %v\n", err)
		return
	}
	updKb, err := e.editKnowledgeBase(cmd.Context(), kb)
	if err != nil {
		e.lg.Error("An error occurred while editing the knowledge base", zap.Error(err))
		fmt.Printf("\n---\nAn error occurred while editing the knowledge base: %v\n", err)
		return
	}

	err = st.UpdateKnowledgeBase(cmd.Context(), updKb)
	if err != nil {
		e.lg.Error("An error occurred while updating the knowledge base", zap.Error(err))
		fmt.Printf("\n---\nAn error occurred while updating the knowledge base: %v\n", err)
		return
	}
	fmt.Print("\n---\nGreat, you've changed the basic information about the knowledge base!\n")
}

func (e *Edit) editKnowledgeBase(ctx context.Context, kb kbEnt.KnowledgeBase) (kbEnt.KnowledgeBase, error) {
	mt := baseinfoKBTui.New(kb.BaseInfo)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		e.lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return kbEnt.KnowledgeBase{}, err
	}
	def := kbEnt.BaseInfo{}
	if m, ok := m.(baseinfoKBTui.Model); ok && m.BaseInfo != def {
		if m.BaseInfo.ID == "" {
			return kbEnt.KnowledgeBase{}, errors.IDIsEmptyErr
		}
		if m.BaseInfo.ShortName == "" {
			return kbEnt.KnowledgeBase{}, errors.ShortNameIsEmptyErr
		}
		kb.BaseInfo = m.BaseInfo
		return kb, nil
	}
	return kbEnt.KnowledgeBase{}, fmt.Errorf("Knowledge base has not been edited")
}
