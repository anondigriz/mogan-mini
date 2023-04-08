package cli

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	chooseKBTui "github.com/anondigriz/mogan-editor-cli/internal/tui/knowledgebase/choose"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/localfinder"
)

var (
	kbID      string
	chooseCmd = &cobra.Command{
		Use:   "choose",
		Short: "Choose a knowledge base to work with",
		Long:  `Choose a knowledge base from the base project directory to be used in the workspace`,
		Run: func(cmd *cobra.Command, args []string) {
			if kbID == "" {
				_, err := chooseKnowledgeBase(cmd.Context())
				if err != nil {
					fmt.Printf("\n---\nError entering the name of the knowledge base name: %v\n", err)
					return
				}
			}

			fmt.Printf("\n---\nYou entered the knowledge base name: %s\n", kbName)

		},
	}
)

func initChooseCmd() {
	chooseCmd.PersistentFlags().StringVar(&kbID, "id", "", "knowledge base project id")
	cobra.OnInitialize(initChooseCmdCfg)
}

func initChooseCmdCfg() {
}

func chooseKnowledgeBase(ctx context.Context) (string, error) {
	lf := localfinder.New(lg, cfg)
	kbs := lf.FindInProjectsDir(ctx)
	mt := chooseKBTui.New(kbs)
	if _, err := tea.NewProgram(mt).Run(); err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return "", nil
	}
	return "", nil
}
