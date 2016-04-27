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

func NewMP4AtomInfo(name string, mandatory bool, onlyOne bool) *MP4AtomInfo {
	atomInfo := &MP4AtomInfo{
		mName:      name,
		mMandatory: mandatory,
		mOnlyOne:   onlyOne,
	}

	return atomInfo
}

type MP4AtomInterface interface{}

type MP4Atom struct {
	mFile          *MP4File
	mStart         uint64
	mEnd           uint64
	mLargesizeMode bool // true if largesize mode
	mSize          uint64
	mType          string
	mUnknownType   bool
	mExtendedType  []uint8

	mPParentAtom *MP4Atom
	mDepth       uint8

	mPProperties     []MP4PropertyInterface
	mPChildAtomInfos []*MP4AtomInfo
	mPChildAtoms     []*MP4Atom
}

func NewMP4Atom(file *MP4File, mtype string) *MP4Atom {
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

func (a *MP4Atom) GetFile() *MP4File {
	return a.mFile
}

func (a *MP4Atom) GetStart() uint64 {
	return a.mStart
}

func (a *MP4Atom) SetStart(val uint64) {
	a.mStart = val
}

func (a *MP4Atom) GetEnd() uint64 {
	return a.mEnd
}

func (a *MP4Atom) SetEnd(val uint64) {
	a.mEnd = val
}

func (a *MP4Atom) GetSize() uint64 {
	return a.mSize
}

func (a *MP4Atom) SetSize(val uint64) {
	a.mSize = val
}

func (a *MP4Atom) GetType() string {
	return a.mType
}

func (a *MP4Atom) SetType(val string) {
	a.mType = val
}

func (a *MP4Atom) GetExtendedType() []uint8 {
	return a.mExtendedType
}

func (a *MP4Atom) SetExtendedType(val []uint8) {
	a.mExtendedType = append([]uint8(nil), val...)
}

func (a *MP4Atom) IsUnknownType() bool {
	return a.mUnknownType
}

func (a *MP4Atom) SetUnknownType(val bool) {
	a.mUnknownType = val
}

func (a *MP4Atom) IsRootAtom() bool {
	return a.mType == ""
}

func (a *MP4Atom) GetParentAtom() *MP4Atom {
	return a.mPParentAtom
}

func (a *MP4Atom) SetParentAtom(val *MP4Atom) {
	a.mPParentAtom = val
}

func (a *MP4Atom) AddChildAtom(childAtom *MP4Atom) {
	childAtom.SetParentAtom(a)
	a.mPChildAtoms = append(a.mPChildAtoms, childAtom)
}

func (a *MP4Atom) InsertChildAtom(childAtom *MP4Atom, index int) {
	childAtom.SetParentAtom(a)
	a.mPChildAtoms = append(a.mPChildAtoms[:index], append([]*MP4Atom{childAtom}, a.mPChildAtoms[index:]...)...)
}

func (a *MP4Atom) DeleteChildAtom(childAtom *MP4Atom) {
	for idx, atom := range a.mPChildAtoms {
		if atom == childAtom {
			a.mPChildAtoms, a.mPChildAtoms[len(a.mPChildAtoms)-1] = append(a.mPChildAtoms[:idx], a.mPChildAtoms[idx+1:]...), nil
		}
	}
}

func (a *MP4Atom) GetNumberOfChildAtoms() int {
	return len(a.mPChildAtoms)
}

func (a *MP4Atom) GetChildAtom(idx int) *MP4Atom {
	return a.mPChildAtoms[idx]
}

func (a *MP4Atom) GetProperty(idx int) MP4PropertyInterface {
	return a.mPProperties[idx]
}

func (a *MP4Atom) GetCount() int {
	return len(a.mPProperties)
}

func (a *MP4Atom) AddProperty(property MP4PropertyInterface) {
	a.mPProperties = append(a.mPProperties, &property)
}

func (a *MP4Atom) AddVersionAndFlags() {
	a.AddProperty(NewMP4Integer8Property(a, "version"))
	a.AddProperty(NewMP4Integer24Property(a, "flags"))
}

func (a *MP4Atom) AddReserved(parentAtom MP4AtomInterface, name string, size int) {
	reserved := NewMP4BytesProperty(parentAtom, name, size, 0)
	reserved.SetReadOnly(true)
	a.AddProperty(reserved)
}

func (a *MP4Atom) ExpectChildAtom(name string, mandatory bool, onlyOne bool) {
	a.mPChildAtomInfos = append(a.mPChildAtomInfos, NewMP4AtomInfo(name, mandatory, onlyOne))
}

func (a *MP4Atom) Generate() {}
