//go:generate genny -in=$GOFILE -out=array_bytes.go gen "ArrayName=MP4Bytes ArrayType=*uint8"
//go:generate genny -in=$GOFILE -out=array_float32.go gen "ArrayName=MP4Float32 ArrayType=float32"
//go:generate genny -in=$GOFILE -out=array_float64.go gen "ArrayName=MP4Float64 ArrayType=float64"
//go:generate genny -in=$GOFILE -out=array_integer16.go gen "ArrayName=MP4Integer16 ArrayType=uint16"
//go:generate genny -in=$GOFILE -out=array_integer32.go gen "ArrayName=MP4Integer32 ArrayType=uint32"
//go:generate genny -in=$GOFILE -out=array_integer64.go gen "ArrayName=MP4Integer64 ArrayType=uint64"
//go:generate genny -in=$GOFILE -out=array_integer8.go gen "ArrayName=MP4Integer8 ArrayType=uint8"
//go:generate genny -in=$GOFILE -out=array_string.go gen "ArrayName=MP4String ArrayType=string"

package mp4v2

import (
	"fmt"
	"math"

	"github.com/cheekybits/genny/generic"
)

// ArrayName is the generic name of the array
type ArrayName generic.Type

// ArrayType is the generic type of the array
type ArrayType generic.Type

// ArrayNameArray needs documentation
type ArrayNameArray struct {
	mElements       []ArrayType
	mNumElements    MP4ArrayIndex
	mMaxNumElements MP4ArrayIndex
}

// ValidIndex needs documentation
func (a *ArrayNameArray) ValidIndex(index MP4ArrayIndex) bool {
	return index < a.mNumElements
}

// Size needs documentation
func (a *ArrayNameArray) Size() MP4ArrayIndex {
	return a.mNumElements
}

// MaxSize needs documentation
func (a *ArrayNameArray) MaxSize() MP4ArrayIndex {
	return a.mMaxNumElements
}

// Add needs documentation
func (a *ArrayNameArray) Add(newElement ArrayType) {
	a.Insert(newElement, a.mNumElements)
}

// Insert needs documentation
func (a *ArrayNameArray) Insert(newElement ArrayType, newIndex MP4ArrayIndex) {
	if newIndex > a.mNumElements {
		panic("Illegal array index")
	}

	if a.mNumElements == a.mMaxNumElements {
		a.mMaxNumElements = MP4ArrayIndex(math.Max(float64(a.mMaxNumElements), 1) * 2)
		newMElements := make([]ArrayType, a.mMaxNumElements)
		copy(newMElements, a.mElements)
		a.mElements = newMElements
	}

	a.mElements[newIndex] = newElement
	a.mNumElements++
}

// Delete needs documentation
func (a *ArrayNameArray) Delete(index MP4ArrayIndex) {
	if !a.ValidIndex(index) {
		panic(fmt.Sprintf("Illegal array index: %v of %v", index, a.mNumElements))
	}

	a.mNumElements--
	if index < a.mNumElements {
		a.mElements = append(a.mElements[:index], a.mElements[index+1:]...)
	}
}

// Resize needs documentation
func (a *ArrayNameArray) Resize(newSize MP4ArrayIndex) {
	if newSize > a.mMaxNumElements {
		newMElements := make([]ArrayType, newSize)
		copy(newMElements, a.mElements)
		a.mElements = newMElements
	} else {
		a.mElements = a.mElements[:newSize]
	}

	a.mNumElements = newSize
	a.mMaxNumElements = newSize
}

// Get needs documentation
func (a *ArrayNameArray) Get(index MP4ArrayIndex) ArrayType {
	if a.ValidIndex(index) {
		return a.mElements[index]
	}

	panic(fmt.Sprintf("Illegal array index: %v of %v", index, a.mNumElements))
}

// Set needs documentation
func (a *ArrayNameArray) Set(index MP4ArrayIndex, value ArrayType) {
	if a.ValidIndex(index) {
		a.mElements[index] = value
	}

	panic(fmt.Sprintf("Illegal array index: %v of %v", index, a.mNumElements))
}
