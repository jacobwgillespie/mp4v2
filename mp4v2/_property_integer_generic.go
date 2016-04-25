//go:generate genny -in=$GOFILE -out=property_integer8.go gen "IPropSize=8 IPropType=uint8"
//go:generate genny -in=$GOFILE -out=property_integer16.go gen "IPropSize=16 IPropType=uint16"
//go:generate genny -in=$GOFILE -out=property_integer24.go gen "IPropSize=24 IPropType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer32.go gen "IPropSize=32 IPropType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer64.go gen "IPropSize=64 IPropType=uint64"

package mp4v2

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// IPropSize is the generic integer size
type IPropSize generic.Type

// IPropType is the generic integer type
type IPropType generic.Type

// MP4IntegerIPropSizeProperty encapsulates a property of type IPropType
type MP4IntegerIPropSizeProperty struct {
	mParentAtom *MP4Atom
	mName       string
	mReadOnly   bool
	mImplicit   bool
	mValues     []IPropType
}

// NewMP4IntegerIPropSizeProperty creates a new MP4IntegerIPropSizeProperty
func NewMP4IntegerIPropSizeProperty(parentAtom *MP4Atom, name string) *MP4IntegerIPropSizeProperty {
	prop := &MP4IntegerIPropSizeProperty{
		mParentAtom: parentAtom,
		mName:       name,
	}

	prop.SetCount(1)
	prop.SetValue(0, 0)

	return prop
}

// GetParentAtom returns the parent atom for the property
func (p *MP4IntegerIPropSizeProperty) GetParentAtom() *MP4Atom {
	return p.mParentAtom
}

// GetName returns the property name
func (p *MP4IntegerIPropSizeProperty) GetName() string {
	return p.mName
}

// GetType returns the property type
func (p *MP4IntegerIPropSizeProperty) GetType() MP4PropertyType {
	return IntegerIPropSizeProperty
}

// IsReadOnly returns true if the property is read-only
func (p *MP4IntegerIPropSizeProperty) IsReadOnly() bool {
	return p.mReadOnly
}

// SetReadOnly enables setting the property to read-only
func (p *MP4IntegerIPropSizeProperty) SetReadOnly(value ...bool) {
	if len(value) == 0 {
		p.mReadOnly = true
	} else {
		p.mReadOnly = value[0]
	}
}

// IsImplicit returns true if the property is implicit
func (p *MP4IntegerIPropSizeProperty) IsImplicit() bool {
	return p.mImplicit
}

// SetImplicit sets the property to implicit
func (p *MP4IntegerIPropSizeProperty) SetImplicit(value ...bool) {
	if len(value) == 0 {
		p.mImplicit = true
	} else {
		p.mImplicit = value[0]
	}
}

// GetCount needs documentation
func (p *MP4IntegerIPropSizeProperty) GetCount() uint32 {
	return uint32(len(p.mValues))
}

// SetCount needs documentation
func (p *MP4IntegerIPropSizeProperty) SetCount(count uint32) {
	if count > uint32(len(p.mValues)) {
		newValues := make([]IPropType, count)
		copy(newValues, p.mValues)
		p.mValues = newValues
	} else {
		p.mValues = p.mValues[:count]
	}
}

// Generate needs documentation
func (p *MP4IntegerIPropSizeProperty) Generate() {
	// Do nothing here
}

// Read needs documentation
func (p *MP4IntegerIPropSizeProperty) Read(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	p.mValues[index] = f.ReadUIntIPropSize()
}

// Write needs documentation
func (p *MP4IntegerIPropSizeProperty) Write(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	f.WriteUIntSize(p.mValues[index])
}

// Dump needs documentation
func (p *MP4IntegerIPropSizeProperty) Dump(indent uint8, dumpImplicits bool, index uint32) {
	// Do nothing here
}

// FindProperty needs documentation
func (p *MP4IntegerIPropSizeProperty) FindProperty(name string, ppProperty **MP4IntegerIPropSizeProperty, pIndex *uint32) bool {
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
func (p *MP4IntegerIPropSizeProperty) GetValue(index uint32) IPropType {
	return p.mValues[index]
}

// SetValue needs documentation
func (p *MP4IntegerIPropSizeProperty) SetValue(value IPropType, index uint32) {
	if p.mReadOnly {
		panic(fmt.Sprintf("property is read-only: %s", p.mName))
	}

	p.mValues[index] = value
}

// AddValue needs documentation
func (p *MP4IntegerIPropSizeProperty) AddValue(value IPropType) {
	p.mValues = append(p.mValues, value)
}

// InsertValue needs documentation
func (p *MP4IntegerIPropSizeProperty) InsertValue(value IPropType, index uint32) {
	if value > IPropType(len(p.mValues)) {
		panic(fmt.Sprintf("invalid index: %v", index))
	}

	p.mValues = append(p.mValues, 0)
	copy(p.mValues[index+1:], p.mValues[index:])
	p.mValues[index] = value
}

// DeleteValue needs documentation
func (p *MP4IntegerIPropSizeProperty) DeleteValue(index uint32) {
	p.mValues = append(p.mValues[:index], p.mValues[index+1:]...)
}

// IncrementValue needs documentation
func (p *MP4IntegerIPropSizeProperty) IncrementValue(increment int32, index uint32) {
	p.SetValue(IPropType(int32(p.GetValue(0))+increment), 0)
}
