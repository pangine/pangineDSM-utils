@0x9c2ec473614d6dbe;

using Go = import "/go.capnp";
$Go.package("pangine");
$Go.import("github.com/pangine");

struct CPOut $Go.doc("output of container parser") {
  name @0   : List(Text);
  offset @1 : List(Int64);
  bytes @2  : List(UInt8);
  loads @3  : List(HeaderLoad);
}

struct HeaderLoad $Go.doc("load infomation for mapping PA and VA") {
  physical @0 : Int64;
  virtual @1  : Int64;
}
