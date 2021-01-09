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