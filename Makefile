.PHONY: test build clean-test

test:
	go test -v ./...

clean-test:
	go test -count=1 -v ./...
