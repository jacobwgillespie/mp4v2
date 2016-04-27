package mp4v2

type MP4FtypAtom struct {
	MP4Atom

	majorBrand       MP4StringProperty
	minorVersion     MP4Integer32Property
	compatibleBrands MP4StringProperty
}
