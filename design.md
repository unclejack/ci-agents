- uisvc
  - returns data from the data service
  - not gated by client certs, but oauth-backed session cookies
  - stuff like repository list, queue list, etc.
  - _LIVES AT BORDER_
- queue svc
  - returns next item in queue
  - accepts submissions from the hook service; injects into queue
- hook svc
  - accepts submissions from github, talks to queue service
  - _LIVES AT BORDER_
- asset svc (called logsvc atm)
  - returns logs and other uploaded assets
  - accepts log submissions
  - operates over a basic tcp protocol; does not use http
- log service
  - accepts log entries from other services
  - returns queries of log data
- metrics service
  - accepts metrics for repositories and other entities
  - outputs prometheus stats for repositories
  - exposes an endpoint for global stats too
- scm api service
  - abstract component; intended to talk to github, bitbucket et al
  - performs actions on behalf of the user, gets data for the user, etc.
  - functionality to manage oauth here (utilized by auth service)
- data service
  - talks to the db
  - records and returns data from the db based on the model
