@0xdebeff1c5643386f;

using Go = import "/go.capnp";
$Go.package("pangine");
$Go.import("github.com/pangine");

struct ResolverOut $Go.doc("input of cpp llvm resolver") {
  isInst @0: Bool;
  takeBytes @1: Int32;
  inst @2: Text;
}
