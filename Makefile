.PHONY: prepare darwin linux windows publish clean release

all: publish

publish:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/main-mac main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/main.exe main.go

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/main main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main main.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/main.exe main.go

prepare:
	go mod vendor

clean:
	rm -rf ./bin