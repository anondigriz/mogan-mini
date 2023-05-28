package knowledgebase

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type Group struct {
	BaseInfo
	Groups     map[string]Group
	Parameters []string
	Rules      []string
}

func (g *Group) Fill(base kbEnt.Group) {
	g.BaseInfo.Fill(base.BaseInfo)
	g.Groups = make(map[string]Group, len(base.Groups))
	for k, v := range base.Groups {
		var child Group
		child.Fill(v)
		g.Groups[k] = child
	}
	g.Parameters = base.Parameters
	g.Rules = base.Rules
}

func (g Group) Extract() kbEnt.Group {
	result := kbEnt.Group{
		BaseInfo: g.BaseInfo.Extract(),
	}
	result.Groups = make(map[string]kbEnt.Group, len(g.Groups))
	for k, v := range g.Groups {
		result.Groups[k] = v.Extract()
	}
	result.Parameters = g.Parameters
	result.Rules = g.Rules
	return result
}
