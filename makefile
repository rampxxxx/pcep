clean:
	GOPATH=$(HOME)/go go clean -i ./...

test:
	go test -race -test.timeout 120s -v ./...

install:
	GOPATH=$(HOME)/go go install -v github.com/rampxxxx/pcep/...
