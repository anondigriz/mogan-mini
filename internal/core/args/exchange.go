package args

import "io"

type ImportProject struct {
	XMLFile  io.ReadSeekCloser
	FileName string
}
