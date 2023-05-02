package args

import "io"

type UploadedFile interface {
	io.Reader
	io.Seeker
	io.Closer
}

type ImportProject struct {
	XMLFile  UploadedFile
	FileName string
}
