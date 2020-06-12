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

## License

MIT