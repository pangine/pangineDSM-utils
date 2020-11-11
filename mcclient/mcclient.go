package mcclient

import (
	"fmt"
	"net"
	"os"

	capnpsrc "github.com/pangine/pangineDSM-utils/capnp"
	capnp "zombiezen.com/go/capnproto2"
)

var serverAddr = "/tmp/resolver.cpp.socket"

// SendResolve try to resolve data at idx and return resolved information
func SendResolve(idx int, data []byte) (res capnpsrc.ResolverOut) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("resolver in allocation error")
	}
	ri, err := capnpsrc.NewRootResolverIn(seg)
	if err != nil {
		panic("resolver in root error")
	}
	ri.SetTerminate(false)
	bl, err := capnp.NewUInt8List(seg, 32)
	if err != nil {
		panic("resolver in bytes error")
	}
	n := len(data) - idx
	if n > 32 {
		n = 32
	}
	for i := 0; i < n; i++ {
		bl.Set(i, data[i+idx])
	}
	ri.SetBytes(bl)
	inbuf, err := msg.Marshal()
	if err != nil {
		panic("resolver in encode error")
	}
	omsg, err := capnp.Unmarshal(sendBytes(inbuf))
	if err != nil {
		panic("resolver out decode error")
	}
	res, err = capnpsrc.ReadRootResolverOut(omsg)
	if err != nil {
		panic("resolver out root error")
	}
	return
}

// SendTerminate send a terminate signal to terminate the server
func SendTerminate() {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("resolver in allocation error")
	}
	ri, err := capnpsrc.NewRootResolverIn(seg)
	if err != nil {
		panic("resolver in root error")
	}
	ri.SetTerminate(true)
	if err != nil {
		panic("resolver in encode error")
	}
	inbuf, err := msg.Marshal()
	if err != nil {
		panic("resolver in encode error")
	}
	sendBytes(inbuf)
}

func sendBytes(inBuf []byte) (outBuf []byte) {
	addr, err := net.ResolveUnixAddr("unix", serverAddr)
	if err != nil {
		panic("cannot solve server addr")
	}

	con, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("connect to resolver via tcp failed")
	}
	defer con.Close()

	outBuf = make([]byte, 1024)
	con.Write(inBuf)
	i, err := con.Read(outBuf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read from server failed")
	} else {
		outBuf = outBuf[:i]
	}
	return
}
