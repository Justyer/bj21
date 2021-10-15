const path = require('path')
const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const packageDefinition = protoLoader.loadSync(
    path.join(__dirname, '../api/bj21/v1/bj21.proto'),
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    },
)
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
console.log(protoDescriptor)
const bjproto = protoDescriptor.bj21.v1

module.exports = bjproto