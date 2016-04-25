//go:generate genny -in=$GOFILE -out=property_integer8.go gen "PropName=Integer8 PropType=uint8"
//go:generate genny -in=$GOFILE -out=property_integer16.go gen "PropName=Integer16 PropType=uint16"
//go:generate genny -in=$GOFILE -out=property_integer24.go gen "PropName=Integer24 PropType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer32.go gen "PropName=Integer32 PropType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer64.go gen "PropName=Integer64 PropType=uint64"

package mp4v2

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// PropName is the generic integer size
type PropName generic.Type

// PropType is the generic integer type
type PropType generic.Type

// MP4PropNameProperty encapsulates a property of type PropType
type MP4PropNameProperty struct {
	mParentAtom *MP4Atom
	mName       string
	mReadOnly   bool
	mImplicit   bool
	mValues     []PropType
}

// NewMP4PropNameProperty creates a new MP4PropNameProperty
func NewMP4PropNameProperty(parentAtom *MP4Atom, name string) *MP4PropNameProperty {
	prop := &MP4PropNameProperty{
		mParentAtom: parentAtom,
		mName:       name,
	}

	prop.SetCount(1)
	prop.SetValue(0, 0)

	return prop
}

// GetParentAtom returns the parent atom for the property
func (p *MP4PropNameProperty) GetParentAtom() *MP4Atom {
	return p.mParentAtom
}

// GetName returns the property name
func (p *MP4PropNameProperty) GetName() string {
	return p.mName
}

// GetType returns the property type
func (p *MP4PropNameProperty) GetType() MP4PropertyType {
	return PropNameProperty
}

// IsReadOnly returns true if the property is read-only
func (p *MP4PropNameProperty) IsReadOnly() bool {
	return p.mReadOnly
}

// SetReadOnly enables setting the property to read-only
func (p *MP4PropNameProperty) SetReadOnly(value ...bool) {
	if len(value) == 0 {
		p.mReadOnly = true
	} else {
		p.mReadOnly = value[0]
	}
}

// IsImplicit returns true if the property is implicit
func (p *MP4PropNameProperty) IsImplicit() bool {
	return p.mImplicit
}

// SetImplicit sets the property to implicit
func (p *MP4PropNameProperty) SetImplicit(value ...bool) {
	if len(value) == 0 {
		p.mImplicit = true
	} else {
		p.mImplicit = value[0]
	}
}

// GetCount needs documentation
func (p *MP4PropNameProperty) GetCount() uint32 {
	return uint32(len(p.mValues))
}

// SetCount needs documentation
func (p *MP4PropNameProperty) SetCount(count uint32) {
	if count > uint32(len(p.mValues)) {
		newValues := make([]PropType, count)
		copy(newValues, p.mValues)
		p.mValues = newValues
	} else {
		p.mValues = p.mValues[:count]
	}
}

// FindProperty needs documentation
func (p *MP4PropNameProperty) FindProperty(name string, ppProperty **MP4PropNameProperty, pIndex *uint32) bool {
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
func (p *MP4PropNameProperty) GetValue(index uint32) PropType {
	return p.mValues[index]
}

// SetValue needs documentation
func (p *MP4PropNameProperty) SetValue(value PropType, index uint32) {
	if p.mReadOnly {
		panic(fmt.Sprintf("property is read-only: %s", p.mName))
	}

	p.mValues[index] = value
}

// AddValue needs documentation
func (p *MP4PropNameProperty) AddValue(value PropType) {
	p.mValues = append(p.mValues, value)
}

// InsertValue needs documentation
func (p *MP4PropNameProperty) InsertValue(value PropType, index uint32) {
	if index > uint32(len(p.mValues)) {
		panic(fmt.Sprintf("invalid index: %v", index))
	}

	p.mValues = append(p.mValues, 0)
	copy(p.mValues[index+1:], p.mValues[index:])
	p.mValues[index] = value
}

// DeleteValue needs documentation
func (p *MP4PropNameProperty) DeleteValue(index uint32) {
	p.mValues = append(p.mValues[:index], p.mValues[index+1:]...)
}
