// Code generated by capnpc-go. DO NOT EDIT.

package pangine

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// output of container parser
type CPOut struct{ capnp.Struct }

// CPOut_TypeID is the unique identifier for the type CPOut.
const CPOut_TypeID = 0xb60d2ba29728e68f

func NewCPOut(s *capnp.Segment) (CPOut, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return CPOut{st}, err
}

func NewRootCPOut(s *capnp.Segment) (CPOut, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return CPOut{st}, err
}

func ReadRootCPOut(msg *capnp.Message) (CPOut, error) {
	root, err := msg.RootPtr()
	return CPOut{root.Struct()}, err
}

func (s CPOut) String() string {
	str, _ := text.Marshal(0xb60d2ba29728e68f, s.Struct)
	return str
}

func (s CPOut) Name() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s CPOut) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CPOut) SetName(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewName sets the name field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s CPOut) NewName(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s CPOut) Offset() (capnp.Int64List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Int64List{List: p.List()}, err
}

func (s CPOut) HasOffset() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CPOut) SetOffset(v capnp.Int64List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewOffset sets the offset field to a newly
// allocated capnp.Int64List, preferring placement in s's segment.
func (s CPOut) NewOffset(n int32) (capnp.Int64List, error) {
	l, err := capnp.NewInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int64List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s CPOut) Bytes() (capnp.UInt8List, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.UInt8List{List: p.List()}, err
}

func (s CPOut) HasBytes() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s CPOut) SetBytes(v capnp.UInt8List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewBytes sets the bytes field to a newly
// allocated capnp.UInt8List, preferring placement in s's segment.
func (s CPOut) NewBytes(n int32) (capnp.UInt8List, error) {
	l, err := capnp.NewUInt8List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt8List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s CPOut) Loads() (HeaderLoad_List, error) {
	p, err := s.Struct.Ptr(3)
	return HeaderLoad_List{List: p.List()}, err
}

func (s CPOut) HasLoads() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s CPOut) SetLoads(v HeaderLoad_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewLoads sets the loads field to a newly
// allocated HeaderLoad_List, preferring placement in s's segment.
func (s CPOut) NewLoads(n int32) (HeaderLoad_List, error) {
	l, err := NewHeaderLoad_List(s.Struct.Segment(), n)
	if err != nil {
		return HeaderLoad_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

// CPOut_List is a list of CPOut.
type CPOut_List struct{ capnp.List }

// NewCPOut creates a new list of CPOut.
func NewCPOut_List(s *capnp.Segment, sz int32) (CPOut_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return CPOut_List{l}, err
}

func (s CPOut_List) At(i int) CPOut { return CPOut{s.List.Struct(i)} }

func (s CPOut_List) Set(i int, v CPOut) error { return s.List.SetStruct(i, v.Struct) }

func (s CPOut_List) String() string {
	str, _ := text.MarshalList(0xb60d2ba29728e68f, s.List)
	return str
}

// CPOut_Promise is a wrapper for a CPOut promised by a client call.
type CPOut_Promise struct{ *capnp.Pipeline }

func (p CPOut_Promise) Struct() (CPOut, error) {
	s, err := p.Pipeline.Struct()
	return CPOut{s}, err
}

// load infomation for mapping PA and VA
type HeaderLoad struct{ capnp.Struct }

// HeaderLoad_TypeID is the unique identifier for the type HeaderLoad.
const HeaderLoad_TypeID = 0xa85299c27791f521

func NewHeaderLoad(s *capnp.Segment) (HeaderLoad, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return HeaderLoad{st}, err
}

func NewRootHeaderLoad(s *capnp.Segment) (HeaderLoad, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return HeaderLoad{st}, err
}

func ReadRootHeaderLoad(msg *capnp.Message) (HeaderLoad, error) {
	root, err := msg.RootPtr()
	return HeaderLoad{root.Struct()}, err
}

func (s HeaderLoad) String() string {
	str, _ := text.Marshal(0xa85299c27791f521, s.Struct)
	return str
}

func (s HeaderLoad) Physical() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s HeaderLoad) SetPhysical(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s HeaderLoad) Virtual() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s HeaderLoad) SetVirtual(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

// HeaderLoad_List is a list of HeaderLoad.
type HeaderLoad_List struct{ capnp.List }

// NewHeaderLoad creates a new list of HeaderLoad.
func NewHeaderLoad_List(s *capnp.Segment, sz int32) (HeaderLoad_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return HeaderLoad_List{l}, err
}

func (s HeaderLoad_List) At(i int) HeaderLoad { return HeaderLoad{s.List.Struct(i)} }

func (s HeaderLoad_List) Set(i int, v HeaderLoad) error { return s.List.SetStruct(i, v.Struct) }

func (s HeaderLoad_List) String() string {
	str, _ := text.MarshalList(0xa85299c27791f521, s.List)
	return str
}

// HeaderLoad_Promise is a wrapper for a HeaderLoad promised by a client call.
type HeaderLoad_Promise struct{ *capnp.Pipeline }

func (p HeaderLoad_Promise) Struct() (HeaderLoad, error) {
	s, err := p.Pipeline.Struct()
	return HeaderLoad{s}, err
}

const schema_9c2ec473614d6dbe = "x\xdal\xd1=\x8b\x13A\x00\xc6\xf1\xe7\x99\xd9\xbdX" +
	"\x18\xcd`J!cw\xbepp\x87\x85\\\xa3\xd1\xc6" +
	"\x13\xc5\x8c\x85\x95(c\xb2\xab\x0b\xc9\xce\xb0;\xeb\x99" +
	"\xca\xc6F\x11D,|AQ\xac\xee\x03\x88VZ\x1c" +
	"\xfa\x15\x04\xed\xae\xd1O`\xbd\xb2\xd1\xe4,\xae\x9b\xf9" +
	"\xf3\x14?f:[g\x84\x8a?\x02f_\xbcT\x1f" +
	"\xf9\xfdds\xfb\xc5\x95-\x986E\xfdyr\xc9\x96" +
	"_V^!j\x01\xaa\xbf\xad6Zj\xa3\xa7\xa6\x9b" +
	"`\xfd\xf8\xe7\xf2\xb3w\xc7\xdb\x1f\xa0\xda\xdc\x1d\xc6\xb3" +
	"\xe5\xf7\x97j\xa7\xa5vz\x87\xda\xfc\x05\xd6C\x7f\xc3" +
	"Uae(\xac\xcf\xfd\xfa\xf9\xc4\x8e\x92\xe2\xa2\xb3#" +
	"\x0cH\x13Q\xd4\xd7\x9f\xbe1\x9f\xbe=\xfc\x0a\x13\x09" +
	"\xf65\xb9\x1fX\xe5\x1a\xeb\xb1\xb3#\x9d\xe5\xa9t\x13" +
	"\x1b2\x97\xeb\xd4\x15zb\xbd\xcf\xf2[z\xd0\xd7\xb6" +
	"\x97\x8f\xf4\xd5~\x83\x97\x11\x10\x11PG/\x00fY" +
	"\xd2\x9c\x14Td\x97M\\=\x0b\x98\x13\x92\xe6\x94`" +
	"\xedoO\xcblh\xc7\x00\x18C0\x06\xef\xdd\xc9\x8a" +
	"P\xd9\xf1\xfc\xbe0sf>7\xb8,\xab\xb07\xf7" +
	"\xf0\x8c\xab\xf8\xa3vU\xf0U\xd0N\xa4z\xe8\xf2`" +
	"\xb3<)\xb4\xb7E)\x93\x020\x9d\x05\xd1\x1e\x03\xcc" +
	"5Is\xf7?b\xb5\x0e\x18/i\x1e\x09*!\xba" +
	"\x14\x80z\xb0\x06\x98\xfb\x92\xe6\xad\xa0\x92\xb2K\x09\xa8" +
	"\xd7M|.i\xde\x0b\x1e\xcc\xed$\xe1\x01p \x1b" +
	"\x87h\x8e\xa7]\x9a\x96I\x98\xd7\xf8o\xed\xdd\x9c\x86" +
	"\xa4\x9c\xc7\xa5\x7f\xb1y\xe1E\xec\xec\xfe?\xd8\xc4?" +
	"\x01\x00\x00\xff\xff\xcb\x91\x7f\xec"

func init() {
	schemas.Register(schema_9c2ec473614d6dbe,
		0xa85299c27791f521,
		0xb60d2ba29728e68f)
}