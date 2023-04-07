package knowledgebase

import "time"

type BaseInfo struct {
	ID           string
	ShortName    string
	CreatedDate  time.Time
	ModifiedDate time.Time
}
