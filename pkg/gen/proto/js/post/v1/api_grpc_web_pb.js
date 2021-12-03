/**
 * @fileoverview gRPC-Web generated client stub for post.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var protoc$gen$validate_validate_validate_pb = require('../../protoc-gen-validate/validate/validate_pb.js')

var post_v1_post_pb = require('../../post/v1/post_pb.js')
const proto = {};
proto.post = {};
proto.post.v1 = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.post.v1.PostServiceClient =
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
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.post.v1.PostServicePromiseClient =
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
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.post.v1.CreateRequest,
 *   !proto.post.v1.CreateResponse>}
 */
const methodDescriptor_PostService_Create = new grpc.web.MethodDescriptor(
  '/post.v1.PostService/Create',
  grpc.web.MethodType.UNARY,
  proto.post.v1.CreateRequest,
  proto.post.v1.CreateResponse,
  /**
   * @param {!proto.post.v1.CreateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.post.v1.CreateResponse.deserializeBinary
);


/**
 * @param {!proto.post.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.post.v1.CreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.post.v1.CreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.post.v1.PostServiceClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/post.v1.PostService/Create',
      request,
      metadata || {},
      methodDescriptor_PostService_Create,
      callback);
};


/**
 * @param {!proto.post.v1.CreateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.post.v1.CreateResponse>}
 *     Promise that resolves to the response
 */
proto.post.v1.PostServicePromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/post.v1.PostService/Create',
      request,
      metadata || {},
      methodDescriptor_PostService_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.post.v1.ReadRequest,
 *   !proto.post.v1.ReadResponse>}
 */
const methodDescriptor_PostService_Read = new grpc.web.MethodDescriptor(
  '/post.v1.PostService/Read',
  grpc.web.MethodType.UNARY,
  proto.post.v1.ReadRequest,
  proto.post.v1.ReadResponse,
  /**
   * @param {!proto.post.v1.ReadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.post.v1.ReadResponse.deserializeBinary
);


/**
 * @param {!proto.post.v1.ReadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.post.v1.ReadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.post.v1.ReadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.post.v1.PostServiceClient.prototype.read =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/post.v1.PostService/Read',
      request,
      metadata || {},
      methodDescriptor_PostService_Read,
      callback);
};


/**
 * @param {!proto.post.v1.ReadRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.post.v1.ReadResponse>}
 *     Promise that resolves to the response
 */
proto.post.v1.PostServicePromiseClient.prototype.read =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/post.v1.PostService/Read',
      request,
      metadata || {},
      methodDescriptor_PostService_Read);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.post.v1.UpdateRequest,
 *   !proto.post.v1.UpdateResponse>}
 */
const methodDescriptor_PostService_Update = new grpc.web.MethodDescriptor(
  '/post.v1.PostService/Update',
  grpc.web.MethodType.UNARY,
  proto.post.v1.UpdateRequest,
  proto.post.v1.UpdateResponse,
  /**
   * @param {!proto.post.v1.UpdateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.post.v1.UpdateResponse.deserializeBinary
);


/**
 * @param {!proto.post.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.post.v1.UpdateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.post.v1.UpdateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.post.v1.PostServiceClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/post.v1.PostService/Update',
      request,
      metadata || {},
      methodDescriptor_PostService_Update,
      callback);
};


/**
 * @param {!proto.post.v1.UpdateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.post.v1.UpdateResponse>}
 *     Promise that resolves to the response
 */
proto.post.v1.PostServicePromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/post.v1.PostService/Update',
      request,
      metadata || {},
      methodDescriptor_PostService_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.post.v1.DeleteRequest,
 *   !proto.post.v1.DeleteResponse>}
 */
const methodDescriptor_PostService_Delete = new grpc.web.MethodDescriptor(
  '/post.v1.PostService/Delete',
  grpc.web.MethodType.UNARY,
  proto.post.v1.DeleteRequest,
  proto.post.v1.DeleteResponse,
  /**
   * @param {!proto.post.v1.DeleteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.post.v1.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.post.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.post.v1.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.post.v1.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.post.v1.PostServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/post.v1.PostService/Delete',
      request,
      metadata || {},
      methodDescriptor_PostService_Delete,
      callback);
};


/**
 * @param {!proto.post.v1.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.post.v1.DeleteResponse>}
 *     Promise that resolves to the response
 */
proto.post.v1.PostServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/post.v1.PostService/Delete',
      request,
      metadata || {},
      methodDescriptor_PostService_Delete);
};


module.exports = proto.post.v1;

