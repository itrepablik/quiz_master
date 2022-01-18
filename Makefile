build:
	go build -o bin/quiz_master
	go install

test:
	go test -mod=readonly -v ./.../

setup:
	go mod tidy
	go mod vendor

compile:
	echo "Compiling to linux platform"
	set GOOS=linux
	set GOARCH=amd64
	GOOS=linux GOARCH=arm64 go build -o bin/quiz_master

all: setup build test compile