package mappers

import (
	"time"

	"github.com/anondigriz/mogan-mini/internal/entity/knowledgebase"
)

type BaseInfoForRow struct {
	UUID         string
	ID           string
	ShortName    string
	CreatedDate  int64
	ModifiedDate int64
}

func (bi *BaseInfoForRow) Fill(base knowledgebase.BaseInfo) {
	bi.UUID = base.UUID
	bi.ID = base.ID
	bi.ShortName = base.ShortName
	bi.CreatedDate = base.CreatedDate.UTC().Unix()
	bi.ModifiedDate = base.ModifiedDate.UTC().Unix()
}

func (bi *BaseInfoForRow) Extract() knowledgebase.BaseInfo {
	var b knowledgebase.BaseInfo
	b.UUID = bi.UUID
	b.ID = bi.ID
	b.ShortName = bi.ShortName
	b.CreatedDate = time.Unix(bi.CreatedDate, 0).UTC()
	b.ModifiedDate = time.Unix(bi.ModifiedDate, 0).UTC()
	return b
}
