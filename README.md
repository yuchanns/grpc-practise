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
### Install php plugins
```sh
pecl install grpc-1.34.0
## find the location of php.ini and enable grpc plugin
php -i | grep php.ini
echo 'extension=grpc.so' >> php.ini
## install composer dependencies
composer install
## install protobuf runtime library
pecl install protobuf
echo 'extension=protobuf.so' >> php.ini
## check plugin status
php -m | egrep 'grpc|protobuf'
```
Before build `grpc_php_plugin`, install `bazel`.
```sh
sudo apt install curl gnupg
curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg
sudo mv bazel.gpg /etc/apt/trusted.gpg.d/
echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
apt update
apt install bazel
```
Then build `grpc_php_plugin`
```sh
git submodule update --init --recursive 
cd grpc
## build grpc_php_plugin with bazel
bazel build @com_google_protobuf//:protoc
bazel build src/compiler:grpc_php_plugin
```
### Install node plugins
```sh
cd node-client
yarn add grpc @grpc/proto-loader
```

## Generate Code
### Go server
From the repository root:
```sh
protoc -I ./proto --go_out=plugins=grpc:./go-server/proto/greeter ./proto/*.proto
```
Then a file `greeter.pb.go` shall show under the path `go-server/proto/greeter`.

You can start the server in one console and test the client request on another.
```sh
# console 1
go run .
# console 2
cd go-server && go test -v .
# === RUN   TestSayHello
#     pkg_test.go:27: get response from SayHello: hello, yuchanns
# --- PASS: TestSayHello (0.00s)
# PASS
# ok      github.com/yuchanns/grpc-practise       0.007s
```
### PHP client
From the repository root:
```sh
protoc -I ./proto --php_out=./php-client/proto --grpc_out=./php-client/proto --plugin=protoc-gen-grpc=./grpc/bazel-bin/src/compiler/grpc_php_plugin ./proto/*.proto 
```
Then two folders `GPBMetadata` and `Greeter` shall show under the path `php-client/proto`

You can run the php client to request go grpc server and get response.
```sh
php php-client/main.php 
# hello, php
```
### Node client
Simply load the protofile
```js
// node-client/index.js

const PROTO_PATH = __dirname + '/../proto/greeter.proto'
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const greeter = grpc.loadPackageDefinition(packageDefinition).greeter
const client = new greeter.Greeter("localhost:9090", grpc.credentials.createInsecure())
client.SayHello({name: "node"}, (error, resp) => {
    if (error) {
        console.log(error)
        return
    }
    console.log(resp)
})
```
Then run in console.
```sh
node index.js
# { msg: 'hello, node' }
```
