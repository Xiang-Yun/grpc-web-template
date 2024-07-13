/**
 * @fileoverview gRPC-Web generated client stub for grpc_api
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.21.12
// source: api_proto/web.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.grpc_api = require('./web_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.grpc_api.APIServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.grpc_api.APIServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.grpc_api.ReadItemRequest,
 *   !proto.grpc_api.MessageResponse>}
 */
const methodDescriptor_APIService_ReadItem = new grpc.web.MethodDescriptor(
  '/grpc_api.APIService/ReadItem',
  grpc.web.MethodType.UNARY,
  proto.grpc_api.ReadItemRequest,
  proto.grpc_api.MessageResponse,
  /**
   * @param {!proto.grpc_api.ReadItemRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpc_api.MessageResponse.deserializeBinary
);


/**
 * @param {!proto.grpc_api.ReadItemRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.grpc_api.MessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.grpc_api.MessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.grpc_api.APIServiceClient.prototype.readItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/grpc_api.APIService/ReadItem',
      request,
      metadata || {},
      methodDescriptor_APIService_ReadItem,
      callback);
};


/**
 * @param {!proto.grpc_api.ReadItemRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.grpc_api.MessageResponse>}
 *     Promise that resolves to the response
 */
proto.grpc_api.APIServicePromiseClient.prototype.readItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/grpc_api.APIService/ReadItem',
      request,
      metadata || {},
      methodDescriptor_APIService_ReadItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.grpc_api.MessageRequest,
 *   !proto.grpc_api.MessageResponse>}
 */
const methodDescriptor_APIService_SendMessage = new grpc.web.MethodDescriptor(
  '/grpc_api.APIService/SendMessage',
  grpc.web.MethodType.UNARY,
  proto.grpc_api.MessageRequest,
  proto.grpc_api.MessageResponse,
  /**
   * @param {!proto.grpc_api.MessageRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpc_api.MessageResponse.deserializeBinary
);


/**
 * @param {!proto.grpc_api.MessageRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.grpc_api.MessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.grpc_api.MessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.grpc_api.APIServiceClient.prototype.sendMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/grpc_api.APIService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_APIService_SendMessage,
      callback);
};


/**
 * @param {!proto.grpc_api.MessageRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.grpc_api.MessageResponse>}
 *     Promise that resolves to the response
 */
proto.grpc_api.APIServicePromiseClient.prototype.sendMessage =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/grpc_api.APIService/SendMessage',
      request,
      metadata || {},
      methodDescriptor_APIService_SendMessage);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.grpc_api.DeleteItemRequest,
 *   !proto.grpc_api.MessageResponse>}
 */
const methodDescriptor_APIService_DeleteItem = new grpc.web.MethodDescriptor(
  '/grpc_api.APIService/DeleteItem',
  grpc.web.MethodType.UNARY,
  proto.grpc_api.DeleteItemRequest,
  proto.grpc_api.MessageResponse,
  /**
   * @param {!proto.grpc_api.DeleteItemRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpc_api.MessageResponse.deserializeBinary
);


/**
 * @param {!proto.grpc_api.DeleteItemRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.grpc_api.MessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.grpc_api.MessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.grpc_api.APIServiceClient.prototype.deleteItem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/grpc_api.APIService/DeleteItem',
      request,
      metadata || {},
      methodDescriptor_APIService_DeleteItem,
      callback);
};


/**
 * @param {!proto.grpc_api.DeleteItemRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.grpc_api.MessageResponse>}
 *     Promise that resolves to the response
 */
proto.grpc_api.APIServicePromiseClient.prototype.deleteItem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/grpc_api.APIService/DeleteItem',
      request,
      metadata || {},
      methodDescriptor_APIService_DeleteItem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.grpc_api.MessageResponse>}
 */
const methodDescriptor_APIService_ShowClock = new grpc.web.MethodDescriptor(
  '/grpc_api.APIService/ShowClock',
  grpc.web.MethodType.SERVER_STREAMING,
  google_protobuf_empty_pb.Empty,
  proto.grpc_api.MessageResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.grpc_api.MessageResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.grpc_api.MessageResponse>}
 *     The XHR Node Readable Stream
 */
proto.grpc_api.APIServiceClient.prototype.showClock =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/grpc_api.APIService/ShowClock',
      request,
      metadata || {},
      methodDescriptor_APIService_ShowClock);
};


/**
 * @param {!proto.google.protobuf.Empty} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.grpc_api.MessageResponse>}
 *     The XHR Node Readable Stream
 */
proto.grpc_api.APIServicePromiseClient.prototype.showClock =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/grpc_api.APIService/ShowClock',
      request,
      metadata || {},
      methodDescriptor_APIService_ShowClock);
};


module.exports = proto.grpc_api;

