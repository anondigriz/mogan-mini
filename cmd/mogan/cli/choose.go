package cli

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	chooseKBTui "github.com/anondigriz/mogan-editor-cli/internal/tui/knowledgebase/choose"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/localfinder"
)

var (
	kbUUID    string
	chooseCmd = &cobra.Command{
		Use:   "choose",
		Short: "Choose a knowledge base project to work with",
		Long:  `Choose a knowledge base project from the base project directory to be used in the workspace`,
		Run: func(cmd *cobra.Command, args []string) {
			if kbUUID == "" {
				uuid, err := chooseKnowledgeBase(cmd.Context())
				if err != nil {
					fmt.Printf("\n---\nThere was a problem when choosing a knowledge base: %v\n", err)
					return
				}
				kbUUID = uuid
			}

			fmt.Printf("\n---\nOkay, you have selected a knowledge base project with UUID %s\n", kbUUID)
			vp.Set("CurrentKnowledgeBase.UUID", kbUUID)
			err := vp.WriteConfig()
			if err != nil {
				lg.Error("fail to write config", zap.Error(err))
				os.Exit(1)
			}
		},
	}
)

func initChooseCmd() {
	chooseCmd.PersistentFlags().StringVar(&kbUUID, "uuid", "", "knowledge base project UUID")
	cobra.OnInitialize(initChooseCmdCfg)
}

func initChooseCmdCfg() {
}

func chooseKnowledgeBase(ctx context.Context) (string, error) {
	lf := localfinder.New(lg, *cfg)
	kbs := lf.FindInProjectsDir(ctx)
	mt := chooseKBTui.New(kbs)
	p := tea.NewProgram(mt)
	m, err := p.Run()
	if err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", err
	}
	if m, ok := m.(chooseKBTui.Model); ok && m.Choice != "" {
		return m.Choice, nil
	}
	return "", fmt.Errorf("Knowledge base was not chosen")
}
