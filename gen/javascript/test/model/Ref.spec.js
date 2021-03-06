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
    instance = new Uisvc.Ref();
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

  describe('Ref', function() {
    it('should create an instance of Ref', function() {
      // uncomment below and update the code to test Ref
      //var instance = new Uisvc.Ref();
      //expect(instance).to.be.a(Uisvc.Ref);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new Uisvc.Ref();
      //expect(instance).to.be();
    });

    it('should have the property repository (base name: "repository")', function() {
      // uncomment below and update the code to test the property repository
      //var instance = new Uisvc.Ref();
      //expect(instance).to.be();
    });

    it('should have the property refName (base name: "ref_name")', function() {
      // uncomment below and update the code to test the property refName
      //var instance = new Uisvc.Ref();
      //expect(instance).to.be();
    });

    it('should have the property sha (base name: "sha")', function() {
      // uncomment below and update the code to test the property sha
      //var instance = new Uisvc.Ref();
      //expect(instance).to.be();
    });

  });

}));
