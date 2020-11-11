package programstruct

// JmpBits specifies the bit length of a jump target
type JmpBits int

const (
	// Default is to keep the same with executable bits
	Default JmpBits = iota
	// Bits32 is 32 bits jmp target
	Bits32
	// Bits64 is 64 bits jmp target
	Bits64
)

// InstFlags record the conditions of the Inst
type InstFlags struct {
	NotValid      bool
	InstSize      int
	OriginInst    string
	IsConditional bool
	IsJmp         bool
	IsCall        bool
	IsRet         bool
	IsIndJmp      bool
	IsHlt         bool
	IsNop         bool
	FlowStop      bool
	JmpBits       JmpBits
	JmpOffset     int
	IndJmpTarget  string
	Prefixes      []string
}
