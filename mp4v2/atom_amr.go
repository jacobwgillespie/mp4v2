package mp4v2

type MP4AmrAtom struct {
	MP4Atom
}

func NewMP4AmrAtom(file *MP4File, atomType string) *MP4AmrAtom {
	atom := &MP4AmrAtom{
		MP4Atom: *NewMP4Atom(file, atomType),
	}

	atom.AddReserved(*atom, "reserved1", 6)
	atom.AddProperty(NewMP4Integer16Property(atom, "dataReferenceIndex"))
	atom.AddReserved(*atom, "reserved2", 16)
	atom.AddProperty(NewMP4Integer16Property(atom, "timeScale"))
	atom.AddReserved(*atom, "reserved3", 2)

	atom.ExpectChildAtom("damr", true, true)

	return atom
}

func (a *MP4AmrAtom) Generate() {
	a.MP4Atom.Generate()

	a.mPProperties[1].(*MP4Integer16Property).SetValue(1, 0)

	reserved2 := []uint8{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x02, 0x00, 0x10,
		0x00, 0x00, 0x00, 0x00,
	}

	a.mPProperties[2].SetReadOnly(false)
	a.mPProperties[2].(*MP4BytesProperty).SetValue(reserved2)
	a.mPProperties[2].SetReadOnly(true)

	// The user should set the samplingRate as appropriate - and create the dac3 atom
}
