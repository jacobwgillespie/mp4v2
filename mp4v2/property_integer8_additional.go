package mp4v2

// Read needs documentation
func (p *MP4Integer8Property) Read(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	p.mValues[index] = f.ReadUInt8()
}

// Write needs documentation
func (p *MP4Integer8Property) Write(f *MP4File, index uint32) {
	if p.mImplicit {
		return
	}
	f.WriteUIntSize(p.mValues[index])
}

// IncrementValue needs documentation
func (p *MP4Integer8Property) IncrementValue(increment int32, index uint32) {
	p.SetValue(uint8(int32(p.GetValue(0))+increment), 0)
}
