package cli

import (
	"fmt"

	kbEnt "github.com/anondigriz/mogan-editor-cli/internal/entity/knowledgebase"
	listshowTui "github.com/anondigriz/mogan-editor-cli/internal/tui/listshow"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/localfinder"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show local knowledge bases",
		Long:  `Show all knowledge base projects that are located in the base project directory`,
		Run: func(cmd *cobra.Command, args []string) {
			lf := localfinder.New(lg, *cfg)
			kbs := lf.FindInProjectsDir(cmd.Context())
			err := showKnowledgeBases(kbs)
			if err != nil {
				fmt.Printf("\n---\nFail to show list of the knowledge bases: %v\n", err)
				return
			}
		},
	}
)

func initShowCmd() {
	cobra.OnInitialize(initShowCmdCfg)
}

func initShowCmdCfg() {
}

func showKnowledgeBases(kbs []kbEnt.KnowledgeBase) error {
	list := make([]string, 0, len(kbs))
	for _, v := range kbs {
		list = append(list, fmt.Sprintf("name: %s; uuid: %s; path: %s", v.ShortName, v.ID, v.Path))
	}
	mt := listshowTui.New(list)

	if _, err := tea.NewProgram(mt).Run(); err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return err
	}
	return nil
}
