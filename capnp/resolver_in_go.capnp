@0xeb0269f5b51f88bd;

using Go = import "/go.capnp";
$Go.package("pangine");
$Go.import("github.com/pangine");

struct ResolverIn $Go.doc("output of cpp llvm resolver") {
  terminate @0 :Bool;
  bytes @1 :List(UInt8);
}
