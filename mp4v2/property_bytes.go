package mp4v2

type MP4BytesProperty struct {
	MP4Property

	mFixedValueSize   int
	mDefaultValueSize int
	mValueSizes       []int32
	mValues           []byte
}

// NewMP4BytesProperty needs documentation
// valueSize - 0
// defaultValueSize - 0
func NewMP4BytesProperty(parentAtom MP4AtomInterface, name string, valueSize int, defaultValueSize int) *MP4BytesProperty {
	property := &MP4BytesProperty{
		MP4Property: MP4Property{
			mParentAtom: &parentAtom,
			mName:       name,
		},
		mDefaultValueSize: defaultValueSize,
	}

	property.SetCount(1)

	return property
}

func (p *MP4BytesProperty) GetType() MP4PropertyType {
	return BytesProperty
}

func (p *MP4BytesProperty) GetCount() int {
	return len(p.mValues)
}

func (p *MP4BytesProperty) SetValue(val []byte) {
	p.mValues = make([]byte, len(val))
	copy(p.mValues, val)
}
