package mp4v2

import "strconv"

// AtomID converts strings into atom IDs
func AtomID(typeVal string) uint32 {
	iVal, err := strconv.ParseUint(typeVal, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(iVal)
}

// IDAtom converts atom IDs into strings
func IDAtom(typeVal uint32) string {
	return strconv.FormatUint(uint64(typeVal), 10)
}
