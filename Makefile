test:
	mkdir -p .coverage/html
	go test -v -cover -coverprofile=.coverage/test.coverage ./...
	go tool cover -html=.coverage/test.coverage -o .coverage/html/test.coverage.html
