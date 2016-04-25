package mp4v2

type MP4ArrayIndex uint32

type MP4Array interface {
	ValidIndex(index MP4ArrayIndex) bool
	Size() MP4ArrayIndex
	MaxSize() MP4ArrayIndex
}
