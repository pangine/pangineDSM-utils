package general

import (
	"fmt"
	"os"
	"strings"
)

// LlvmTripleStruct is the llvm triple in format of a struct of string
type LlvmTripleStruct struct {
	Arch    string
	SubArch string
	Vendor  string
	OS      string
	Env     string
	Obj     string
}

// ParseLlvmTriple is used to convert string version of llvm triple into the struct version
func ParseLlvmTriple(tripleStr string) (tripleStruct LlvmTripleStruct) {
	fields := strings.Split(tripleStr, "-")
	if len(fields) != 5 && len(fields) != 6 {
		panic("Input llvm triple illegal: " + tripleStr)
	}
	tripleStruct.Arch = fields[0]
	if len(fields) == 5 {
		// This llvm triple has a sub arch field
		tripleStruct.SubArch = fields[1]
	}
	tripleStruct.Vendor = fields[len(fields)-4]
	tripleStruct.OS = fields[len(fields)-3]
	tripleStruct.Env = fields[len(fields)-2]
	tripleStruct.Obj = fields[len(fields)-1]
	return
}

//PrintSupportLlvmTriple is a print function that print input string list with format
func PrintSupportLlvmTriple(supportTripleStr []string) {
	fmt.Println("There are supported llvm triples for this tool:")
	for _, triple := range supportTripleStr {
		fmt.Printf("\t%s\n", triple)
	}
	os.Exit(0)
}

//CheckLlvmTriple is to check whether the input LLVM triple is supported
func CheckLlvmTriple(inputTripleStr string, supportTripleStr []string) (outputTripleStr string) {
	for _, t := range supportTripleStr {
		if strings.ToLower(inputTripleStr) == strings.ToLower(t) {
			outputTripleStr = t
			break
		}
	}
	if outputTripleStr == "" {
		fmt.Printf("ERROR: input llvm triple: %s is not supported", inputTripleStr)
		os.Exit(1)
	}
	return
}
