@0x965502164927c095;

using Go = import "/go.capnp";
$Go.package("pangine");
$Go.import("github.com/pangine");

struct StageAPI $Go.doc("The message sent to a stage service (including rt)") {
  terminate @0 :Bool;
  id @1 :Int64;
  offsets @2 :List(Int64);
  rollbacked @3 :Bool;
  rollbackedToId @4 :Int64;
}
