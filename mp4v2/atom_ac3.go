package mp4v2

type MP4Ac3Atom struct {
	MP4Atom
}

func NewMP4Ac3Atom(file *MP4File) *MP4Ac3Atom {
	atom := &MP4Ac3Atom{
		MP4Atom: *NewMP4Atom(file, "ac-3"),
	}

	atom.AddReserved(*atom, "reserved1", 6)
	atom.AddProperty(NewMP4Integer16Property(atom, "dataReferenceIndex"))
	atom.AddReserved(*atom, "reserved2", 8)
	atom.AddProperty(NewMP4Integer16Property(atom, "channelCount"))
	atom.AddProperty(NewMP4Integer16Property(atom, "sampleSize"))
	atom.AddReserved(*atom, "reserved3", 4)
	atom.AddProperty(NewMP4Integer16Property(atom, "samplingRate"))
	atom.AddReserved(*atom, "reserved4", 2)

	atom.ExpectChildAtom("dac3", true, true)

	return atom
}

func (a *MP4Ac3Atom) Generate() {
	a.MP4Atom.Generate()

	a.mPProperties[1].(*MP4Integer16Property).SetValue(1, 0)
	a.mPProperties[3].(*MP4Integer16Property).SetValue(2, 0)
	a.mPProperties[4].(*MP4Integer16Property).SetValue(0x0010, 0)

	// The user should set the samplingRate as appropriate - and create the dac3 atom
}
