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
    define(['ApiClient', 'model/RunSettings', 'model/Task'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./RunSettings'), require('./Task'));
  } else {
    // Browser globals (root is window)
    if (!root.Uisvc) {
      root.Uisvc = {};
    }
    root.Uisvc.Run = factory(root.Uisvc.ApiClient, root.Uisvc.RunSettings, root.Uisvc.Task);
  }
}(this, function(ApiClient, RunSettings, Task) {
  'use strict';




  /**
   * The Run model module.
   * @module model/Run
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Run</code>.
   * @alias module:model/Run
   * @class
   */
  var exports = function() {
    var _this = this;









  };

  /**
   * Constructs a <code>Run</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Run} obj Optional instance to populate.
   * @return {module:model/Run} The populated <code>Run</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('settings')) {
        obj['settings'] = RunSettings.constructFromObject(data['settings']);
      }
      if (data.hasOwnProperty('name')) {
        obj['name'] = ApiClient.convertToType(data['name'], 'String');
      }
      if (data.hasOwnProperty('created_at')) {
        obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
      }
      if (data.hasOwnProperty('started_at')) {
        obj['started_at'] = ApiClient.convertToType(data['started_at'], 'Date');
      }
      if (data.hasOwnProperty('finished_at')) {
        obj['finished_at'] = ApiClient.convertToType(data['finished_at'], 'Date');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'Boolean');
      }
      if (data.hasOwnProperty('task')) {
        obj['task'] = Task.constructFromObject(data['task']);
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'Number');
      }
    }
    return obj;
  }

  /**
   * @member {module:model/RunSettings} settings
   */
  exports.prototype['settings'] = undefined;
  /**
   * @member {String} name
   */
  exports.prototype['name'] = undefined;
  /**
   * @member {Date} created_at
   */
  exports.prototype['created_at'] = undefined;
  /**
   * @member {Date} started_at
   */
  exports.prototype['started_at'] = undefined;
  /**
   * @member {Date} finished_at
   */
  exports.prototype['finished_at'] = undefined;
  /**
   * @member {Boolean} status
   */
  exports.prototype['status'] = undefined;
  /**
   * @member {module:model/Task} task
   */
  exports.prototype['task'] = undefined;
  /**
   * @member {Number} id
   */
  exports.prototype['id'] = undefined;



  return exports;
}));

