/**
 * uisvc
 * API for the user interface service; the service that is directly responsible for presenting data to users. 
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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Uisvc);
  }
}(this, function(expect, Uisvc) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Uisvc.UserError();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('UserError', function() {
    it('should create an instance of UserError', function() {
      // uncomment below and update the code to test UserError
      //var instance = new Uisvc.UserError();
      //expect(instance).to.be.a(Uisvc.UserError);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new Uisvc.UserError();
      //expect(instance).to.be();
    });

    it('should have the property error (base name: "error")', function() {
      // uncomment below and update the code to test the property error
      //var instance = new Uisvc.UserError();
      //expect(instance).to.be();
    });

  });

}));
