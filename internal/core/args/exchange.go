package args

import "io"

type ImportKnowledgeBase struct {
	XMLFile  io.ReadSeekCloser
	FileName string
}
