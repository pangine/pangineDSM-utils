package programstruct

// P2VConv convert PA to VA according to
func P2VConv(ph []ProgramHeader, pAddr int) (vAddr int) {
	vAddr = pAddr
	largest := -1
	for _, e := range ph {
		if e.PAddr <= pAddr && e.PAddr > largest {
			largest = e.PAddr
			vAddr = pAddr - e.PAddr + e.VAddr
		}
	}
	return
}

// V2PConv convert VA to PA according to
func V2PConv(ph []ProgramHeader, vAddr int) (pAddr int) {
	pAddr = vAddr
	largest := -1
	for _, e := range ph {
		if e.VAddr <= vAddr && e.VAddr > largest {
			largest = e.VAddr
			pAddr = vAddr - e.VAddr + e.PAddr
		}
	}
	return
}

// VAisValid is to test whether the input vAddr is valid
func VAisValid(ph []ProgramHeader, vAddr int) bool {
	return P2VConv(ph, V2PConv(ph, vAddr)) == vAddr
}
