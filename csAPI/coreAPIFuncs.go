package csapi

import (
	"io"
	"net"

	capnpsrc "github.com/pangine/pangineDSM-utils/capnp"
	capnp "zombiezen.com/go/capnproto2"
)

// CoreAPIType is the go version of capnp CoreAPIType enum
type CoreAPIType int

const (
	// CoreResponse represent that this message is a response with new instructions.
	CoreResponse = iota
	// CoreRequest represent that this message is a request to the instructions
	// of the function that the first element of the offset list is in.
	CoreRequest
	// CoreRollback represent that this message is a rollback requirement to the
	// to the offset in the list that was created in an earliest iteration.
	CoreRollback
)

// RespondCore is the function that generates CoreAPI capnp message contains a response
func RespondCore(
	id int,
	offsets []int,
	socket string,
) {
	addr, err := net.ResolveUnixAddr("unix", socket)
	if err != nil {
		panic("cannot solve server addr")
	}
	con, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("connect to resolver via tcp failed")
	}
	defer con.Close()

	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("server message allocation error")
	}
	data, err := capnpsrc.NewRootCoreAPI(seg)
	if err != nil {
		panic("server message root error")
	}
	capnpOffsets, err := capnp.NewInt64List(seg, int32(len(offsets)))
	if err != nil {
		panic("server offset list create error")
	}
	for i, o := range offsets {
		capnpOffsets.Set(i, int64(o))
	}
	data.SetId(int64(id))
	data.SetMessageType(capnpsrc.CoreAPIType_response)
	data.SetOffsets(capnpOffsets)

	buf, err := msg.Marshal()
	if err != nil {
		panic("server message encode fail")
	}

	con.Write(buf)
}

// RequestCore is the function that generates CoreAPI capnp message contains a request
// and read the reply insn information
func RequestCore(
	id int,
	offset int,
	socket string,
) (
	accepted bool,
	insns []int,
) {
	addr, err := net.ResolveUnixAddr("unix", socket)
	if err != nil {
		panic("cannot solve server addr")
	}
	con, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("connect to resolver via tcp failed")
	}
	defer con.Close()

	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("server message allocation error")
	}
	data, err := capnpsrc.NewRootCoreAPI(seg)
	if err != nil {
		panic("server message root error")
	}
	data.SetId(int64(id))
	data.SetMessageType(capnpsrc.CoreAPIType_request)
	capnpOffsets, err := capnp.NewInt64List(seg, 1)
	if err != nil {
		panic("server offset list create error")
	}
	capnpOffsets.Set(0, int64(offset))
	data.SetOffsets(capnpOffsets)

	buf, err := msg.Marshal()
	if err != nil {
		panic("server message encode fail")
	}

	con.Write(buf)

	// read reply
	var rid int
	rid, accepted, insns = ReadCoreAPIReplyMessage(con)
	if rid != id {
		accepted = false
	}
	return
}

// RollbackCore is the function that generates CoreAPI capnp message contains a rollback
// and read the reply insn information
func RollbackCore(
	id int,
	offsets []int,
	socket string,
) (
	nid int,
	accepted bool,
	insns []int,
) {
	addr, err := net.ResolveUnixAddr("unix", socket)
	if err != nil {
		panic("cannot solve server addr")
	}
	con, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("connect to resolver via tcp failed")
	}
	defer con.Close()

	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("server message allocation error")
	}
	data, err := capnpsrc.NewRootCoreAPI(seg)
	if err != nil {
		panic("server message root error")
	}
	capnpOffsets, err := capnp.NewInt64List(seg, int32(len(offsets)))
	if err != nil {
		panic("server offset list create error")
	}
	for i, o := range offsets {
		capnpOffsets.Set(i, int64(o))
	}
	data.SetId(int64(id))
	data.SetMessageType(capnpsrc.CoreAPIType_rollback)
	data.SetOffsets(capnpOffsets)

	buf, err := msg.Marshal()
	if err != nil {
		panic("server message encode fail")
	}

	con.Write(buf)

	// read reply
	nid, accepted, insns = ReadCoreAPIReplyMessage(con)
	return
}

// ReadCoreAPIReplyMessage reads the CoreAPIResponse capnp message from input reader
func ReadCoreAPIReplyMessage(
	reader io.Reader,
) (
	id int,
	accepted bool,
	offsets []int,
) {
	msg, err := capnp.NewDecoder(reader).Decode()
	if err != nil {
		panic("core reply message decode error")
	}
	data, err := capnpsrc.ReadRootCoreAPIReply(msg)
	if err != nil {
		panic("core reply message parsing error")
	}
	id = int(data.Id())
	accepted = data.Accepted()
	offsets = make([]int, 0)
	if data.HasOffsets() {
		capnpOffsets, err := data.Offsets()
		if err != nil {
			panic("client message get offsets error")
		}
		for i := 0; i < capnpOffsets.Len(); i++ {
			offsets = append(offsets, int(capnpOffsets.At(i)))
		}
	}
	return
}

// ReadCoreMessage reads the CoreAPI capnp message from input reader
func ReadCoreMessage(
	reader io.Reader,
) (
	id int,
	MessageType CoreAPIType,
	offsets []int,
) {
	msg, err := capnp.NewDecoder(reader).Decode()
	if err != nil {
		panic("server message decode error")
	}
	data, err := capnpsrc.ReadRootCoreAPI(msg)
	if err != nil {
		panic("server message parsing error")
	}
	id = int(data.Id())
	switch data.MessageType() {
	case capnpsrc.CoreAPIType_response:
		MessageType = CoreResponse
	case capnpsrc.CoreAPIType_request:
		MessageType = CoreRequest
	case capnpsrc.CoreAPIType_rollback:
		MessageType = CoreRollback
	}
	offsets = make([]int, 0)
	if data.HasOffsets() {
		capnpOffsets, err := data.Offsets()
		if err != nil {
			panic("server message get offsets error")
		}
		for i := 0; i < capnpOffsets.Len(); i++ {
			offsets = append(offsets, int(capnpOffsets.At(i)))
		}
	}
	return
}

// ReplyCoreAPIBuf is the function that generates CoreAPIReply capnp message
func ReplyCoreAPIBuf(
	id int,
	accepted bool,
	offsets []int,
) (
	buf []byte,
) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("core reply message allocation error")
	}
	data, err := capnpsrc.NewRootCoreAPIReply(seg)
	if err != nil {
		panic("core reply message root error")
	}
	capnpOffsets, err := capnp.NewInt64List(seg, int32(len(offsets)))
	if err != nil {
		panic("core reply message offset list create error")
	}
	for i, o := range offsets {
		capnpOffsets.Set(i, int64(o))
	}
	data.SetId(int64(id))
	data.SetAccepted(accepted)
	data.SetOffsets(capnpOffsets)

	buf, err = msg.Marshal()
	if err != nil {
		panic("core reply message encode fail")
	}

	return
}
