package mp4v2

import (
	"log"
)

type MP4AtomInfo struct {
	mName      string
	mMandatory bool
	mOnlyOne   bool
	mCount     uint32
}

type MP4Atom struct {
	mFile          *MP4File
	mStart         uint64
	mEnd           uint64
	mLargesizeMode bool // true if largesize mode
	mSize          uint64
	mType          string
	mUnknownType   bool
	mExtendedType  [16]uint8

	mPParentAtom *MP4Atom
	mDepth       uint8

	mPChildAtomInfos []*MP4AtomInfo
	mPChildAtoms     []*MP4Atom
}

func CreateAtom(file *MP4File, mtype string) *MP4Atom {
	atom := &MP4Atom{
		mFile:  file,
		mType:  mtype,
		mDepth: 0xFF,
	}

	return atom
}

func ReadAtom(file *MP4File, pParentAtom *MP4Atom) *MP4Atom {
	var hdrSize uint8 = 8
	var extendedType [16]uint8

	var pos uint64 = file.GetPosition()

	log.Printf("\"%s\": pos = %v", file.GetFilename(), pos)
}
