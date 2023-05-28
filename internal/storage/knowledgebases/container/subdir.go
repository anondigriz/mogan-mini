package container

import "path"

const (
	GroupsSubDir     = "groups"
	ParametersSubDir = "parameters"
	PatternsSubDir   = "patterns"
	RulesSubDir      = "rules"
)

func (c Container) getGroupsSubDir() string {
	return path.Join(c.knowledgeBaseDir, GroupsSubDir)
}

func (c Container) getParametersSubDir() string {
	return path.Join(c.knowledgeBaseDir, ParametersSubDir)
}

func (c Container) getPatternsSubDir() string {
	return path.Join(c.knowledgeBaseDir, PatternsSubDir)
}

func (c Container) getRulesSubDir() string {
	return path.Join(c.knowledgeBaseDir, RulesSubDir)
}
