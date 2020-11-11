package programstruct

// Sections format of binary sections
type Sections struct {
	Name   []string
	Offset []int
	Data   []uint8
}

// ProgramHeader records header information written in binaries for execution
type ProgramHeader struct {
	PAddr int
	VAddr int
}

// BinaryInfo records original bytes in binary with operated headers
type BinaryInfo struct {
	ProgramHeaders []ProgramHeader
	Sections       Sections
}
