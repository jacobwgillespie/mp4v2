package mp4v2

type MP4RootAtom struct {
	MP4Atom

	mRewriteFtyp         *MP4FtypAtom
	mRewriteFtypPosition uint64
	mRewriteFree         *MP4FreeAtom
	mRewriteFreePosition uint64
}
