# uisvc

Uisvc - JavaScript client for uisvc
API for the user interface service; the service that is directly responsible for presenting data to users. This service typically runs at the border, and leverages session cookies or authentication tokens that we generate for users. It also is responsible for handling the act of oauth and user creation through its login hooks. uisvc typically talks to the datasvc and other services to accomplish its goal, it does not save anything locally or carry state. 
This SDK is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.JavascriptClientCodegen

## Installation

### For [Node.js](https://nodejs.org/)

#### npm

To publish the library as a [npm](https://www.npmjs.com/),
please follow the procedure in ["Publishing npm packages"](https://docs.npmjs.com/getting-started/publishing-npm-packages).

Then install it via:

```shell
npm install uisvc --save
```

##### Local development

To use the library locally without publishing to a remote npm registry, first install the dependencies by changing 
into the directory containing `package.json` (and this README). Let's call this `JAVASCRIPT_CLIENT_DIR`. Then run:

```shell
npm install
```

Next, [link](https://docs.npmjs.com/cli/link) it globally in npm with the following, also from `JAVASCRIPT_CLIENT_DIR`:

```shell
npm link
```

Finally, switch to the directory you want to use your uisvc from, and run:

```shell
npm link /path/to/<JAVASCRIPT_CLIENT_DIR>
```

You should now be able to `require('uisvc')` in javascript files from the directory you ran the last 
command above from.

#### git
#
If the library is hosted at a git repository, e.g.
https://github.com/GIT_USER_ID/GIT_REPO_ID
then install it via:

```shell
    npm install GIT_USER_ID/GIT_REPO_ID --save
```

### For browser

The library also works in the browser environment via npm and [browserify](http://browserify.org/). After following
the above steps with Node.js and installing browserify with `npm install -g browserify`,
perform the following (assuming *main.js* is your entry file, that's to say your javascript file where you actually 
use this library):

```shell
browserify main.js > bundle.js
```

Then include *bundle.js* in the HTML pages.

### Webpack Configuration

Using Webpack you may encounter the following error: "Module not found: Error:
Cannot resolve module", most certainly you should disable AMD loader. Add/merge
the following section to your webpack config:

```javascript
module: {
  rules: [
    {
      parser: {
        amd: false
      }
    }
  ]
}
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following JS code:

```javascript
var Uisvc = require('uisvc');

var defaultClient = Uisvc.ApiClient.instance;

// Configure API key authorization: session
var session = defaultClient.authentications['session'];
session.apiKey = "YOUR API KEY"
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//session.apiKeyPrefix['Cookie'] = "Token"

// Configure API key authorization: token
var token = defaultClient.authentications['token'];
token.apiKey = "YOUR API KEY"
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//token.apiKeyPrefix['Authorization'] = "Token"

var api = new Uisvc.DefaultApi()

var runId = 56; // {Number} The ID of the run to retrieve


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
api.cancelRunIdPost(runId, callback);

```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Uisvc.DefaultApi* | [**cancelRunIdPost**](docs/DefaultApi.md#cancelRunIdPost) | **POST** /cancel/{run_id} | Cancel by Run ID
*Uisvc.DefaultApi* | [**errorsGet**](docs/DefaultApi.md#errorsGet) | **GET** /errors | Retrieve errors
*Uisvc.DefaultApi* | [**logAttachIdGet**](docs/DefaultApi.md#logAttachIdGet) | **GET** /log/attach/{id} | Attach to a running log
*Uisvc.DefaultApi* | [**loggedinGet**](docs/DefaultApi.md#loggedinGet) | **GET** /loggedin | Check logged in state
*Uisvc.DefaultApi* | [**loginGet**](docs/DefaultApi.md#loginGet) | **GET** /login | Log into the system
*Uisvc.DefaultApi* | [**logoutGet**](docs/DefaultApi.md#logoutGet) | **GET** /logout | Log out of the system
*Uisvc.DefaultApi* | [**repositoriesCiAddOwnerRepoGet**](docs/DefaultApi.md#repositoriesCiAddOwnerRepoGet) | **GET** /repositories/ci/add/{owner}/{repo} | Add a specific repository to CI.
*Uisvc.DefaultApi* | [**repositoriesCiDelOwnerRepoGet**](docs/DefaultApi.md#repositoriesCiDelOwnerRepoGet) | **GET** /repositories/ci/del/{owner}/{repo} | Removes a specific repository from CI.
*Uisvc.DefaultApi* | [**repositoriesMyGet**](docs/DefaultApi.md#repositoriesMyGet) | **GET** /repositories/my | Fetch all the writable repositories for the user.
*Uisvc.DefaultApi* | [**repositoriesSubAddOwnerRepoGet**](docs/DefaultApi.md#repositoriesSubAddOwnerRepoGet) | **GET** /repositories/sub/add/{owner}/{repo} | Subscribe to a repository running CI
*Uisvc.DefaultApi* | [**repositoriesSubDelOwnerRepoGet**](docs/DefaultApi.md#repositoriesSubDelOwnerRepoGet) | **GET** /repositories/sub/del/{owner}/{repo} | Unsubscribe from a repository
*Uisvc.DefaultApi* | [**repositoriesSubscribedGet**](docs/DefaultApi.md#repositoriesSubscribedGet) | **GET** /repositories/subscribed | List all subscribed repositories
*Uisvc.DefaultApi* | [**repositoriesVisibleGet**](docs/DefaultApi.md#repositoriesVisibleGet) | **GET** /repositories/visible | Fetch all the repositories the user can view.
*Uisvc.DefaultApi* | [**runRunIdGet**](docs/DefaultApi.md#runRunIdGet) | **GET** /run/{run_id} | Get a run by ID
*Uisvc.DefaultApi* | [**runsCountGet**](docs/DefaultApi.md#runsCountGet) | **GET** /runs/count | Count the runs
*Uisvc.DefaultApi* | [**runsGet**](docs/DefaultApi.md#runsGet) | **GET** /runs | Obtain the run list for the user
*Uisvc.DefaultApi* | [**submitGet**](docs/DefaultApi.md#submitGet) | **GET** /submit | Perform a manual submission to tinyCI
*Uisvc.DefaultApi* | [**tasksCountGet**](docs/DefaultApi.md#tasksCountGet) | **GET** /tasks/count | Count the Tasks
*Uisvc.DefaultApi* | [**tasksGet**](docs/DefaultApi.md#tasksGet) | **GET** /tasks | Obtain the task list optionally filtering by repository and sha.
*Uisvc.DefaultApi* | [**tasksRunsIdCountGet**](docs/DefaultApi.md#tasksRunsIdCountGet) | **GET** /tasks/runs/{id}/count | Count the runs corresponding to the task ID.
*Uisvc.DefaultApi* | [**tasksRunsIdGet**](docs/DefaultApi.md#tasksRunsIdGet) | **GET** /tasks/runs/{id} | Obtain the run list based on the task ID.
*Uisvc.DefaultApi* | [**tasksSubscribedGet**](docs/DefaultApi.md#tasksSubscribedGet) | **GET** /tasks/subscribed | Obtain the list of tasks that belong to repositories you are subscribed to.
*Uisvc.DefaultApi* | [**tokenDelete**](docs/DefaultApi.md#tokenDelete) | **DELETE** /token | Remove and reset your tinyCI access token
*Uisvc.DefaultApi* | [**tokenGet**](docs/DefaultApi.md#tokenGet) | **GET** /token | Get a tinyCI access token
*Uisvc.DefaultApi* | [**userPropertiesGet**](docs/DefaultApi.md#userPropertiesGet) | **GET** /user/properties | Get information about the current user


## Documentation for Models

 - [Uisvc.Error](docs/Error.md)
 - [Uisvc.Ref](docs/Ref.md)
 - [Uisvc.RepoConfig](docs/RepoConfig.md)
 - [Uisvc.Repository](docs/Repository.md)
 - [Uisvc.RepositoryList](docs/RepositoryList.md)
 - [Uisvc.Run](docs/Run.md)
 - [Uisvc.RunList](docs/RunList.md)
 - [Uisvc.RunSettings](docs/RunSettings.md)
 - [Uisvc.Task](docs/Task.md)
 - [Uisvc.TaskList](docs/TaskList.md)
 - [Uisvc.TaskSettings](docs/TaskSettings.md)
 - [Uisvc.UserError](docs/UserError.md)


## Documentation for Authorization


### session

- **Type**: API key
- **API key parameter name**: Cookie
- **Location**: HTTP header

### token

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header
