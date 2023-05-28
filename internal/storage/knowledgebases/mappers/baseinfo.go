package knowledgebase

import (
	"time"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
)

type BaseInfo struct {
	UUID         string
	ID           string
	ShortName    string
	Description  string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

func (bi *BaseInfo) Fill(base kbEnt.BaseInfo) {
	bi.UUID = base.UUID
	bi.ID = base.ID
	bi.ShortName = base.ShortName
	bi.CreatedDate = base.CreatedDate
	bi.ModifiedDate = base.ModifiedDate
}

func (bi BaseInfo) Extract() kbEnt.BaseInfo {
	result := kbEnt.BaseInfo{
		UUID:         bi.UUID,
		ID:           bi.ID,
		ShortName:    bi.ShortName,
		CreatedDate:  bi.CreatedDate,
		ModifiedDate: bi.ModifiedDate,
	}
	return result
}
