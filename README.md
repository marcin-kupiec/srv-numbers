# srv-numbers
## Introduction

This service handles number search.
Its responsibility is to expose REST API.

## Development
```
./script/lint
./script/test
```
or
```
make lint
make test
```

To run server in `docker` please run:
```
./script/run
```

To run server directly please run:
```
make run
```

## Documentation

Service uses `_env` configuration file, where log level (`DEBUG`, `INFO`, `ERROR`) and service port can be configured.\
Numbers are taken from `input.txt` file.\
Search algorithm is based on binary search.\
After starting the service, frontend is by default available on http://localhost:5411.

### REST API
- HTTP API Documentation in [swagger format](./docs/api/swagger.yaml)
