[![Build Status](https://travis-ci.org/moemoe89/integration-test-golang.svg?branch=master)](https://travis-ci.org/moemoe89/integration-test-golang)
[![codecov](https://codecov.io/gh/moemoe89/integration-test-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/integration-test-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/integration-test-golang)](https://goreportcard.com/report/github.com/moemoe89/integration-test-golang)

# INTEGRATION-TEST-GOLANG #

Example integration test using Dockertest

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ integration-test-golang/
  |     |
  |     +--+ main.go
  |        + repository/
  |        |
  |        +--+ repository.go
  |        |
  |        +--+ cassandra
  |        |  |
  |        |  +--+ cassandra.go
  |        |     + cassandra_test.go
  |        +--+ mysql
  |        |  |
  |        |  +--+ mysql.go
  |        |     + mysql_test.go
  |        |
  |        +--+ postgres
  |           |
  |           +--+ postgres.go
  |              + postgres_test.go
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup

* Setup Golang <https://golang.org>
* Setup Docker <https://www.docker.com>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> integration-test-golang
```

## How to Run Test
```
$ go test ./...
```
or
```
$ make test
```

## License

MIT