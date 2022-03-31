## gRPC introduction

![grpc-overview](https://blog.jdriven.com/uploads/2018/10/grpc-overview.png)

gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

#### The main usage scenarios

Efficiently connecting polyglot services in microservices style architecture
Connecting mobile devices, browser clients to backend services
Generating efficient client libraries

#### Core features that make it awesome

Idiomatic client libraries in 11 languages
Highly efficient on wire and with a simple service definition framework
Bi-directional streaming with http/2 based transport
Pluggable auth, tracing, load balancing and health checking

### RPCs (Remote Procedure Calls)

A remote procedure call is an interprocess communication technique that is used for client-server based applications. It is also known as a subroutine call or a function call.

A client has a request message that the RPC translates and sends to the server. This request may be a procedure or a function call to a remote server. When the server receives the request, it sends the required response back to the client. The client is blocked while the server is processing the call and only resumed execution after the server is finished.

The sequence of events in a remote procedure call are given as follows −

The client stub is called by the client.
The client stub makes a system call to send the message to the server and puts the parameters in the message.
The message is sent from the client to the server by the client’s operating system.
The message is passed to the server stub by the server operating system.
The parameters are removed from the message by the server stub.
Then, the server procedure is called by the server stub.

A diagram that demonstrates this is as follows −
![rpc](https://www.tutorialspoint.com/assets/questions/media/12686/RPC.PNG)

#### Advantages of Remote Procedure Call

Some of the advantages of RPC are as follows −

- Remote procedure calls support process oriented and thread oriented models.
- The internal message passing mechanism of RPC is hidden from the user.
- The effort to re-write and re-develop the code is minimum in remote procedure calls.
- Remote procedure calls can be used in distributed environment as well as the local environment.
- Many of the protocol layers are omitted by RPC to improve performance.

#### Disadvantages of Remote Procedure Call

Some of the disadvantages of RPC are as follows −

- The remote procedure call is a concept that can be implemented in different ways. It is not a standard.
- There is no flexibility in RPC for hardware architecture. It is only interaction based.
- There is an increase in costs because of remote procedure call.

### What are Remote Procedure Calls?

Remote procedure calls are the function calls which look like general/local function calls but differ in the fact that the execution of remote functions calls typically take place on a different machine. However, for the developer writing the code, there is minimal difference between a function call and a remote call. The calls typically follow the client-server model, where the machine which executes the call acts as the server.

### Why do we need remote procedure calls?

Remote procedure calls provide a way to execute codes on another machine. It becomes of utmost importance in big, bulky products where a single machine cannot host all the code which is required for the overall product to function.

In microservice architecture, the application is broken down into small services and these services communicate with each other via messaging queue and APIs. And all of this communication takes place over a network where different machines/nodes serve different functionality based on the service they host. So, creating remote procedure calls becomes a critical aspect when it comes to working in a distributed environment.

### Why gRPC?

Google Remote Procedure Calls (gRPC) provides a framework to perform the remote procedure calls. But there are some other libraries and mechanisms to execute code on remote machine. So, what makes gRPC special? Let's find out.

- Language independent − gRPC uses Google Protocol Buffer internally. So, multiple languages can be used such as Java, Python, Go, Dart, etc. A Java client can make a procedure call and a server that uses Python can respond, effectively, achieving language independence.

- Efficient Data Compaction − In microservice environment, given that multiple communications take place over a network, it is critical that the data that we are sending is as succinct as possible. We need to avoid any superfluous data to ensure that the data is quickly transferred. Given that gRPC uses Google Protocol Buffer internally, it has this feature to its advantage.

- Efficient serialization and deserialization − In microservice environment, given that multiple communications take place over a network, it is critical that we serialize and deserialize the data as quickly as possible. Given that gRPC uses Google Protocol Buffer internally, it ensures quick serializing and deserializing of data.

- Simple to use − gRPC already has a library and plugins that auto-generate procedure code (as we will see in the upcoming chapters). For simple use-cases, it can be used as local function calls.
