package choose

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"

	errMsgs "github.com/anondigriz/mogan-mini/cmd/mogan/cli/errors/messages"
	"github.com/anondigriz/mogan-mini/cmd/mogan/cli/messages"
	navigatorTUI "github.com/anondigriz/mogan-mini/internal/tui/group/navigator"
	kbsUC "github.com/anondigriz/mogan-mini/internal/usecase/knowledgebases"
)

func (c *Choose) Init(kbsu *kbsUC.KnowledgeBases, knowledgeBaseUUID string) error {
	parameters, err := kbsu.GetAllParameters(knowledgeBaseUUID)
	if err != nil {
		c.lg.Error(errMsgs.GetParametersFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetKnowledgeBaseFail)
		return err
	}
	c.Parameters = parameters

	rules, err := kbsu.GetAllRules(knowledgeBaseUUID)
	if err != nil {
		c.lg.Error(errMsgs.GetRulesFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetRulesFail)
		return err
	}
	c.Rules = rules

	groups, err := kbsu.GetAllGroups(knowledgeBaseUUID)
	if err != nil {
		c.lg.Error(errMsgs.GetKnowledgeBaseFail, zap.Error(err))
		messages.PrintFail(errMsgs.GetKnowledgeBaseFail)
		return err
	}
	c.Root = c.buildRootGroup(groups)

	return nil
}

func (c *Choose) buildRootGroup(groups map[string]kbEnt.Group) *navigatorTUI.Group {
	root := &navigatorTUI.Group{}

	var unusedParameters, unusedRules []string
	for k := range c.Parameters {
		unusedParameters = append(unusedParameters, k)

	}
	for k := range c.Rules {
		unusedRules = append(unusedRules, k)
	}

	for _, v := range groups {
		g, usedParameters, usedRules := c.buildChildGroup(v, root)
		root.Groups = append(root.Groups, g)
		for _, v := range usedParameters {
			unusedParameters = remove(unusedParameters, v)
		}
		for _, v := range usedRules {
			unusedRules = remove(unusedRules, v)
		}
	}

	_ = c.addParameters(root, unusedParameters, []string{})
	_ = c.addRules(root, unusedRules, []string{})

	return root
}

func remove(slice []string, s string) []string {
	index := slices.Index(slice, s)
	if index == -1 {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func (c *Choose) buildChildGroup(group kbEnt.Group, parent *navigatorTUI.Group) (*navigatorTUI.Group, []string, []string) {
	root := &navigatorTUI.Group{
		BaseInfo: group.BaseInfo,
		Parent:   parent,
	}
	var usedParameters, usedRules []string

	for _, v := range group.Groups {
		g, p, r := c.buildChildGroup(v, root)
		root.Groups = append(root.Groups, g)
		usedParameters = append(usedParameters, p...)
		usedRules = append(usedRules, r...)
	}
	usedParameters = c.addParameters(root, group.Parameters, usedParameters)
	usedRules = c.addRules(root, group.Rules, usedRules)
	return root, usedParameters, usedRules
}

func (c *Choose) addParameters(root *navigatorTUI.Group, parameters []string, usedParameters []string) []string {
	for _, v := range parameters {
		if p, ok := c.Parameters[v]; ok {
			root.Parameters = append(root.Parameters, &navigatorTUI.Parameter{
				BaseInfo: p.BaseInfo})
			usedParameters = append(usedParameters, v)
		}
	}
	return usedParameters
}

func (c *Choose) addRules(root *navigatorTUI.Group, rules []string, usedRules []string) []string {
	for _, v := range rules {
		if p, ok := c.Rules[v]; ok {
			root.Rules = append(root.Rules, &navigatorTUI.Rule{
				BaseInfo: p.BaseInfo})
			usedRules = append(usedRules, v)
		}
	}
	return usedRules
}
