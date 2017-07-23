# Go Template

## Summary

__go_template__ is a fully developed starter package for building go applications. It includes a nginx server that sits in front of the main application server, a worker pool for managing long running processes, jwt authentication, etc.

## Development

The `Taskfile.yml` is all of the different tasks that can be run. To view all of them, run `task -l`. To run the application server and wait for changes run `task serve --watch`.

## TODO

* add jwt auth service
* add s3 service
