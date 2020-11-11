package objectsapi

import (
	mcclient "github.com/pangine/pangineDSM-utils/mcclient"
	pstruct "github.com/pangine/pangineDSM-utils/program-struct"
)

// Object is the interface to provide the same function APIs across different binary object types
type Object interface {
	FindObjectText(sec pstruct.Sections) (lo, hi int)
	TypeInst(insn string, l int) (flg pstruct.InstFlags)
	InstLstFixForPrefix(inque []int, bi pstruct.BinaryInfo) (outque []int)
	ParseObj(file string) (bi pstruct.BinaryInfo)
}

// MapAdd2Inst append map[int]bool to detailed InstFlags array
func MapAdd2Inst(bi pstruct.BinaryInfo, inMap map[int]bool, outList map[int]pstruct.InstFlags, object Object) {
	for k := range inMap {
		if _, ok := outList[k]; ok {
			// Already resolved
			continue
		}
		res := mcclient.SendResolve(pstruct.V2PConv(bi.ProgramHeaders, k), bi.Sections.Data)
		if !res.IsInst() || res.TakeBytes() == 0 {
			outList[k] = pstruct.InstFlags{NotValid: true}
			continue
		}
		inst, err := res.Inst()
		if err != nil {
			inst = "##INST"
		}
		outList[k] = object.TypeInst(inst, int(res.TakeBytes()))
	}
	return
}

// ListAdd2Inst append []int to detailed InstFlags array
func ListAdd2Inst(bi pstruct.BinaryInfo, inList []int, outList map[int]pstruct.InstFlags, object Object) {
	for _, k := range inList {
		if _, ok := outList[k]; ok {
			// Already resolved
			continue
		}
		res := mcclient.SendResolve(pstruct.V2PConv(bi.ProgramHeaders, k), bi.Sections.Data)
		if !res.IsInst() || res.TakeBytes() == 0 {
			outList[k] = pstruct.InstFlags{NotValid: true}
			continue
		}
		inst, err := res.Inst()
		if err != nil {
			inst = "##INST"
		}
		outList[k] = object.TypeInst(inst, int(res.TakeBytes()))
	}
	return
}
