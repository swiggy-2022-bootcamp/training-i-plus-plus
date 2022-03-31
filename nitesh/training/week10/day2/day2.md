## gRPC implementation example

### Prerequisites

- Go
- Protocol buffer
- Go plugins for the protocol compiler:

  Install the protocol compiler plugins for Go using the following commands:

  ```sh
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
  ```

  Update your PATH so that the protoc compiler can find the plugins:

  ```sh
  $ export PATH="$PATH:$(go env GOPATH)/bin"
  ```

### Example Code

clone the repo:

```sh
$ git clone -b v1.45.0 --depth 1 https://github.com/grpc/grpc-go
```

Change to the quick start example directory:

```sh
$ cd grpc-go/examples/helloworld
```

### Run the example

From the examples/helloworld directory:

Compile and execute the server code:

```sh
$ go run greeter_server/main.go
```

From a different terminal, compile and execute the client code to see the client output:

```sh
$ go run greeter_client/main.go
Greeting: Hello world
```

Read the example code the play with it.

## Must read resources

- [gRPC.io](https://grpc.io/docs/languages/go/basics/)
- [gRPC guide](https://www.tutorialspoint.com/grpc/grpc_quick_guide.htm)
