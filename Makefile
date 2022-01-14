lint:
	golint ./...

test:
	go test -v ./... -coverprofile=converage.txt -covermode=atomic