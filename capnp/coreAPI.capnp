@0xfe982ce7e37aff22;

using Go = import "/go.capnp";
$Go.package("pangine");
$Go.import("github.com/pangine");

struct CoreAPI $Go.doc("The message sent from other services to core service") {
  id @0 :Int64;
  messageType @1 :CoreAPIType;
  offsets @2 :List(Int64);
}

struct CoreAPIReply $Go.doc("The message sent from core service replying a request or a rollback") {
  id @0 :Int64;
  accepted @1 :Bool;
  offsets @2 :List(Int64);
}

enum CoreAPIType $Go.doc("The API message type") {
  response @0;
  request @1;
  rollback @2;
}
