/* eslint-disable */
/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./dare_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.TaskServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.TaskServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.GetAllRequest,
 *   !proto.GetAllResponse>}
 */
const methodInfo_TaskService_GetAll = new grpc.web.AbstractClientBase.MethodInfo(
  proto.GetAllResponse,
  /** @param {!proto.GetAllRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.GetAllResponse.deserializeBinary
);


/**
 * @param {!proto.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.GetAllResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.GetAllResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.TaskServiceClient.prototype.getAll =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/TaskService/GetAll',
      request,
      metadata || {},
      methodInfo_TaskService_GetAll,
      callback);
};


/**
 * @param {!proto.GetAllRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.GetAllResponse>}
 *     A native promise that resolves to the response
 */
proto.TaskServicePromiseClient.prototype.getAll =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/TaskService/GetAll',
      request,
      metadata || {},
      methodInfo_TaskService_GetAll);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.UpsertRequest,
 *   !proto.UpsertResponse>}
 */
const methodInfo_TaskService_Upsert = new grpc.web.AbstractClientBase.MethodInfo(
  proto.UpsertResponse,
  /** @param {!proto.UpsertRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.UpsertResponse.deserializeBinary
);


/**
 * @param {!proto.UpsertRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.UpsertResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.UpsertResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.TaskServiceClient.prototype.upsert =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/TaskService/Upsert',
      request,
      metadata || {},
      methodInfo_TaskService_Upsert,
      callback);
};


/**
 * @param {!proto.UpsertRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.UpsertResponse>}
 *     A native promise that resolves to the response
 */
proto.TaskServicePromiseClient.prototype.upsert =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/TaskService/Upsert',
      request,
      metadata || {},
      methodInfo_TaskService_Upsert);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.DeleteRequest,
 *   !proto.DeleteResponse>}
 */
const methodInfo_TaskService_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.DeleteResponse,
  /** @param {!proto.DeleteRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.TaskServiceClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/TaskService/Delete',
      request,
      metadata || {},
      methodInfo_TaskService_Delete,
      callback);
};


/**
 * @param {!proto.DeleteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.DeleteResponse>}
 *     A native promise that resolves to the response
 */
proto.TaskServicePromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/TaskService/Delete',
      request,
      metadata || {},
      methodInfo_TaskService_Delete);
};


module.exports = proto;

