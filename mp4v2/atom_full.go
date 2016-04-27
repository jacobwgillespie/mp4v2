package mp4v2

type MP4FullAtom struct {
	MP4Atom
	version MP4Integer8Property
	flags   MP4Integer24Property
}
