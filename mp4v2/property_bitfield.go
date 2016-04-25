package mp4v2

import "fmt"

// MP4BitfieldProperty encapsulates a property of type uint64
type MP4BitfieldProperty struct {
	mParentAtom *MP4Atom
	mName       string
	mReadOnly   bool
	mImplicit   bool
	mValues     []uint64
	mNumBits    uint8
}

// NewMP4BitfieldProperty creates a new MP4BitfieldProperty
func NewMP4BitfieldProperty(parentAtom *MP4Atom, name string, numBits uint8) *MP4BitfieldProperty {
	if numBits == 0 || numBits > 64 {
		panic("Invalid number of bits")
	}

	prop := &MP4BitfieldProperty{
		mParentAtom: parentAtom,
		mName:       name,
		mNumBits:    numBits,
	}

	prop.SetCount(1)
	prop.SetValue(0, 0)

	return prop
}

// GetNumBits returns the number of bits
func (p *MP4BitfieldProperty) GetNumBits() uint8 {
	return p.mNumBits
}

// SetNumBits sets the number of bits
func (p *MP4BitfieldProperty) SetNumBits(numBits uint8) {
	p.mNumBits = numBits
}

// GetParentAtom returns the parent atom for the property
func (p *MP4BitfieldProperty) GetParentAtom() *MP4Atom {
	return p.mParentAtom
}

// GetName returns the property name
func (p *MP4BitfieldProperty) GetName() string {
	return p.mName
}

// GetType returns the property type
func (p *MP4BitfieldProperty) GetType() MP4PropertyType {
	return Integer64Property
}

// IsReadOnly returns true if the property is read-only
func (p *MP4BitfieldProperty) IsReadOnly() bool {
	return p.mReadOnly
}

// SetReadOnly enables setting the property to read-only
func (p *MP4BitfieldProperty) SetReadOnly(value ...bool) {
	if len(value) == 0 {
		p.mReadOnly = true
	} else {
		p.mReadOnly = value[0]
	}
}

// IsImplicit returns true if the property is implicit
func (p *MP4BitfieldProperty) IsImplicit() bool {
	return p.mImplicit
}

// SetImplicit sets the property to implicit
func (p *MP4BitfieldProperty) SetImplicit(value ...bool) {
	if len(value) == 0 {
		p.mImplicit = true
	} else {
		p.mImplicit = value[0]
	}
}

// GetCount needs documentation
func (p *MP4BitfieldProperty) GetCount() uint32 {
	return uint32(len(p.mValues))
}

// SetCount needs documentation
func (p *MP4BitfieldProperty) SetCount(count uint32) {
	if count > uint32(len(p.mValues)) {
		newValues := make([]uint64, count)
		copy(newValues, p.mValues)
		p.mValues = newValues
	} else {
		p.mValues = p.mValues[:count]
	}
}

// Generate needs documentation
func (p *MP4BitfieldProperty) Generate() {
	// Do nothing here
}

// Read needs documentation
func (p *MP4BitfieldProperty) Read(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	p.mValues[index] = f.ReadBits(p.mNumBits)
}

// Write needs documentation
func (p *MP4BitfieldProperty) Write(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	f.WriteBits(p.mValues[index], p.mNumBits)
}

// Dump needs documentation
func (p *MP4BitfieldProperty) Dump(indent uint8, dumpImplicits bool, index uint32) {
	// Do nothing here
}

// FindProperty needs documentation
func (p *MP4BitfieldProperty) FindProperty(name string, ppProperty **MP4BitfieldProperty, pIndex *uint32) bool {
	if name == "" {
		return false
	}

	if p.mName == name {
		*ppProperty = p
		return true
	}
	return false
}

// GetValue needs documentation
func (p *MP4BitfieldProperty) GetValue(index uint32) uint64 {
	return p.mValues[index]
}

// SetValue needs documentation
func (p *MP4BitfieldProperty) SetValue(value uint64, index uint32) {
	if p.mReadOnly {
		panic(fmt.Sprintf("property is read-only: %s", p.mName))
	}

	p.mValues[index] = value
}

// AddValue needs documentation
func (p *MP4BitfieldProperty) AddValue(value uint64) {
	p.mValues = append(p.mValues, value)
}

// InsertValue needs documentation
func (p *MP4BitfieldProperty) InsertValue(value uint64, index uint32) {
	if value > uint64(len(p.mValues)) {
		panic(fmt.Sprintf("invalid index: %v", index))
	}

	p.mValues = append(p.mValues, 0)
	copy(p.mValues[index+1:], p.mValues[index:])
	p.mValues[index] = value
}

// DeleteValue needs documentation
func (p *MP4BitfieldProperty) DeleteValue(index uint32) {
	p.mValues = append(p.mValues[:index], p.mValues[index+1:]...)
}

// IncrementValue needs documentation
func (p *MP4BitfieldProperty) IncrementValue(increment int32, index uint32) {
	p.SetValue(uint64(int32(p.GetValue(0))+increment), 0)
}
