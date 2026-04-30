#format all the files
format fmt:
	gofmt -w ./

# get test coverage for overall repo
test-cover-all:
	go test ./... -coverprofile=coverage.txt covermode=atomic

# run test for all projects
test-all: 
	go test -v \
	./challenge-01 \
	./challenge-02