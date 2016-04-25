package mp4v2

import (
	"os"
)

type MP4File struct {
	mFile             *os.File
	mFileOriginalSize uint64
	mCreateFlags      uint64

	mPRootAtom *MP4Atom
}
