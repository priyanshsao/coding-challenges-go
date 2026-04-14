#format all the files
format fmt:
	gofmt -w ./

# get test coverage for overall repo
test-cover-all:
	go test ./... -coverprofile=coverage.txt covermode=atomic