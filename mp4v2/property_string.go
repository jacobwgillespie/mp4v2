package mp4v2

type MP4StringProperty struct {
	MP4Property

	mFixedValueSize   int
	mDefaultValueSize int
	mValueSizes       []int
	mValues           []byte
}
