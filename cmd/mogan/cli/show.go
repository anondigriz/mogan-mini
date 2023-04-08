package cli

import (
	"fmt"

	showTui "github.com/anondigriz/mogan-editor-cli/internal/tui/show"
	"github.com/anondigriz/mogan-editor-cli/internal/utility/knowledgebase/show"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show knowledge bases in the workspace",
		Long:  `Show all knowledge base projects that are located in the base project directory`,
		Run: func(cmd *cobra.Command, args []string) {
			sw := show.New(lg, cfg)
			kbs := sw.FindProjects(cmd.Context())
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

func showKnowledgeBases(kbs []show.KnowledgeBase) error {
	m := showTui.ListModel{}
	for _, v := range kbs {
		m.List = append(m.List, fmt.Sprintf("id: %s, name: %s, path: %s", v.Id, v.Name, v.Path))
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		lg.Error("Alas, there's been an error: %v", zap.Error(err))
		return err
	}
	return nil
}
