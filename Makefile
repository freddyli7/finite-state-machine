export GO111MODULE=on

.DEFAULT_GOAL := install

# Update the dependencies
update:
	go mod tidy
	go mod vendor

# Build and install a copy in bin
install:
	go build -o fsm ./

# Build for debug
debug:
	go install -i -gcflags="all=-N -l" .

# run unit tests on project packages
utest:

check:
	gosec ./...

coverage:
	go tool cover -html=a.out -o cover.html
