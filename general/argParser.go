package general

import (
	"flag"
)

// Args is the parsed version of input arguments
type Args struct {
	File        string
	V           bool
	StageSocket string
	RtSocket    string
	CoreSocket  string
}

// ArgParser parses exe args to a uniformed structure
func ArgParser() (outarg Args) {
	file := flag.String("f", "", "Input file")
	v := flag.Bool("v", false, "verbose")
	sso := flag.String("ss", "/tmp/PangineStage.socket", "Unix socket file for this stage")
	rso := flag.String("rs", "/tmp/PangineRt.socket", "Unix socket file for recursive traversal process")
	cso := flag.String("cs", "/tmp/PangineCore.socket", "Unix socket file for core state process")
	flag.Parse()
	outarg.File = *file
	outarg.V = *v
	outarg.StageSocket = *sso
	outarg.RtSocket = *rso
	outarg.CoreSocket = *cso
	return
}
