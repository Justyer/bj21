// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var bj21_pb = require('./bj21_pb.js');

function serialize_bj21_v1_Message(arg) {
  if (!(arg instanceof bj21_pb.Message)) {
    throw new Error('Expected argument of type bj21.v1.Message');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_bj21_v1_Message(buffer_arg) {
  return bj21_pb.Message.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlackJackService = exports.BlackJackService = {
  logicConn: {
    path: '/bj21.v1.BlackJack/LogicConn',
    requestStream: true,
    responseStream: true,
    requestType: bj21_pb.Message,
    responseType: bj21_pb.Message,
    requestSerialize: serialize_bj21_v1_Message,
    requestDeserialize: deserialize_bj21_v1_Message,
    responseSerialize: serialize_bj21_v1_Message,
    responseDeserialize: deserialize_bj21_v1_Message,
  },
};

exports.BlackJackClient = grpc.makeGenericClientConstructor(BlackJackService);
