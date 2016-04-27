package mp4v2

type MP4PropertyType int

const (
	Integer8Property MP4PropertyType = iota
	Integer16Property
	Integer24Property
	Integer32Property
	Integer64Property
	Float32Property
	StringProperty
	BytesProperty
	TableProperty
	DescriptorProperty
	LanguageCodeProperty
	BasicTypeProperty
	PropNameProperty = BasicTypeProperty
)

type MP4PropertyInterface interface {
	// GetParentAtom() *MP4Atom
	// GetName() string
	// GetType() MP4PropertyType
	// IsReadOnly() bool
	SetReadOnly(value ...bool)
	// IsImplicit() bool
	// SetImplicit(value ...bool)
	// GetCount() uint32
	// SetCount(count uint32)
	// Generate()
	// Read(file *MP4File, index uint32)
	// Write(file *MP4File, index uint32)
	// Dump(indent uint8, dumpImplicits bool, index uint32)
	// FindProperty(name string, ppProperty **MP4Property, pIndex *uint32) bool
}

type MP4Property struct {
	mParentAtom MP4AtomInterface
	mName       string
	mReadOnly   bool
	mImplicit   bool
}

func NewMP4Property(parentAtom *MP4Atom, name string) *MP4Property {
	return &MP4Property{
		mParentAtom: parentAtom,
		mName:       name,
	}
}

func (p *MP4Property) GetParentAtom() MP4AtomInterface {
	return p.mParentAtom
}

func (p *MP4Property) GetName() string {
	return p.mName
}

func (p *MP4Property) GetType() MP4PropertyType {
	return 0
}

func (p *MP4Property) IsReadOnly() bool {
	return p.mReadOnly
}

func (p *MP4Property) SetReadOnly(value ...bool) {
	if len(value) == 0 {
		p.mReadOnly = true
	} else {
		p.mReadOnly = value[0]
	}
}

func (p *MP4Property) IsImplicit() bool {
	return p.mImplicit
}

func (p *MP4Property) SetImplicit(value ...bool) {
	if len(value) == 0 {
		p.mImplicit = true
	} else {
		p.mImplicit = value[0]
	}
}

func (p *MP4Property) GetCount() uint32 {
	return 0
}

func (p *MP4Property) SetCount(count uint32) {
	// Do nothing here
}

func (p *MP4Property) Generate() {
	// Do nothing here
}

func (p *MP4Property) Read(file *MP4File, index uint32) {
	// Do nothing here
}

func (p *MP4Property) Write(file *MP4File, index uint32) {
	// Do nothing here
}

func (p *MP4Property) Dump(indent uint8, dumpImplicits bool, index uint32) {
	// Do nothing here
}

func (p *MP4Property) FindProperty(name string, ppProperty **MP4Property, pIndex *uint32) bool {
	if name == "" {
		return false
	}

	if p.mName == name {
		*ppProperty = p
		return true
	}
	return false
}
