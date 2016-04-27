package mp4v2

type MP4DataAtom struct {
	MP4Atom

	typeReserved      MP4Integer16Property
	typeSetIdentifier MP4Integer8Property
	typeCode          MP4BasicTypeProperty
	locale            MP4Integer32Property
	metadata          MP4BytesProperty
}
