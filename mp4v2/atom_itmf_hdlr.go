package mp4v2

type MP4ItmfHdlrAtom struct {
	MP4Atom

	reserved1   MP4Integer32Property
	handlerType MP4BytesProperty
	reserved2   MP4BytesProperty
	name        MP4BytesProperty
}
