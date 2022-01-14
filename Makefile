lint:
	golint ./...

test:
	go test -v ./... -coverprofile=.coverage.txt -covermode=atomic