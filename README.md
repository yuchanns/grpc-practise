# grpc-practise
grpc-practise: using go as server and php | node as client

## Prepare
Clone this repository or open it on [Github Codespaces](https://github.com/codespaces).

## Install environment
### Install protoc
```sh
# download protobuf compiler and make it
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protobuf-all-3.14.0.tar.gz
tar -zxvf protobuf-all-3.14.0.tar.gz
cd protobuf-3.14.0 
./configure
make
make install
# load the dynamic library
ldconfig
# check version
protoc --version
```
### Install go plugins
```sh
GO111MODULE=on GOPROXY=https://goproxy.cn go get github.com/golang/protobuf/protoc-gen-go@v1.3
GO111MODULE=on GOPROXY=https://goproxy.cn go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1
```

## Generate Code
### Go server
From the root path:
```sh
protoc -I ./proto --go_out=plugins=grpc:./go-server/proto/greeter ./proto/*.proto
```
Then a file `greeter.pb.go` shall show under the path `go-server/proto/greeter`.

You can start the server in one console and test the client request on another.
```sh
# console 1
go run .
# console 2
go test -v .
```
