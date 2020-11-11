package csapi

import (
	"io"
	"net"

	capnpsrc "github.com/pangine/pangineDSM-utils/capnp"
	capnp "zombiezen.com/go/capnproto2"
)

// RequestRTService is a wrapped version of RequestStage that does not contain rollback info
func RequestRTService(
	id int,
	offsets []int,
	socket string,
) {
	RequestStage(id, offsets, socket, false, 0)
}

// RequestStage is the function that generates StageAPI capnp message contains a request
// and send over target unix socket
func RequestStage(
	id int,
	offsets []int,
	socket string,
	rollbacked bool,
	rollbackedToID int,
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

	con.Write(RequestStageBuf(id, offsets, rollbacked, rollbackedToID))
}

// RequestStageBuf is the function that generates StageAPI capnp message contains a request
// and return the message in bytes
func RequestStageBuf(
	id int,
	offsets []int,
	rollbacked bool,
	rollbackedToID int,
) (
	buf []byte,
) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic("client message allocation error")
	}
	data, err := capnpsrc.NewRootStageAPI(seg)
	if err != nil {
		panic("client message root error")
	}
	CapnpOffsets, err := capnp.NewInt64List(seg, int32(len(offsets)))
	if err != nil {
		panic("client response offset list create error")
	}
	for i, o := range offsets {
		CapnpOffsets.Set(i, int64(o))
	}
	data.SetId(int64(id))
	data.SetOffsets(CapnpOffsets)
	data.SetRollbacked(rollbacked)
	data.SetRollbackedToId(int64(rollbackedToID))

	buf, err = msg.Marshal()
	if err != nil {
		panic("client message encode fail")
	}

	return
}

// TerminateClient is the function that generates StageAPI capnp message to terminate client
func TerminateClient(socket string) {
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
		panic("client message allocation error")
	}
	data, err := capnpsrc.NewRootStageAPI(seg)
	if err != nil {
		panic("client message root error")
	}
	data.SetTerminate(true)

	buf, err := msg.Marshal()
	if err != nil {
		panic("client message encode fail")
	}

	con.Write(buf)
}

// ReadStageMessage reads the StageAPI capnp message from input reader
func ReadStageMessage(
	reader io.Reader,
) (
	id int,
	terminate bool,
	offsets []int,
	rollbacked bool,
	rollbackedToID int,
) {
	msg, err := capnp.NewDecoder(reader).Decode()
	if err != nil {
		panic("client message decode error")
	}
	data, err := capnpsrc.ReadRootStageAPI(msg)
	if err != nil {
		panic("client message parsing error")
	}
	id = int(data.Id())
	terminate = data.Terminate()
	offsets = make([]int, 0)
	if data.HasOffsets() {
		CapnpOffsets, err := data.Offsets()
		if err != nil {
			panic("client message get offsets error")
		}
		for i := 0; i < CapnpOffsets.Len(); i++ {
			offsets = append(offsets, int(CapnpOffsets.At(i)))
		}
	}
	rollbacked = data.Rollbacked()
	rollbackedToID = int(data.RollbackedToId())
	return
}
