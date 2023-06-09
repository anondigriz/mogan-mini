package choose

import (
	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type Group struct {
	kbEnt.BaseInfo
	Parent     *Group
	Groups     []*Group
	Parameters []*Parameter
	Rules      []*Rule
}

type Parameter struct {
	kbEnt.BaseInfo
}

type Rule struct {
	kbEnt.BaseInfo
}
