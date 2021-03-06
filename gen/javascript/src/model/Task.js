/**
 * uisvc
 * API for the user interface service; the service that is directly responsible for presenting data to users. This service typically runs at the border, and leverages session cookies or authentication tokens that we generate for users. It also is responsible for handling the act of oauth and user creation through its login hooks. uisvc typically talks to the datasvc and other services to accomplish its goal, it does not save anything locally or carry state. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.0-SNAPSHOT
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Ref', 'model/Repository', 'model/TaskSettings'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Ref'), require('./Repository'), require('./TaskSettings'));
  } else {
    // Browser globals (root is window)
    if (!root.Uisvc) {
      root.Uisvc = {};
    }
    root.Uisvc.Task = factory(root.Uisvc.ApiClient, root.Uisvc.Ref, root.Uisvc.Repository, root.Uisvc.TaskSettings);
  }
}(this, function(ApiClient, Ref, Repository, TaskSettings) {
  'use strict';




  /**
   * The Task model module.
   * @module model/Task
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Task</code>.
   * @alias module:model/Task
   * @class
   */
  var exports = function() {
    var _this = this;













  };

  /**
   * Constructs a <code>Task</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Task} obj Optional instance to populate.
   * @return {module:model/Task} The populated <code>Task</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('path')) {
        obj['path'] = ApiClient.convertToType(data['path'], 'String');
      }
      if (data.hasOwnProperty('settings')) {
        obj['settings'] = TaskSettings.constructFromObject(data['settings']);
      }
      if (data.hasOwnProperty('parent')) {
        obj['parent'] = Repository.constructFromObject(data['parent']);
      }
      if (data.hasOwnProperty('ref')) {
        obj['ref'] = Ref.constructFromObject(data['ref']);
      }
      if (data.hasOwnProperty('base_sha')) {
        obj['base_sha'] = ApiClient.convertToType(data['base_sha'], 'String');
      }
      if (data.hasOwnProperty('pull_request_id')) {
        obj['pull_request_id'] = ApiClient.convertToType(data['pull_request_id'], 'Number');
      }
      if (data.hasOwnProperty('canceled')) {
        obj['canceled'] = ApiClient.convertToType(data['canceled'], 'Boolean');
      }
      if (data.hasOwnProperty('finished_at')) {
        obj['finished_at'] = ApiClient.convertToType(data['finished_at'], 'Date');
      }
      if (data.hasOwnProperty('started_at')) {
        obj['started_at'] = ApiClient.convertToType(data['started_at'], 'Date');
      }
      if (data.hasOwnProperty('created_at')) {
        obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'Boolean');
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'Number');
      }
    }
    return obj;
  }

  /**
   * @member {String} path
   */
  exports.prototype['path'] = undefined;
  /**
   * @member {module:model/TaskSettings} settings
   */
  exports.prototype['settings'] = undefined;
  /**
   * @member {module:model/Repository} parent
   */
  exports.prototype['parent'] = undefined;
  /**
   * @member {module:model/Ref} ref
   */
  exports.prototype['ref'] = undefined;
  /**
   * @member {String} base_sha
   */
  exports.prototype['base_sha'] = undefined;
  /**
   * @member {Number} pull_request_id
   */
  exports.prototype['pull_request_id'] = undefined;
  /**
   * @member {Boolean} canceled
   */
  exports.prototype['canceled'] = undefined;
  /**
   * @member {Date} finished_at
   */
  exports.prototype['finished_at'] = undefined;
  /**
   * @member {Date} started_at
   */
  exports.prototype['started_at'] = undefined;
  /**
   * @member {Date} created_at
   */
  exports.prototype['created_at'] = undefined;
  /**
   * @member {Boolean} status
   */
  exports.prototype['status'] = undefined;
  /**
   * @member {Number} id
   */
  exports.prototype['id'] = undefined;



  return exports;
}));


