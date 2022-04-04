# gRPC (Google Remote Procedure Call)

gRPC (gRPC Remote Procedure Calls) is an open source remote procedure call (RPC) system initially developed at Google in 2015 as the next generation of the RPC infrastructure. It uses HTTP/2 for transport, Protocol Buffers as the interface description language, and provides features such as authentication, bidirectional streaming and flow control, blocking or nonblocking bindings, and cancellation and timeouts. It generates cross-platform client and server bindings for many languages. Most common usage scenarios include connecting services in a microservices style architecture, or connecting mobile device clients to backend services. 

## Working


- A client application can call methods directly on a server-side application present on other machines.
- Service is defined, methods are specified, which can be further remotely called with their parameters and return types.
- On the other hand, the server runs a gRPC server to handle client calls.
- It uses protocol buffers as the Interface Definition Language to enable communication between two systems used to describe the service interface and the structure of payload messages.
- HTTP/2 - gRPC is basically a protocol built on top of HTTP/2. HTTP/2 is used as transport.
- Protobuf serialization - Messages that we serialize both for the request and response are encoded with protocol buffers.
- Clients open one long-lived connection to the gRPC server.
- A new HTTP/2 stream for each RPC call.
- Allows Client-Side and Server-Side Streaming.
- Bidirectional Streaming.

## Benefits


- Easy to understand.
- Web infrastructure already built on top of HTTP.
- Great tooling for testing, inspection, and modification.
- Loose coupling between clients/servers makes changes easy.
- High-quality HTTP implementations in every language.



