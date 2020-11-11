package general

import pstruct "github.com/pangine/pangineDSM-utils/program-struct"

// InstSuccessors calculates the deterministic successor(s) of input instruction
func InstSuccessors(inst pstruct.InstFlags, ip int) (successors []int) {
	successors = make([]int, 0)
	// The next offset
	if !(inst.IsJmp || inst.IsCall || inst.FlowStop) {
		successors = append(successors, ip)
	}
	// The jmp target
	if !inst.IsIndJmp &&
		(inst.IsJmp ||
			inst.IsCall ||
			inst.IsConditional) {
		successors = append(successors, ip+inst.JmpOffset)
	}
	return
}
