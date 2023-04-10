package args

import "io"

type UploadedFile interface {
	io.Reader
	io.Seeker
	io.Closer
}

type ImportKnowledgeBase struct {
	KnowledgeBaseUUID string
	XMLFile           UploadedFile
	FileName          string
}
