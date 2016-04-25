// +build ignore
//go:generate genny -in=$GOFILE -out=property_integer8_integer.go gen "GenericPropertyIntegerSize=8 GenericPropertyIntegerType=uint8"
//go:generate genny -in=$GOFILE -out=property_integer16_integer.go gen "GenericPropertyIntegerSize=16 GenericPropertyIntegerType=uint16"
//go:generate genny -in=$GOFILE -out=property_integer24_integer.go gen "GenericPropertyIntegerSize=24 GenericPropertyIntegerType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer32_integer.go gen "GenericPropertyIntegerSize=32 GenericPropertyIntegerType=uint32"
//go:generate genny -in=$GOFILE -out=property_integer64_integer.go gen "GenericPropertyIntegerSize=64 GenericPropertyIntegerType=uint64"

package mp4v2

import "github.com/cheekybits/genny/generic"

// MP4IntegerGenericPropertyIntegerSizeProperty is the generic parent struct
type MP4IntegerGenericPropertyIntegerSizeProperty generic.Type

// GenericPropertyIntegerSize is the generic integer size
type GenericPropertyIntegerSize generic.Type

// GenericPropertyIntegerType is the generic integer type
type GenericPropertyIntegerType generic.Type

// Read needs documentation
func (p *MP4IntegerGenericPropertyIntegerSizeProperty) Read(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	p.mValues[index] = f.ReadUIntGenericPropertyIntegerSize()
}

// Write needs documentation
func (p *MP4IntegerGenericPropertyIntegerSizeProperty) Write(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	f.WriteUIntGenericPropertyIntegerSize(p.mValues[index])
}

// IncrementValue needs documentation
func (p *MP4IntegerGenericPropertyIntegerSizeProperty) IncrementValue(increment int32, index uint32) {
	p.SetValue(GenericPropertyIntegerType(int32(p.GetValue(0))+increment), 0)
}
