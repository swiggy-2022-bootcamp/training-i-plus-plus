protoc services/auth/authpb/auth.proto --go-grpc_out=. --go_out=.

#Server Start
go run services/auth/server/server.go

#Client Start
go run greet/auth/client/client.go

#MongoDB Start
mongod --dbpath="C:\Program Files\MongoDB\Server\4.2\data".